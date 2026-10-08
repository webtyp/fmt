#!/bin/bash

set -e

BENCHMARK_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MEMORY_BENCH_DIR="$BENCHMARK_DIR/bench-memory-alloc"

# Function to get the correct analyzer binary name based on OS
get_analyzer_name() {
    if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" || "$OSTYPE" == "win32" ]]; then
        echo "analyzer.exe"
    else
        echo "analyzer"
    fi
}

ANALYZER_BINARY=$(get_analyzer_name)

echo "🧠 Running memory allocation benchmarks..."

# Run standard library benchmarks
echo "📊 Running standard library benchmarks..."
cd "$MEMORY_BENCH_DIR/standard"
STANDARD_RESULTS=$(go test -bench=. -benchmem | grep -E '^Benchmark')

# Run fmt benchmarks
echo "📊 Running fmt benchmarks..."
cd "$MEMORY_BENCH_DIR/webtyp"
WEBTYP_RESULTS=$(go test -bench=. -benchmem | grep -E '^Benchmark')

# Generate memory benchmark section for README
echo "📝 Generating memory benchmark results..."

# Create temporary benchmark results file
TEMP_RESULTS="$BENCHMARK_DIR/benchmark_results.md"

cat > "$TEMP_RESULTS" << EOF

## Memory Allocation Benchmarks

### Standard Library vs fmt Performance

#### String Processing Benchmarks
\`\`\`
Standard Library:
$STANDARD_RESULTS

fmt:
$WEBTYP_RESULTS
\`\`\`

### Performance Analysis

The benchmarks show memory allocation differences between:
- **Standard Library**: Traditional Go string operations using \`strings\`, \`fmt\`, \`strconv\` packages
- **fmt**: Custom implementations optimized for minimal allocations
- **fmt Pointers**: Zero-allocation operations using string pointer modifications

**Key Benefits:**
- 🔥 **Reduced Allocations**: fmt minimizes memory allocations through efficient implementations
- ⚡ **Pointer Optimization**: Using \`Apply()\` method with string pointers eliminates temporary allocations
- 📦 **Memory Efficiency**: Lower memory footprint especially important for embedded systems and WebAssembly

EOF

echo "✅ Memory benchmark results generated"
echo "📄 Results saved to: $TEMP_RESULTS"

# Now update the main README with these results
cd "$BENCHMARK_DIR"
go build -o "$ANALYZER_BINARY" .
./"$ANALYZER_BINARY" memory "$TEMP_RESULTS"

# Clean up temporary file
# rm -f "$TEMP_RESULTS"

echo "✅ Memory benchmarks completed and README updated"
