# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.25.6) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.11.0).

## Running the benchmarks

Run the following commands in your shell (preferably on an idle machine):

```shell
git clone https://github.com/jub0bs/cors-benchmarks
cd cors-benchmarks
go test -run ^$ -bench . -count 10 > bench.out
benchstat -col "/mw@(rs-cors jub0bs-cors)" -row "/req" bench.out
```

## Results

```text
goos: darwin
goarch: arm64
pkg: github.com/jub0bs/cors-benchmarks
cpu: Apple M4
                              │   rs-cors   │             jub0bs-cors             │
                              │   sec/op    │   sec/op     vs base                │
single_vs_actual                494.2n ± 1%   515.0n ± 1%   +4.21% (p=0.000 n=10)
multiple_vs_actual              512.3n ± 2%   525.9n ± 1%   +2.65% (p=0.000 n=10)
pathological_vs_actual          546.9n ± 1%   530.8n ± 1%   -2.95% (p=0.000 n=10)
many_vs_actual                  526.5n ± 1%   509.8n ± 1%   -3.16% (p=0.000 n=10)
any_vs_actual                   511.9n ± 1%   504.8n ± 1%   -1.39% (p=0.000 n=10)
all_CORS_headers_vs_actual      540.9n ± 0%   569.6n ± 1%   +5.31% (p=0.000 n=10)
single_vs_preflight             436.2n ± 1%   385.9n ± 1%  -11.53% (p=0.000 n=10)
multiple_vs_preflight           436.0n ± 1%   385.5n ± 0%  -11.58% (p=0.000 n=10)
pathological_vs_preflight       415.8n ± 1%   131.0n ± 1%  -68.49% (p=0.000 n=10)
many_vs_preflight               394.8n ± 1%   108.4n ± 1%  -72.55% (p=0.000 n=10)
any_vs_preflight                431.9n ± 1%   376.0n ± 0%  -12.94% (p=0.000 n=10)
ACRH_vs_preflight               401.3n ± 1%   110.8n ± 1%  -72.39% (p=0.000 n=10)
all_CORS_headers_vs_preflight   402.7n ± 1%   131.3n ± 0%  -67.39% (p=0.000 n=10)
malicious_ACRH_vs_preflight     431.1n ± 1%   115.1n ± 1%  -73.30% (p=0.000 n=10)
geomean                         459.9n        288.8n       -37.22%

                              │   rs-cors    │              jub0bs-cors               │
                              │     B/op     │     B/op      vs base                  │
single_vs_actual                1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
multiple_vs_actual              1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
pathological_vs_actual          1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
many_vs_actual                  1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
any_vs_actual                   1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
all_CORS_headers_vs_actual      1.062Ki ± 0%   1.078Ki ± 0%   +1.47% (p=0.000 n=10)
single_vs_preflight               976.0 ± 0%     928.0 ± 0%   -4.92% (p=0.000 n=10)
multiple_vs_preflight             976.0 ± 0%     928.0 ± 0%   -4.92% (p=0.000 n=10)
pathological_vs_preflight         960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
many_vs_preflight                 960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
any_vs_preflight                  976.0 ± 0%     928.0 ± 0%   -4.92% (p=0.000 n=10)
ACRH_vs_preflight                 968.0 ± 0%     208.0 ± 0%  -78.51% (p=0.000 n=10)
all_CORS_headers_vs_preflight     968.0 ± 0%     208.0 ± 0%  -78.51% (p=0.000 n=10)
malicious_ACRH_vs_preflight       968.0 ± 0%     208.0 ± 0%  -78.51% (p=0.000 n=10)
geomean                          1012.9          578.0       -42.93%
¹ all samples are equal

                              │  rs-cors   │             jub0bs-cors              │
                              │ allocs/op  │ allocs/op   vs base                  │
single_vs_actual                11.00 ± 0%   11.00 ± 0%        ~ (p=1.000 n=10) ¹
multiple_vs_actual              11.00 ± 0%   11.00 ± 0%        ~ (p=1.000 n=10) ¹
pathological_vs_actual          12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
many_vs_actual                  12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
any_vs_actual                   11.00 ± 0%   11.00 ± 0%        ~ (p=1.000 n=10) ¹
all_CORS_headers_vs_actual      11.00 ± 0%   12.00 ± 0%   +9.09% (p=0.000 n=10)
single_vs_preflight             8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
multiple_vs_preflight           8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
pathological_vs_preflight       9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
many_vs_preflight               9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
any_vs_preflight                8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
ACRH_vs_preflight               9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
all_CORS_headers_vs_preflight   9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
malicious_ACRH_vs_preflight     9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
geomean                         9.683        7.000       -27.71%
¹ all samples are equal
```

For more details about the benchmark labeled "malicious_ACRH_vs_preflight",
see https://github.com/rs/cors/issues/170.
