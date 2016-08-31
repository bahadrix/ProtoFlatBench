Protocol Buffers vs  Flat Buffers Benchmark
===================

----------

Libraries
-------------

Protocol Buffers
: github.com/golang/protobuf
Flat Buffers
: github.com/google/flatbuffers
JSON
: github.com/dustin/gojson

Results
-------------

Benchmark Results
```
BenchmarkFlatWrite-8     1000000     1672 ns/op     3144 B/op    18 allocs/op
BenchmarkFlatRead-8     30000000       42.4 ns/op     32 B/op     1 allocs/op
BenchmarkProtoWrite-8    2000000      755 ns/op      936 B/op     9 allocs/op
BenchmarkProtoRead-8     2000000      694 ns/op      768 B/op     6 allocs/op
BenchmarkJSONWrite-8      500000     2815 ns/op     1976 B/op     9 allocs/op
BenchmarkJSONRead-8       200000    10300 ns/op     1360 B/op    16 allocs/op
```

Serialized byte sizes:

| Serializer | Size |
|-----------------|-----------|
| ProtocolBuffers | 482 Bytes |
| Flatbuffers     | 528 Bytes |
| JSON            | 676 Bytes |