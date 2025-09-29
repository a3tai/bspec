#!/usr/bin/env swift

/**
 * BSpec Viewer - macOS Native Application
 *
 * A simple native macOS application that uses the Rust SDK
 * to provide native .bspec file handling through the FFI interface.
 */

import Foundation
import Cocoa

// Import the Rust SDK through C interface
// In a real app, this would be in a bridging header
/*
 #include "bspec.h"
 */

@main
class BSpecViewer: NSObject, NSApplicationDelegate {

    var window: NSWindow!
    var textView: NSTextView!
    var currentArchive: OpaquePointer?

    func applicationDidFinishLaunching(_ notification: Notification) {
        setupUI()

        // Handle file opening if launched with a file
        if let filePath = getFilePathFromArgs() {
            openBSpecFile(path: filePath)
        }
    }

    func applicationWillTerminate(_ notification: Notification) {
        // Clean up any open archives
        if let archive = currentArchive {
            bspec_close_archive(archive)
        }
    }

    func setupUI() {
        // Create main window
        window = NSWindow(
            contentRect: NSRect(x: 100, y: 100, width: 800, height: 600),
            styleMask: [.titled, .closable, .resizable, .miniaturizable],
            backing: .buffered,
            defer: false
        )

        window.title = "BSpec Viewer"
        window.makeKeyAndOrderFront(nil)

        // Create scroll view with text view
        let scrollView = NSScrollView(frame: window.contentView!.bounds)
        scrollView.autoresizingMask = [.width, .height]
        scrollView.hasVerticalScroller = true
        scrollView.hasHorizontalScroller = true

        textView = NSTextView(frame: scrollView.bounds)
        textView.isEditable = false
        textView.isSelectable = true
        textView.font = NSFont.monospacedSystemFont(ofSize: 12, weight: .regular)

        scrollView.documentView = textView
        window.contentView?.addSubview(scrollView)

        // Set up menu
        setupMenu()

        // Show welcome message
        showWelcomeMessage()
    }

