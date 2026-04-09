# minitools

## run bench

```shell
# beachmark
go test -bench=BenchmarkDataJson -benchmem github.com/qmaru/minitools/v2

# run test
go test -v -count 1 github.com/qmaru/minitools/v2
go test -v -count 1 github.com/qmaru/minitools/v2 -run='TestCompatibility'

# jsonv2
GOEXPERIMENT=jsonv2
go test -tags jsonv2 -bench=BenchmarkDataJson -benchmem github.com/qmaru/minitools/v2
go test -tags jsonv2 -v -count 1 github.com/qmaru/minitools/v2 -run='TestCompatibility'
```
