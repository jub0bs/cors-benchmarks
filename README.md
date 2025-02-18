# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.24.0) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.5.6).

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
                              │   rs-cors   │             jub0bs-cors             │
                              │   sec/op    │   sec/op     vs base                │
single_vs_actual                633.5n ± 1%   669.1n ± 1%   +5.62% (p=0.000 n=10)
multiple_vs_actual              646.1n ± 1%   674.0n ± 1%   +4.31% (p=0.000 n=10)
pathological_vs_actual          704.2n ± 6%   787.0n ± 5%  +11.75% (p=0.000 n=10)
many_vs_actual                  660.6n ± 2%   644.6n ± 1%   -2.42% (p=0.000 n=10)
any_vs_actual                   631.5n ± 1%   631.9n ± 1%        ~ (p=0.782 n=10)
all_CORS_headers_vs_actual      676.4n ± 2%   733.9n ± 1%   +8.52% (p=0.000 n=10)
single_vs_preflight             516.9n ± 1%   491.9n ± 1%   -4.85% (p=0.000 n=10)
multiple_vs_preflight           522.9n ± 1%   492.7n ± 1%   -5.78% (p=0.000 n=10)
pathological_vs_preflight       517.8n ± 1%   546.6n ± 1%   +5.56% (p=0.000 n=10)
many_vs_preflight               476.3n ± 0%   428.6n ± 0%  -10.00% (p=0.000 n=10)
any_vs_preflight                508.2n ± 1%   481.3n ± 1%   -5.28% (p=0.000 n=10)
ACRH_vs_preflight               482.3n ± 1%   447.3n ± 1%   -7.26% (p=0.000 n=10)
all_CORS_headers_vs_preflight   488.7n ± 1%   479.6n ± 1%   -1.86% (p=0.000 n=10)
malicious_ACRH_vs_preflight     520.2n ± 1%   480.4n ± 1%   -7.64% (p=0.000 n=10)
geomean                         565.0n        560.1n        -0.87%

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
