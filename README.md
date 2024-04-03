# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.22.1) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.10.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.1.1).

## Running the benchmarks

Run the following commands in your shell:

```shell
git clone https://github.com/jub0bs/cors-benchmarks
cd cors-benchmarks
go test -run ^$ -bench . -benchtime 5s
```

## Some results

I've slightly redacted the results below for better readability. In particular,
I've annotated results with colored dots to ease visual comparison:

- A 🔴 marks a case in which jub0bs/cors fares worse than rs/cors,
  i.e. it both runs more slowly and allocates more memory.
- A 🟠 marks a case in which jub0bs/cors fares neither better nor worse than
  rs/cors, i.e. it either runs more slowly or allocates more memory
  (but not both).
- A 🟢 marks a case in which jub0bs/cors fares better than rs/cors,
  i.e. it both runs faster and allocates less memory.

```text
goos: darwin
goarch: amd64
pkg: github.com/jub0bs/cors-benchmarks
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz

no_CORS_________________________________-8       532 ns/op       1024 B/op    10 allocs/op

rs_cors_________________single_vs_actual-8       625 ns/op       1056 B/op    10 allocs/op
jub0bs_cors_____________single_vs_actual-8       665 ns/op       1072 B/op    11 allocs/op 🔴

rs_cors_______________multiple_vs_actual-8       625 ns/op       1056 B/op    10 allocs/op
jub0bs_cors___________multiple_vs_actual-8       669 ns/op       1072 B/op    11 allocs/op 🔴

rs_cors___________pathological_vs_actual-8       696 ns/op       1072 B/op    12 allocs/op
jub0bs_cors_______pathological_vs_actual-8       739 ns/op       1056 B/op    11 allocs/op 🟠

rs_cors___________________many_vs_actual-8       654 ns/op       1072 B/op    12 allocs/op
jub0bs_cors_______________many_vs_actual-8       639 ns/op       1056 B/op    11 allocs/op 🟢

rs_cors____________________any_vs_actual-8       615 ns/op       1056 B/op    10 allocs/op
jub0bs_cors________________any_vs_actual-8       625 ns/op       1056 B/op    11 allocs/op 🔴

rs_cors______________single_vs_preflight-8       507 ns/op        960 B/op     7 allocs/op
jub0bs_cors__________single_vs_preflight-8       456 ns/op        944 B/op     7 allocs/op 🟢

rs_cors____________multiple_vs_preflight-8       513 ns/op        960 B/op     7 allocs/op
jub0bs_cors________multiple_vs_preflight-8       462 ns/op        944 B/op     7 allocs/op 🟢

rs_cors________pathological_vs_preflight-8       525 ns/op        960 B/op     9 allocs/op
jub0bs_cors____pathological_vs_preflight-8       510 ns/op        928 B/op     7 allocs/op 🟢

rs_cors________________many_vs_preflight-8       483 ns/op        960 B/op     9 allocs/op
jub0bs_cors____________many_vs_preflight-8       422 ns/op        928 B/op     7 allocs/op 🟢

rs_cors_________________any_vs_preflight-8       503 ns/op        960 B/op     7 allocs/op
jub0bs_cors_____________any_vs_preflight-8       446 ns/op        944 B/op     7 allocs/op 🟢

rs_cors________________ACRH_vs_preflight-8       522 ns/op        984 B/op    10 allocs/op
jub0bs_cors____________ACRH_vs_preflight-8       419 ns/op        928 B/op     7 allocs/op 🟢

rs_cors______malicious_ACRH_vs_preflight-8  47203943 ns/op  121177076 B/op  1060 allocs/op 😱
jub0bs_cors__malicious_ACRH_vs_preflight-8       427 ns/op        928 B/op     7 allocs/op 🟢
```
