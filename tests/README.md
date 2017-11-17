Run tests:
```go	
    go test -v      // 'v' stands for verbose output
```
Run benchmark:
```go
    go test -bench=.
```
While benchmarking, the benchmark function must run the target code b.N times.
During benchmark execution, b.N is adjusted until the benchmark function lasts
long to be timely reliable. The output
```go
    BenchmarkHello 100000   100 ns/op
```
means, the loop ran 100000 times at a speed of 100 ns/op.