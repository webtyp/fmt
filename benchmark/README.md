# fmt Benchmark Suite

Automated benchmark tools to measure and compare performance between standard Go libraries and fmt implementations.



## Binary Size Comparison

[Standard Library Example](bench-binary-size/standard-lib/main.go) | [fmt Example](bench-binary-size/webtyp-lib/main.go)

<!-- This table is automatically generated from build-and-measure.sh -->
*Last updated: 2026-11-05 12:43:26*

| Build Type | Parameters | Standard Library<br/>`go build` | fmt<br/>`tinygo build` | Size Reduction | Performance |
|------------|------------|------------------|------------|----------------|-------------|
| 🖥️ **Default Native** | `-ldflags="-s -w"` | 1.5 MB | 1.4 MB | **-124.0 KB** | ➖ **7.8%** |
| 🌐 **Default WASM** | `(default -opt=z)` | 705.4 KB | 272.0 KB | **-433.3 KB** | ✅ **61.4%** |
| 🌐 **Ultra WASM** | `-no-debug -panic=trap -scheduler=none -gc=leaking -target wasm` | 156.2 KB | 26.1 KB | **-130.2 KB** | 🏆 **83.3%** |
| 🌐 **Speed WASM** | `-opt=2 -target wasm` | 959.2 KB | 410.5 KB | **-548.7 KB** | ✅ **57.2%** |
| 🌐 **Debug WASM** | `-opt=0 -target wasm` | 2.0 MB | 795.9 KB | **-1.2 MB** | ✅ **60.3%** |

### 🎯 Performance Summary

- 🏆 **Peak Reduction: 83.3%** (Best optimization)
- ✅ **Average WebAssembly Reduction: 65.6%**
- ✅ **Average Native Reduction: 7.8%**
- 📦 **Total Size Savings: 2.4 MB across all builds**

#### Performance Legend
- ❌ Poor (<5% reduction)
- ➖ Fair (5-15% reduction)
- ✅ Good (15-70% reduction)
- 🏆 Outstanding (>70% reduction)


## Memory Usage Comparison

[Standard Library Example](bench-memory-alloc/standard) | [fmt Example](bench-memory-alloc/webtyp)

<!-- This table is automatically generated from memory-benchmark.sh -->
*Last updated: 2026-11-05 12:49:56*

Performance benchmarks comparing memory allocation patterns between standard Go library and fmt:

| 🧪 **Benchmark Category** | 📚 **Library** | 💾 **Memory/Op** | 🔢 **Allocs/Op** | ⏱️ **Time/Op** | 📈 **Memory Trend** | 🎯 **Alloc Trend** | 🏆 **Performance** |
|----------------------------|----------------|-------------------|-------------------|-----------------|---------------------|---------------------|--------------------|
| 📝 **String Processing** | 📊 Standard | `640 B / 925741 OP` | `25` | `1.2μs` | - | - | - |
| | 🚀 fmt | `464 B / 244772 OP` | `17` | `4.8μs` | 🏆 **27.5% less** | 🏆 **32.0% less** | 🏆 **Excellent** |
| 🔢 **Number Processing** | 📊 Standard | `256 B / 2994324 OP` | `9` | `398ns` | - | - | - |
| | 🚀 fmt | `320 B / 1202761 OP` | `17` | `1.0μs` | ⚠️ **25.0% more** | ❌ **88.9% more** | ❌ **Poor** |
| 🔄 **Mixed Operations** | 📊 Standard | `208 B / 1773916 OP` | `12` | `682ns` | - | - | - |
| | 🚀 fmt | `176 B / 554125 OP` | `12` | `2.1μs` | ✅ **15.4% less** | ➖ **Same** | ✅ **Good** |

### 🎯 Performance Summary

- 💾 **Memory Efficiency**: ✅ **Good** (Memory efficient) (-6.0% average change)
- 🔢 **Allocation Efficiency**: ⚠️ **Caution** (More allocations) (19.0% average change)
- 📊 **Benchmarks Analyzed**: 3 categories
- 🎯 **Optimization Focus**: Binary size reduction vs runtime efficiency

