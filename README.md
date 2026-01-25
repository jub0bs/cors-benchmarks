# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.25.6) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.10.0).

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
single_vs_actual                497.3n ± 1%   511.2n ± 1%   +2.79% (p=0.000 n=10)
multiple_vs_actual              511.8n ± 1%   524.6n ± 1%   +2.52% (p=0.000 n=10)
pathological_vs_actual          542.6n ± 1%   523.8n ± 1%   -3.48% (p=0.000 n=10)
many_vs_actual                  521.9n ± 1%   507.8n ± 2%   -2.70% (p=0.000 n=10)
any_vs_actual                   511.3n ± 1%   514.3n ± 3%        ~ (p=0.481 n=10)
all_CORS_headers_vs_actual      551.5n ± 1%   568.6n ± 1%   +3.09% (p=0.000 n=10)
single_vs_preflight             436.5n ± 1%   409.2n ± 1%   -6.24% (p=0.000 n=10)
multiple_vs_preflight           434.7n ± 1%   411.4n ± 1%   -5.36% (p=0.000 n=10)
pathological_vs_preflight       415.2n ± 1%   387.4n ± 1%   -6.68% (p=0.000 n=10)
many_vs_preflight               393.6n ± 1%   359.4n ± 1%   -8.68% (p=0.000 n=10)
any_vs_preflight                429.3n ± 1%   399.5n ± 0%   -6.94% (p=0.000 n=10)
ACRH_vs_preflight               406.1n ± 2%   370.2n ± 1%   -8.86% (p=0.000 n=10)
all_CORS_headers_vs_preflight   403.8n ± 1%   385.6n ± 1%   -4.50% (p=0.000 n=10)
malicious_ACRH_vs_preflight     430.5n ± 1%   369.0n ± 1%  -14.28% (p=0.000 n=10)
geomean                         460.2n        440.3n        -4.32%

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
