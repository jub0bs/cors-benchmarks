# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.22.1) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.10.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.1.1).

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
single_vs_actual                639.9n ± 1%   680.8n ± 2%   +6.38% (p=0.000 n=10)
multiple_vs_actual              644.9n ± 1%   685.0n ± 1%   +6.23% (p=0.000 n=10)
pathological_vs_actual          714.0n ± 1%   750.6n ± 1%   +5.13% (p=0.000 n=10)
many_vs_actual                  677.4n ± 1%   649.0n ± 1%   -4.20% (p=0.000 n=10)
any_vs_actual                   636.5n ± 1%   635.2n ± 2%        ~ (p=1.000 n=10)
single_vs_preflight             519.2n ± 0%   464.6n ± 1%  -10.52% (p=0.000 n=10)
multiple_vs_preflight           523.2n ± 1%   467.9n ± 1%  -10.57% (p=0.000 n=10)
pathological_vs_preflight       530.7n ± 1%   513.6n ± 1%   -3.22% (p=0.000 n=10)
many_vs_preflight               489.1n ± 1%   424.0n ± 2%  -13.31% (p=0.000 n=10)
any_vs_preflight                509.4n ± 1%   451.2n ± 1%  -11.42% (p=0.000 n=10)
ACRH_vs_preflight               527.2n ± 2%   425.9n ± 2%  -19.21% (p=0.000 n=10)
malicious_ACRH_vs_preflight   16589.5n ± 1%   430.9n ± 2%  -97.40% (p=0.000 n=10)
geomean                         764.6n        536.3n       -29.86%

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
