# go-lang

### Reveal Go language crucial gears and idiomatic features.

📌 Edit module name:
```
go mod edit -module github.com/oleksiivelychko/go-lang
```
📌 List of module versions:
```
go list -m -versions github.com/oleksiivelychko/go-lang
```
📌 Run tests with benchmarks for the package:
```
go test -bench=. ./package_name
```
...when the output `BenchmarkPackageName-8                 41410035                29.00 ns/op`
means that the loop ran 41410035 times at a speed of 29.00 ns per loop on 8 CPU numbers.