### ⚖️ Trade-offs Analysis

The benchmarks reveal important trade-offs between **binary size** and **runtime performance**:

#### 📦 **Binary Size Benefits** ✅
- 🏆 **16-84% smaller** compiled binaries
- 🌐 **Superior WebAssembly** compression ratios
- 🚀 **Faster deployment** and distribution
- 💾 **Lower storage** requirements

#### 🧠 **Runtime Memory Considerations** ⚠️
- 📈 **Higher allocation overhead** during execution
- 🗑️ **Increased GC pressure** due to allocation patterns
- ⚡ **Trade-off optimizes** for distribution size over runtime efficiency
- 🔄 **Different optimization strategy** than standard library

#### 🎯 **Optimization Recommendations**
| 🎯 **Use Case** | 💡 **Recommendation** | 🔧 **Best For** |
|-----------------|------------------------|------------------|
| 🌐 WebAssembly Apps | ✅ **fmt** | Size-critical web deployment |
| 📱 Embedded Systems | ✅ **fmt** | Resource-constrained devices |
| ☁️ Edge Computing | ✅ **fmt** | Fast startup and deployment |
| 🏢 Memory-Intensive Server | ⚠️ **Standard Library** | High-throughput applications |
| 🔄 High-Frequency Processing | ⚠️ **Standard Library** | Performance-critical workloads |

#### 📊 **Performance Legend**
- 🏆 **Excellent** (Better performance)
- ✅ **Good** (Acceptable trade-off)
- ⚠️ **Caution** (Higher resource usage)
- ❌ **Poor** (Significant overhead)


## Quick Usage 🚀

```bash
# Run complete benchmark (recommended)
./build-and-measure.sh

# Clean generated files
./clean-all.sh

# Update README with existing data only (does not re-run benchmarks)
./update-readme.sh

# Run all memory and binary size benchmarks (without updating README)
./run-all-benchmarks.sh

# Run only memory benchmarks
./memory-benchmark.sh
```

## What Gets Measured 📊

1.  **Binary Size Comparison**: Native + WebAssembly builds with multiple optimization levels. This compares the compiled output size of projects using the standard Go library versus fmt.
2.  **Memory Allocation**: Measures Bytes/op, Allocations/op, and execution time (ns/op) for benchmark categories. This helps in understanding the memory efficiency of fmt compared to standard library operations.
    *   **String Processing**: Benchmarks operations like case conversion, text manipulation, etc.
    *   **Number Processing**: Benchmarks numeric formatting, conversion operations, etc.
    *   **Mixed Operations**: Benchmarks scenarios involving a combination of string and numeric operations.

## Current Performance Status

**Target**: Achieve memory usage close to standard library while maintaining binary size benefits.

**Latest Results**: the two generated tables above are the only source of numbers (run `./build-and-measure.sh` to refresh them). In short: fmt is much smaller in WebAssembly, uses less memory in string work, and is slower per operation on the server; Number Processing still allocates more than the standard library and is the next optimization target.

📋 **Memory Optimization Guide**: See [`MEMORY_REDUCTION.md`](./MEMORY_REDUCTION.md) for comprehensive techniques and best practices to replace Go standard libraries with fmt's optimized implementations. Essential reading for efficient string and numeric processing in TinyGo WebAssembly applications.

## Requirements

- **Go 1.21+**
- **TinyGo** (optional, but recommended for full WebAssembly testing and to achieve smallest binary sizes).


## Troubleshooting

**TinyGo Not Found:**
```
❌ TinyGo is not installed. Building only standard Go binaries.
```
Install TinyGo from: https://tinygo.org/getting-started/install/

**Permission Issues (Linux/macOS/WSL):**
If you encounter permission errors when trying to run the shell scripts, make them executable:
```bash
chmod +x *.sh
```

**Build Failures:**
- Ensure you're in the `benchmark/` directory
- Verify fmt library is available in the parent directory




