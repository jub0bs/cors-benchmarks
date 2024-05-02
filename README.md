# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.22.2) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.0), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.1.3).

## Running the benchmarks

Run the following commands in your shell (preferably on an idle machine):

```shell
git clone https://github.com/jub0bs/cors-benchmarks
cd cors-benchmarks
go test -run ^$ -bench . -count 10 > new.txt
benchstat -col "/mw@(rs-cors jub0bs-cors)" -row "/req" new.txt
```

## Results

```text
goos: darwin
goarch: amd64
pkg: github.com/jub0bs/cors-benchmarks
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
                              │   rs-cors   │             jub0bs-cors             │
                              │   sec/op    │   sec/op     vs base                │
single_vs_actual                645.0n ± 0%   681.8n ± 1%   +5.71% (p=0.000 n=10)
multiple_vs_actual              654.6n ± 1%   695.6n ± 0%   +6.27% (p=0.000 n=10)
pathological_vs_actual          723.5n ± 1%   764.9n ± 1%   +5.72% (p=0.000 n=10)
many_vs_actual                  689.4n ± 1%   665.6n ± 1%   -3.45% (p=0.000 n=10)
any_vs_actual                   657.8n ± 2%   655.2n ± 1%        ~ (p=0.542 n=10)
all_CORS_headers_vs_actual      705.4n ± 1%   779.4n ± 1%  +10.48% (p=0.000 n=10)
single_vs_preflight             532.1n ± 2%   478.7n ± 2%  -10.04% (p=0.000 n=10)
multiple_vs_preflight           537.5n ± 2%   481.3n ± 2%  -10.45% (p=0.000 n=10)
pathological_vs_preflight       550.0n ± 2%   531.7n ± 1%   -3.32% (p=0.000 n=10)
many_vs_preflight               504.8n ± 3%   434.8n ± 3%  -13.87% (p=0.000 n=10)
any_vs_preflight                524.1n ± 2%   466.8n ± 2%  -10.93% (p=0.000 n=10)
ACRH_vs_preflight               495.8n ± 2%   437.1n ± 1%  -11.83% (p=0.000 n=10)
all_CORS_headers_vs_preflight   502.8n ± 3%   462.4n ± 1%   -8.03% (p=0.000 n=10)
malicious_ACRH_vs_preflight     491.4n ± 2%   440.6n ± 2%  -10.34% (p=0.000 n=10)
geomean                         580.9n        556.5n        -4.20%

                              │   rs-cors    │              jub0bs-cors              │
                              │     B/op     │     B/op      vs base                 │
single_vs_actual                1.031Ki ± 0%   1.047Ki ± 0%  +1.52% (p=0.000 n=10)
multiple_vs_actual              1.031Ki ± 0%   1.047Ki ± 0%  +1.52% (p=0.000 n=10)
pathological_vs_actual          1.047Ki ± 0%   1.031Ki ± 0%  -1.49% (p=0.000 n=10)
many_vs_actual                  1.047Ki ± 0%   1.031Ki ± 0%  -1.49% (p=0.000 n=10)
any_vs_actual                   1.031Ki ± 0%   1.031Ki ± 0%       ~ (p=1.000 n=10) ¹
all_CORS_headers_vs_actual      1.047Ki ± 0%   1.078Ki ± 0%  +2.99% (p=0.000 n=10)
single_vs_preflight               960.0 ± 0%     944.0 ± 0%  -1.67% (p=0.000 n=10)
multiple_vs_preflight             960.0 ± 0%     944.0 ± 0%  -1.67% (p=0.000 n=10)
pathological_vs_preflight         960.0 ± 0%     928.0 ± 0%  -3.33% (p=0.000 n=10)
many_vs_preflight                 960.0 ± 0%     928.0 ± 0%  -3.33% (p=0.000 n=10)
any_vs_preflight                  960.0 ± 0%     944.0 ± 0%  -1.67% (p=0.000 n=10)
ACRH_vs_preflight                 960.0 ± 0%     928.0 ± 0%  -3.33% (p=0.000 n=10)
all_CORS_headers_vs_preflight     960.0 ± 0%     928.0 ± 0%  -3.33% (p=0.000 n=10)
malicious_ACRH_vs_preflight       960.0 ± 0%     928.0 ± 0%  -3.33% (p=0.000 n=10)
geomean                          1003.3          989.7       -1.35%
¹ all samples are equal

                              │  rs-cors   │             jub0bs-cors              │
                              │ allocs/op  │ allocs/op   vs base                  │
single_vs_actual                10.00 ± 0%   11.00 ± 0%  +10.00% (p=0.000 n=10)
multiple_vs_actual              10.00 ± 0%   11.00 ± 0%  +10.00% (p=0.000 n=10)
pathological_vs_actual          12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
many_vs_actual                  12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
any_vs_actual                   10.00 ± 0%   11.00 ± 0%  +10.00% (p=0.000 n=10)
all_CORS_headers_vs_actual      10.00 ± 0%   12.00 ± 0%  +20.00% (p=0.000 n=10)
single_vs_preflight             7.000 ± 0%   7.000 ± 0%        ~ (p=1.000 n=10) ¹
multiple_vs_preflight           7.000 ± 0%   7.000 ± 0%        ~ (p=1.000 n=10) ¹
pathological_vs_preflight       9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
many_vs_preflight               9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
any_vs_preflight                7.000 ± 0%   7.000 ± 0%        ~ (p=1.000 n=10) ¹
ACRH_vs_preflight               9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
all_CORS_headers_vs_preflight   9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
malicious_ACRH_vs_preflight     9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
geomean                         9.157        8.549        -6.64%
¹ all samples are equal
```

For more details about the benchmark labeled "malicious_ACRH_vs_preflight",
see https://github.com/rs/cors/issues/170.