    func setupMenu() {
        let mainMenu = NSMenu()

        // App menu
        let appMenuItem = NSMenuItem()
        mainMenu.addItem(appMenuItem)
        let appMenu = NSMenu()
        appMenuItem.submenu = appMenu

        appMenu.addItem(NSMenuItem(title: "About BSpec Viewer", action: #selector(showAbout), keyEquivalent: ""))
        appMenu.addItem(NSMenuItem.separator())
        appMenu.addItem(NSMenuItem(title: "Quit BSpec Viewer", action: #selector(NSApplication.terminate(_:)), keyEquivalent: "q"))

        // File menu
        let fileMenuItem = NSMenuItem(title: "File", action: nil, keyEquivalent: "")
        mainMenu.addItem(fileMenuItem)
        let fileMenu = NSMenu(title: "File")
        fileMenuItem.submenu = fileMenu

        fileMenu.addItem(NSMenuItem(title: "Open...", action: #selector(openFile), keyEquivalent: "o"))
        fileMenu.addItem(NSMenuItem(title: "Close", action: #selector(closeFile), keyEquivalent: "w"))

        NSApplication.shared.mainMenu = mainMenu
    }

    func getFilePathFromArgs() -> String? {
        let args = CommandLine.arguments
        if args.count > 1 {
            return args[1]
        }
        return nil
    }

    @objc func openFile() {
        let openPanel = NSOpenPanel()
        openPanel.allowedFileTypes = ["bspec"]
        openPanel.canChooseFiles = true
        openPanel.canChooseDirectories = false
        openPanel.allowsMultipleSelection = false

        if openPanel.runModal() == .OK {
            if let url = openPanel.url {
                openBSpecFile(path: url.path)
            }
        }
    }

    @objc func closeFile() {
        if let archive = currentArchive {
            bspec_close_archive(archive)
            currentArchive = nil
        }
        showWelcomeMessage()
    }

    @objc func showAbout() {
        let alert = NSAlert()
        alert.messageText = "BSpec Viewer"
        alert.informativeText = """
        Native macOS viewer for Business Specification Standard (.bspec) files.

        Powered by BSpec Rust SDK v1.0.0
        """
        alert.addButton(withTitle: "OK")
        alert.runModal()
    }

    func openBSpecFile(path: String) {
        // Close any existing archive
        if let archive = currentArchive {
            bspec_close_archive(archive)
        }

        // Open the new archive
        currentArchive = bspec_open_archive(path)

        guard let archive = currentArchive else {
            showError("Failed to open BSpec file: \(path)")
            return
        }

        // Get archive statistics
        guard let statsPtr = bspec_get_archive_stats(archive) else {
            showError("Failed to read archive statistics")
            return
        }

        let statsString = String(cString: statsPtr)
        bspec_free_string(UnsafeMutablePointer<CChar>(mutating: statsPtr))

        // Parse and display the statistics
        displayArchiveInfo(statsJSON: statsString, filePath: path)

        // Query all documents
        queryAndDisplayDocuments()
    }

    func displayArchiveInfo(statsJSON: String, filePath: String) {
        guard let data = statsJSON.data(using: .utf8),
              let stats = try? JSONSerialization.jsonObject(with: data) as? [String: Any] else {
            showError("Failed to parse archive statistics")
            return
        }

        var content = "# BSpec Archive Information\\n\\n"
        content += "**File:** `\\(URL(fileURLWithPath: filePath).lastPathComponent)`\\n"
        content += "**Path:** `\\(filePath)`\\n\\n"

        content += "## Archive Details\\n\\n"
        content += "- **Name:** \\(stats["name"] as? String ?? "Unknown")\\n"
        content += "- **Total Documents:** \\(stats["total_documents"] as? Int ?? 0)\\n"
        content += "- **Conformance Level:** \\(stats["conformance_level"] as? String ?? "Unknown")\\n"
        content += "- **Industry Profile:** \\(stats["industry_profile"] as? String ?? "Unknown")\\n"

        if let documentTypes = stats["document_types"] as? [String], !documentTypes.isEmpty {
            content += "- **Document Types:** \\(documentTypes.joined(separator: ", "))\\n"
        }

        if let domains = stats["domains"] as? [String], !domains.isEmpty {
            content += "- **Domains:** \\(domains.joined(separator: ", "))\\n"
        }

        if let createdAt = stats["created_at"] as? String {
            content += "- **Created:** \\(createdAt)\\n"
        }

        if let updatedAt = stats["updated_at"] as? String {
            content += "- **Updated:** \\(updatedAt)\\n"
        }

        content += "\\n---\\n\\n"

        textView.string = content
        window.title = "BSpec Viewer - \\(URL(fileURLWithPath: filePath).lastPathComponent)"
    }

    func queryAndDisplayDocuments() {
        guard let archive = currentArchive else { return }

        // Query all documents
        let queryJSON = "{\\"type\\": \\"*\\"}" // Get all document types
        guard let queryResult = bspec_query_documents(archive, queryJSON) else {
            appendToDisplay("\\n**Error:** Failed to query documents\\n")
            return
        }

        let documentCount = bspec_query_result_count(queryResult)

        var content = "## Documents (\\(documentCount))\\n\\n"

        if documentCount == 0 {
            content += "*No documents found in this archive.*\\n"
        } else {
            for i in 0..<documentCount {
                if let docPtr = bspec_get_document_json(queryResult, i) {
                    let docJSON = String(cString: docPtr)
                    bspec_free_string(UnsafeMutablePointer<CChar>(mutating: docPtr))

                    if let doc = parseDocument(json: docJSON) {
                        content += "### \\(doc.title)\\n\\n"
                        content += "- **ID:** `\\(doc.id)`\\n"
                        content += "- **Type:** \\(doc.type)\\n"
                        content += "- **Status:** \\(doc.status)\\n"
                        content += "- **Owner:** \\(doc.owner)\\n"
                        content += "- **Version:** \\(doc.version)\\n"

                        if !doc.content.isEmpty {
                            let preview = String(doc.content.prefix(200))
                            content += "- **Content Preview:** \\(preview)\\(doc.content.count > 200 ? "..." : "")\\n"
                        }

                        content += "\\n"
                    }
                }
            }
        }

        bspec_free_query_result(queryResult)
        appendToDisplay(content)
    }

    func parseDocument(json: String) -> DocumentInfo? {
        guard let data = json.data(using: .utf8),
              let doc = try? JSONSerialization.jsonObject(with: data) as? [String: Any] else {
            return nil
        }

        return DocumentInfo(
            id: doc["id"] as? String ?? "",
            title: doc["title"] as? String ?? "Untitled",
            type: doc["type"] as? String ?? "Unknown",
            status: doc["status"] as? String ?? "Unknown",
            owner: doc["owner"] as? String ?? "Unknown",
            version: doc["version"] as? String ?? "Unknown",
            content: doc["content"] as? String ?? ""
        )
    }

    func appendToDisplay(_ text: String) {
        DispatchQueue.main.async {
            self.textView.string += text
        }
    }

    func showWelcomeMessage() {
        let welcome = """
        # Welcome to BSpec Viewer

        This is a native macOS application for viewing Business Specification Standard (.bspec) files.

        ## Getting Started

        - **Open a file:** Use File → Open... or drag a .bspec file onto this window
        - **Supported formats:** .bspec archive files
        - **Features:**
          - Archive information display
          - Document listing and preview
          - Native macOS integration
          - Powered by Rust SDK for performance

        ## About BSpec

        The Business Specification Standard (BSpec) is a comprehensive framework for documenting enterprise software architecture and business requirements.

        Visit [bspec.dev](https://bspec.dev) for more information.
        """

        textView.string = welcome
        window.title = "BSpec Viewer"
    }

    func showError(_ message: String) {
        let alert = NSAlert()
        alert.messageText = "Error"
        alert.informativeText = message
        alert.alertStyle = .warning
        alert.addButton(withTitle: "OK")
        alert.runModal()
    }
}

struct DocumentInfo {
    let id: String
    let title: String
    let type: String
    let status: String
    let owner: String
    let version: String
    let content: String
}

// C function declarations (these would normally be in a bridging header)
@_silgen_name("bspec_open_archive")
func bspec_open_archive(_ path: UnsafePointer<CChar>) -> OpaquePointer?

@_silgen_name("bspec_close_archive")
func bspec_close_archive(_ archive: OpaquePointer)

@_silgen_name("bspec_get_archive_stats")
func bspec_get_archive_stats(_ archive: OpaquePointer) -> UnsafePointer<CChar>?

@_silgen_name("bspec_query_documents")
func bspec_query_documents(_ archive: OpaquePointer, _ query: UnsafePointer<CChar>) -> OpaquePointer?

@_silgen_name("bspec_query_result_count")
func bspec_query_result_count(_ result: OpaquePointer) -> Int32

@_silgen_name("bspec_get_document_json")
func bspec_get_document_json(_ result: OpaquePointer, _ index: Int) -> UnsafePointer<CChar>?

@_silgen_name("bspec_free_string")
func bspec_free_string(_ str: UnsafeMutablePointer<CChar>)

@_silgen_name("bspec_free_query_result")
func bspec_free_query_result(_ result: OpaquePointer)