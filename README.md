# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.22.1) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.10.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.1.2).

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
                            │    rs-cors    │             jub0bs-cors             │
                            │    sec/op     │   sec/op     vs base                │
single_vs_actual                636.2n ± 1%   678.7n ± 1%   +6.68% (p=0.000 n=10)
multiple_vs_actual              646.2n ± 2%   678.9n ± 1%   +5.04% (p=0.000 n=10)
pathological_vs_actual          712.3n ± 1%   746.7n ± 1%   +4.83% (p=0.000 n=10)
many_vs_actual                  671.5n ± 2%   652.6n ± 2%   -2.81% (p=0.001 n=10)
any_vs_actual                   633.9n ± 2%   632.7n ± 1%        ~ (p=0.957 n=10)
single_vs_preflight             513.5n ± 1%   464.3n ± 2%   -9.59% (p=0.000 n=10)
multiple_vs_preflight           519.6n ± 2%   465.8n ± 1%  -10.37% (p=0.000 n=10)
pathological_vs_preflight       527.2n ± 1%   515.1n ± 1%   -2.31% (p=0.000 n=10)
many_vs_preflight               487.7n ± 1%   425.3n ± 3%  -12.79% (p=0.000 n=10)
any_vs_preflight                504.4n ± 2%   452.0n ± 1%  -10.40% (p=0.000 n=10)
ACRH_vs_preflight               520.3n ± 2%   427.1n ± 2%  -17.90% (p=0.000 n=10)
malicious_ACRH_vs_preflight   16989.5n ± 3%   427.9n ± 2%  -97.48% (p=0.000 n=10)
geomean                         761.7n        535.5n       -29.70%

                            │   rs-cors    │              jub0bs-cors               │
                            │     B/op     │     B/op      vs base                  │
single_vs_actual              1.031Ki ± 0%   1.047Ki ± 0%   +1.52% (p=0.000 n=10)
multiple_vs_actual            1.031Ki ± 0%   1.047Ki ± 0%   +1.52% (p=0.000 n=10)
pathological_vs_actual        1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
many_vs_actual                1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
any_vs_actual                 1.031Ki ± 0%   1.031Ki ± 0%        ~ (p=1.000 n=10) ¹
single_vs_preflight             960.0 ± 0%     944.0 ± 0%   -1.67% (p=0.000 n=10)
multiple_vs_preflight           960.0 ± 0%     944.0 ± 0%   -1.67% (p=0.000 n=10)
pathological_vs_preflight       960.0 ± 0%     928.0 ± 0%   -3.33% (p=0.000 n=10)
many_vs_preflight               960.0 ± 0%     928.0 ± 0%   -3.33% (p=0.000 n=10)
any_vs_preflight                960.0 ± 0%     944.0 ± 0%   -1.67% (p=0.000 n=10)
ACRH_vs_preflight               984.0 ± 0%     928.0 ± 0%   -5.69% (p=0.000 n=10)
malicious_ACRH_vs_preflight   37832.0 ± 0%     928.0 ± 0%  -97.55% (p=0.000 n=10)
geomean                       1.331Ki          986.0       -27.66%
¹ all samples are equal

                            │   rs-cors   │             jub0bs-cors              │
                            │  allocs/op  │ allocs/op   vs base                  │
single_vs_actual               10.00 ± 0%   11.00 ± 0%  +10.00% (p=0.000 n=10)
multiple_vs_actual             10.00 ± 0%   11.00 ± 0%  +10.00% (p=0.000 n=10)
pathological_vs_actual         12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
many_vs_actual                 12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
any_vs_actual                  10.00 ± 0%   11.00 ± 0%  +10.00% (p=0.000 n=10)
single_vs_preflight            7.000 ± 0%   7.000 ± 0%        ~ (p=1.000 n=10) ¹
multiple_vs_preflight          7.000 ± 0%   7.000 ± 0%        ~ (p=1.000 n=10) ¹
pathological_vs_preflight      9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
many_vs_preflight              9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
any_vs_preflight               7.000 ± 0%   7.000 ± 0%        ~ (p=1.000 n=10) ¹
ACRH_vs_preflight             10.000 ± 0%   7.000 ± 0%  -30.00% (p=0.000 n=10)
malicious_ACRH_vs_preflight   11.000 ± 0%   7.000 ± 0%  -36.36% (p=0.000 n=10)
geomean                        9.339        8.451        -9.51%
¹ all samples are equal
```

For more details about the benchmark labeled "malicious_ACRH_vs_preflight",
see https://github.com/rs/cors/issues/170.
