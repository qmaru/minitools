# minitools

## run bench

```shell
# benchmark
go test -bench=BenchmarkDataJson -benchmem github.com/qmaru/minitools/v2

# run test
go test -v -count 1 github.com/qmaru/minitools/v2
go test -v -count 1 github.com/qmaru/minitools/v2 -run='TestCompatibility'
```
