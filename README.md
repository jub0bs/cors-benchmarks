# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.25.5) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.9.2).

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
single_vs_actual                505.1n ± 2%   523.8n ± 1%   +3.70% (p=0.000 n=10)
multiple_vs_actual              517.4n ± 1%   526.0n ± 1%   +1.68% (p=0.000 n=10)
pathological_vs_actual          548.9n ± 1%   587.1n ± 1%   +6.96% (p=0.000 n=10)
many_vs_actual                  533.1n ± 1%   531.8n ± 1%        ~ (p=0.670 n=10)
any_vs_actual                   535.7n ± 1%   525.5n ± 0%   -1.89% (p=0.000 n=10)
all_CORS_headers_vs_actual      548.4n ± 1%   578.5n ± 1%   +5.49% (p=0.000 n=10)
single_vs_preflight             438.9n ± 1%   416.4n ± 1%   -5.12% (p=0.000 n=10)
multiple_vs_preflight           442.6n ± 1%   417.4n ± 1%   -5.68% (p=0.000 n=10)
pathological_vs_preflight       418.2n ± 1%   432.8n ± 4%   +3.49% (p=0.000 n=10)
many_vs_preflight               398.4n ± 0%   363.6n ± 1%   -8.75% (p=0.000 n=10)
any_vs_preflight                436.1n ± 1%   408.6n ± 1%   -6.32% (p=0.000 n=10)
ACRH_vs_preflight               404.8n ± 1%   374.4n ± 1%   -7.50% (p=0.000 n=10)
all_CORS_headers_vs_preflight   405.9n ± 0%   390.6n ± 0%   -3.77% (p=0.000 n=10)
malicious_ACRH_vs_preflight     431.6n ± 1%   377.5n ± 0%  -12.54% (p=0.000 n=10)
geomean                         465.6n        454.7n        -2.34%

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
