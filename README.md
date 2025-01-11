# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.23.0) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.5.0).

## Running the benchmarks

Run the following commands in your shell (preferably on an idle machine):

```shell
git clone https://github.com/jub0bs/cors-benchmarks
cd cors-benchmarks
go test -run ^$ -bench . -count 10 > benchmark_results.txt
benchstat -col "/mw@(rs-cors jub0bs-cors)" -row "/req" benchmark_results.txt
```

## Results

```text
goos: darwin
goarch: amd64
pkg: github.com/jub0bs/cors-benchmarks
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
                              │   rs-cors    │             jub0bs-cors             │
                              │    sec/op    │   sec/op     vs base                │
single_vs_actual                647.8n ±  1%   688.7n ± 1%   +6.31% (p=0.000 n=10)
multiple_vs_actual              664.5n ±  0%   694.0n ± 2%   +4.45% (p=0.000 n=10)
pathological_vs_actual          712.2n ±  1%   778.4n ± 1%   +9.28% (p=0.000 n=10)
many_vs_actual                  683.8n ±  1%   661.8n ± 2%   -3.22% (p=0.002 n=10)
any_vs_actual                   649.7n ±  1%   641.9n ± 3%        ~ (p=0.052 n=10)
all_CORS_headers_vs_actual      694.1n ±  3%   770.1n ± 2%  +10.95% (p=0.000 n=10)
single_vs_preflight             549.6n ± 48%   499.1n ± 1%   -9.18% (p=0.000 n=10)
multiple_vs_preflight           533.7n ±  3%   499.4n ± 1%   -6.41% (p=0.000 n=10)
pathological_vs_preflight       529.4n ±  1%   546.6n ± 1%   +3.25% (p=0.000 n=10)
many_vs_preflight               494.1n ±  1%   436.1n ± 1%  -11.73% (p=0.000 n=10)
any_vs_preflight                519.8n ±  1%   482.0n ± 1%   -7.26% (p=0.000 n=10)
ACRH_vs_preflight               490.6n ±  2%   445.2n ± 1%   -9.25% (p=0.000 n=10)
all_CORS_headers_vs_preflight   500.0n ±  1%   484.9n ± 1%   -3.01% (p=0.000 n=10)
malicious_ACRH_vs_preflight     526.5n ±  1%   483.9n ± 3%   -8.08% (p=0.000 n=10)
geomean                         579.9n         568.0n        -2.05%

                              │   rs-cors    │              jub0bs-cors              │
                              │     B/op     │     B/op      vs base                 │
single_vs_actual                1.047Ki ± 0%   1.047Ki ± 0%       ~ (p=1.000 n=10) ¹
multiple_vs_actual              1.047Ki ± 0%   1.047Ki ± 0%       ~ (p=1.000 n=10) ¹
pathological_vs_actual          1.047Ki ± 0%   1.031Ki ± 0%  -1.49% (p=0.000 n=10)
many_vs_actual                  1.047Ki ± 0%   1.031Ki ± 0%  -1.49% (p=0.000 n=10)
any_vs_actual                   1.047Ki ± 0%   1.031Ki ± 0%  -1.49% (p=0.000 n=10)
all_CORS_headers_vs_actual      1.062Ki ± 0%   1.078Ki ± 0%  +1.47% (p=0.000 n=10)
single_vs_preflight               976.0 ± 0%     944.0 ± 0%  -3.28% (p=0.000 n=10)
multiple_vs_preflight             976.0 ± 0%     944.0 ± 0%  -3.28% (p=0.000 n=10)
pathological_vs_preflight         960.0 ± 0%     928.0 ± 0%  -3.33% (p=0.000 n=10)
many_vs_preflight                 960.0 ± 0%     928.0 ± 0%  -3.33% (p=0.000 n=10)
any_vs_preflight                  976.0 ± 0%     944.0 ± 0%  -3.28% (p=0.000 n=10)
ACRH_vs_preflight                 968.0 ± 0%     928.0 ± 0%  -4.13% (p=0.000 n=10)
all_CORS_headers_vs_preflight     968.0 ± 0%     928.0 ± 0%  -4.13% (p=0.000 n=10)
malicious_ACRH_vs_preflight       968.0 ± 0%     928.0 ± 0%  -4.13% (p=0.000 n=10)
geomean                          1012.9          989.7       -2.29%
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
pathological_vs_preflight       9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
many_vs_preflight               9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
any_vs_preflight                8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
ACRH_vs_preflight               9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
all_CORS_headers_vs_preflight   9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
malicious_ACRH_vs_preflight     9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
geomean                         9.683        8.549       -11.71%
¹ all samples are equal
```

For more details about the benchmark labeled "malicious_ACRH_vs_preflight",
see https://github.com/rs/cors/issues/170.
