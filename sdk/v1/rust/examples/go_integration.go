package main

/*
#cgo CFLAGS: -I../lib
#cgo LDFLAGS: -L../lib -lbspec_sdk
#include "bspec.h"
#include <stdlib.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"log"
	"unsafe"
)

// BSpecArchive wraps the C Archive pointer
type BSpecArchive struct {
	archive *C.Archive
}

// QueryResult wraps the C QueryResult pointer
type QueryResult struct {
	result *C.QueryResult
}

// NewArchive creates a new BSpec archive
func NewArchive(name string) (*BSpecArchive, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	archive := C.bspec_create_archive(cName)
	if archive == nil {
		return nil, fmt.Errorf("failed to create archive")
	}

	return &BSpecArchive{archive: archive}, nil
}

// OpenArchive opens an existing .bspec file
func OpenArchive(path string) (*BSpecArchive, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	archive := C.bspec_open_archive(cPath)
	if archive == nil {
		return nil, fmt.Errorf("failed to open archive: %s", path)
	}

	return &BSpecArchive{archive: archive}, nil
}

// Save saves the archive to a file
func (a *BSpecArchive) Save(path string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	result := C.bspec_save_archive(a.archive, cPath)
	if result != C.BSPEC_SUCCESS {
		return fmt.Errorf("failed to save archive: error code %d", int(result))
	}

	return nil
}

// Close closes and frees the archive
func (a *BSpecArchive) Close() {
	if a.archive != nil {
		C.bspec_close_archive(a.archive)
		a.archive = nil
	}
}

// Query executes a query against the archive
func (a *BSpecArchive) Query(query map[string]interface{}) (*QueryResult, error) {
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query: %v", err)
	}

	cQuery := C.CString(string(queryJSON))
	defer C.free(unsafe.Pointer(cQuery))

	result := C.bspec_query_documents(a.archive, cQuery)
	if result == nil {
		return nil, fmt.Errorf("query failed")
	}

	return &QueryResult{result: result}, nil
}

// GetStats returns archive statistics
func (a *BSpecArchive) GetStats() (map[string]interface{}, error) {
	cStats := C.bspec_get_archive_stats(a.archive)
	if cStats == nil {
		return nil, fmt.Errorf("failed to get stats")
	}
	defer C.bspec_free_string((*C.char)(unsafe.Pointer(cStats)))

	statsJSON := C.GoString(cStats)
	var stats map[string]interface{}
	if err := json.Unmarshal([]byte(statsJSON), &stats); err != nil {
		return nil, fmt.Errorf("failed to parse stats: %v", err)
	}

	return stats, nil
}

// Count returns the number of documents in the query result
func (r *QueryResult) Count() int {
	return int(C.bspec_query_result_count(r.result))
}

// GetDocument returns a document by index as a JSON string
func (r *QueryResult) GetDocument(index int) (string, error) {
	cDoc := C.bspec_get_document_json(r.result, C.size_t(index))
	if cDoc == nil {
		return "", fmt.Errorf("invalid index or document not found")
	}
	defer C.bspec_free_string((*C.char)(unsafe.Pointer(cDoc)))

	return C.GoString(cDoc), nil
}

// GetAllDocuments returns all documents in the query result
func (r *QueryResult) GetAllDocuments() ([]map[string]interface{}, error) {
	count := r.Count()
	documents := make([]map[string]interface{}, count)

	for i := 0; i < count; i++ {
		docJSON, err := r.GetDocument(i)
		if err != nil {
			return nil, fmt.Errorf("failed to get document %d: %v", i, err)
		}

		var doc map[string]interface{}
		if err := json.Unmarshal([]byte(docJSON), &doc); err != nil {
			return nil, fmt.Errorf("failed to parse document %d: %v", i, err)
		}

		documents[i] = doc
	}

	return documents, nil
}

// Close frees the query result
func (r *QueryResult) Close() {
	if r.result != nil {
		C.bspec_free_query_result(r.result)
		r.result = nil
	}
}

// Example usage
func main() {
	// Get SDK version
	version := C.GoString(C.bspec_get_version())
	fmt.Printf("BSpec SDK Version: %s\n", version)

	// Create a new archive
	archive, err := NewArchive("Test Archive")
	if err != nil {
		log.Fatalf("Failed to create archive: %v", err)
	}
	defer archive.Close()

	// Get archive stats
	stats, err := archive.GetStats()
	if err != nil {
		log.Fatalf("Failed to get stats: %v", err)
	}

	fmt.Printf("Archive Stats: %+v\n", stats)

	// Query all documents (empty archive)
	query := map[string]interface{}{
		"type": "MSN",
	}

	result, err := archive.Query(query)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer result.Close()

	fmt.Printf("Found %d documents\n", result.Count())

	// Example of opening an existing archive
	// archive2, err := OpenArchive("example.bspec")
	// if err != nil {
	//     log.Printf("Could not open example.bspec: %v", err)
	// } else {
	//     defer archive2.Close()
	//     // Work with existing archive...
	// }

	fmt.Println("Integration test completed successfully!")
}