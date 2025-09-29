#!/bin/bash
# Build script for BSpec Rust SDK
# Generates cross-platform libraries for integration with other languages

set -e

PROJECT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$PROJECT_DIR"

echo "🦀 Building BSpec Rust SDK..."

# Clean previous builds
echo "🧹 Cleaning previous builds..."
cargo clean
rm -rf target/
rm -rf lib/
mkdir -p lib include

# Development build
echo "📦 Building development version..."
cargo build --features=c-api

# Release build
echo "🚀 Building release version..."
cargo build --release --features=c-api

# Copy header file
echo "📄 Copying header file..."
cp include/bspec.h lib/

if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "🍎 Building for macOS..."

    # Build for different architectures
    cargo build --release --target x86_64-apple-darwin --features=c-api
    cargo build --release --target aarch64-apple-darwin --features=c-api

    # Create universal binary
    echo "🔗 Creating universal binary..."
    lipo -create \
        target/x86_64-apple-darwin/release/libbspec_sdk.a \
        target/aarch64-apple-darwin/release/libbspec_sdk.a \
        -output lib/libbspec_sdk_universal.a

    # Copy individual architectures
    cp target/x86_64-apple-darwin/release/libbspec_sdk.a lib/libbspec_sdk_x86_64.a
    cp target/aarch64-apple-darwin/release/libbspec_sdk.a lib/libbspec_sdk_aarch64.a

    # Copy dynamic library
    cp target/release/libbspec_sdk.dylib lib/ 2>/dev/null || echo "Dynamic library not built"

    echo "✅ macOS libraries built:"
    echo "  - lib/libbspec_sdk_universal.a (Universal binary)"
    echo "  - lib/libbspec_sdk_x86_64.a (Intel)"
    echo "  - lib/libbspec_sdk_aarch64.a (Apple Silicon)"

elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    echo "🐧 Building for Linux..."

    # Copy libraries
    cp target/release/libbspec_sdk.a lib/
    cp target/release/libbspec_sdk.so lib/ 2>/dev/null || echo "Dynamic library not built"

    echo "✅ Linux libraries built:"
    echo "  - lib/libbspec_sdk.a (Static)"
    echo "  - lib/libbspec_sdk.so (Dynamic)"

elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" ]]; then
    echo "🪟 Building for Windows..."

    # Copy libraries
    cp target/release/bspec_sdk.lib lib/ 2>/dev/null || cp target/release/libbspec_sdk.a lib/bspec_sdk.lib
    cp target/release/bspec_sdk.dll lib/ 2>/dev/null || echo "Dynamic library not built"

    echo "✅ Windows libraries built:"
    echo "  - lib/bspec_sdk.lib (Static)"
    echo "  - lib/bspec_sdk.dll (Dynamic)"
fi

# Generate pkg-config file
echo "📋 Generating pkg-config file..."
cat > lib/bspec-sdk.pc << EOF
prefix=$PROJECT_DIR
libdir=\${prefix}/lib
includedir=\${prefix}/include

Name: BSpec SDK
Description: Business Specification Standard SDK
Version: 1.0.0
Libs: -L\${libdir} -lbspec_sdk
Cflags: -I\${includedir}
EOF

# Run tests
echo "🧪 Running tests..."
cargo test --features=c-api

# Generate documentation
echo "📚 Generating documentation..."
cargo doc --no-deps --features=c-api

echo ""
echo "🎉 Build complete!"
echo ""
echo "Integration files:"
echo "  - lib/bspec.h (C header)"
echo "  - lib/bspec-sdk.pc (pkg-config)"
if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "  - lib/libbspec_sdk_universal.a (Universal static library)"
fi
echo ""
echo "To integrate with Go CLI:"
echo "  CGO_CFLAGS=\"-I$PROJECT_DIR/lib\""
echo "  CGO_LDFLAGS=\"-L$PROJECT_DIR/lib -lbspec_sdk\""
echo ""
echo "To integrate with macOS app:"
echo "  Add lib/libbspec_sdk_universal.a to Xcode project"
echo "  Add lib/bspec.h to bridging header"
echo ""
echo "Documentation: file://$PROJECT_DIR/target/doc/bspec_sdk/index.html"