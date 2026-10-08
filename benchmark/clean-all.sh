#!/bin/bash

# clean-all.sh - Clean all benchmark artifacts and temporary files
# This script removes generated binaries, test artifacts, and build cache

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "🧹 Cleaning fmt Benchmark Artifacts"
echo "==========================================="

# Function to clean binary artifacts
clean_binary_artifacts() {
    echo "📋 Cleaning binary size benchmark artifacts..."
    
    if [[ -d "bench-binary-size" ]]; then
        # Clean binary files in subdirectories
        find bench-binary-size -type f \( -name "standard*" -o -name "webtyp*" \) ! -name "*.go" ! -name "*.mod" -exec rm -f {} \;
        
        # Clean WASM files specifically
        find bench-binary-size -name "*.wasm" -exec rm -f {} \;
        
        echo "✅ Binary artifacts cleaned"
    else
        echo "ℹ️  Binary benchmark directory not found, skipping"
    fi
}

# Function to clean memory benchmark artifacts
clean_memory_artifacts() {
    echo "📋 Cleaning memory benchmark artifacts..."
    
    if [[ -d "bench-memory-alloc" ]]; then
        # Clean test cache and binaries
        find bench-memory-alloc -name "*.test" -exec rm -f {} \;
        find bench-memory-alloc -name "*.out" -exec rm -f {} \;
        
        echo "✅ Memory benchmark artifacts cleaned"
    else
        echo "ℹ️  Memory benchmark directory not found, skipping"
    fi
}

# Function to clean build artifacts
clean_build_artifacts() {
    echo "📋 Cleaning build artifacts..."
    
    # Remove analyzer binaries (both Unix and Windows versions)
    if [[ -f "analyzer" ]]; then
        rm -f analyzer
        echo "✅ Analyzer binary (Unix) removed"
    fi
    
    if [[ -f "analyzer.exe" ]]; then
        rm -f analyzer.exe
        echo "✅ Analyzer binary (Windows) removed"
    fi
    
    # Remove benchmark-test binary
    if [[ -f "benchmark-test" ]]; then
        rm -f benchmark-test
        echo "✅ Benchmark-test binary removed"
    fi
    
    # Clean Go build cache for this module
    go clean -cache -testcache
    echo "✅ Go build cache cleaned"
}

# Function to clean temporary files
clean_temp_files() {
    echo "📋 Cleaning temporary files..."
    
    # Remove temporary README files
    find . -name "*.tmp" -exec rm -f {} \;
    find . -name "README.md.tmp" -exec rm -f {} \;
    
    # Remove log files
    find . -name "*.log" -exec rm -f {} \;
    
    # Remove backup files
    find . -name "*~" -exec rm -f {} \;
    find . -name "*.bak" -exec rm -f {} \;
    
    echo "✅ Temporary files cleaned"
}

# Function to clean memory tool artifacts (legacy)
clean_legacy_artifacts() {
    echo "📋 Cleaning legacy artifacts..."
    
    if [[ -d "memory-tool" ]]; then
        find memory-tool -name "*.test" -exec rm -f {} \;
        find memory-tool -name "*.out" -exec rm -f {} \;
        echo "✅ Legacy memory-tool artifacts cleaned"
    fi
}

# Parse command line arguments
case "${1:-all}" in
    "binary")
        clean_binary_artifacts
        ;;
    "memory")
        clean_memory_artifacts
        ;;
    "build")
        clean_build_artifacts
        ;;
    "temp")
        clean_temp_files
        ;;
    "all")
        clean_binary_artifacts
        clean_memory_artifacts
        clean_build_artifacts
        clean_temp_files
        clean_legacy_artifacts
        ;;
    "help"|"-h"|"--help")
        echo "Usage: $0 [binary|memory|build|temp|all|help]"
        echo ""
        echo "Commands:"
        echo "  binary  - Clean only binary size artifacts"
        echo "  memory  - Clean only memory benchmark artifacts"
        echo "  build   - Clean only build artifacts (binaries, cache)"
        echo "  temp    - Clean only temporary files (logs, backups)"
        echo "  all     - Clean everything (default)"
        echo "  help    - Show this help message"
        echo ""
        echo "Note: This will remove all generated binaries and test results."
        echo "You'll need to rebuild using build-and-measure.sh and run-all-benchmarks.sh"
        exit 0
        ;;
    *)
        echo "❌ Unknown command: $1"
        echo "Use '$0 help' for usage information"
        exit 1
        ;;
esac

echo "✅ Cleanup completed successfully!"
echo "==========================================="
