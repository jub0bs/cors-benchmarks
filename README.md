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
single_vs_actual                  643.1n ± 1%   679.5n ± 2%   +5.65% (p=0.000 n=10)
multiple_vs_actual                645.6n ± 1%   677.9n ± 1%   +5.00% (p=0.000 n=10)
pathological_vs_actual            711.3n ± 1%   750.3n ± 1%   +5.47% (p=0.000 n=10)
many_vs_actual                    679.8n ± 1%   647.4n ± 1%   -4.77% (p=0.000 n=10)
any_vs_actual                     637.4n ± 1%   635.7n ± 1%        ~ (p=0.565 n=10)
all_CORS_headers_vs_actual        681.5n ± 1%   755.4n ± 1%  +10.84% (p=0.000 n=10)
single_vs_preflight               523.6n ± 2%   468.1n ± 3%  -10.59% (p=0.000 n=10)
multiple_vs_preflight             529.5n ± 2%   471.9n ± 2%  -10.88% (p=0.000 n=10)
pathological_vs_preflight         532.4n ± 2%   524.7n ± 2%        ~ (p=0.127 n=10)
many_vs_preflight                 496.3n ± 2%   425.0n ± 2%  -14.37% (p=0.000 n=10)
any_vs_preflight                  511.8n ± 2%   467.8n ± 3%   -8.60% (p=0.000 n=10)
ACRH_vs_preflight                 532.1n ± 2%   435.1n ± 4%  -18.22% (p=0.000 n=10)
all_CORS_headers_vs_preflight     548.0n ± 4%   459.6n ± 1%  -16.14% (p=0.000 n=10)
malicious_ACRH_vs_preflight     17238.0n ± 3%   438.2n ± 5%  -97.46% (p=0.000 n=10)
geomean                           745.7n        547.4n       -26.59%

                              │   rs-cors    │              jub0bs-cors               │
                              │     B/op     │     B/op      vs base                  │
single_vs_actual                1.031Ki ± 0%   1.047Ki ± 0%   +1.52% (p=0.000 n=10)
multiple_vs_actual              1.031Ki ± 0%   1.047Ki ± 0%   +1.52% (p=0.000 n=10)
pathological_vs_actual          1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
many_vs_actual                  1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
any_vs_actual                   1.031Ki ± 0%   1.031Ki ± 0%        ~ (p=1.000 n=10) ¹
all_CORS_headers_vs_actual      1.047Ki ± 0%   1.078Ki ± 0%   +2.99% (p=0.000 n=10)
single_vs_preflight               960.0 ± 0%     944.0 ± 0%   -1.67% (p=0.000 n=10)
multiple_vs_preflight             960.0 ± 0%     944.0 ± 0%   -1.67% (p=0.000 n=10)
pathological_vs_preflight         960.0 ± 0%     928.0 ± 0%   -3.33% (p=0.000 n=10)
many_vs_preflight                 960.0 ± 0%     928.0 ± 0%   -3.33% (p=0.000 n=10)
any_vs_preflight                  960.0 ± 0%     944.0 ± 0%   -1.67% (p=0.000 n=10)
ACRH_vs_preflight                 984.0 ± 0%     928.0 ± 0%   -5.69% (p=0.000 n=10)
all_CORS_headers_vs_preflight     984.0 ± 0%     928.0 ± 0%   -5.69% (p=0.000 n=10)
malicious_ACRH_vs_preflight     37832.0 ± 0%     928.0 ± 0%  -97.55% (p=0.000 n=10)
geomean                         1.278Ki          989.7       -24.39%
¹ all samples are equal

                              │   rs-cors   │             jub0bs-cors              │
                              │  allocs/op  │ allocs/op   vs base                  │
single_vs_actual                 10.00 ± 0%   11.00 ± 0%  +10.00% (p=0.000 n=10)
multiple_vs_actual               10.00 ± 0%   11.00 ± 0%  +10.00% (p=0.000 n=10)
pathological_vs_actual           12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
many_vs_actual                   12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
any_vs_actual                    10.00 ± 0%   11.00 ± 0%  +10.00% (p=0.000 n=10)
all_CORS_headers_vs_actual       10.00 ± 0%   12.00 ± 0%  +20.00% (p=0.000 n=10)
single_vs_preflight              7.000 ± 0%   7.000 ± 0%        ~ (p=1.000 n=10) ¹
multiple_vs_preflight            7.000 ± 0%   7.000 ± 0%        ~ (p=1.000 n=10) ¹
pathological_vs_preflight        9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
many_vs_preflight                9.000 ± 0%   7.000 ± 0%  -22.22% (p=0.000 n=10)
any_vs_preflight                 7.000 ± 0%   7.000 ± 0%        ~ (p=1.000 n=10) ¹
ACRH_vs_preflight               10.000 ± 0%   7.000 ± 0%  -30.00% (p=0.000 n=10)
all_CORS_headers_vs_preflight   10.000 ± 0%   7.000 ± 0%  -30.00% (p=0.000 n=10)
malicious_ACRH_vs_preflight     11.000 ± 0%   7.000 ± 0%  -36.36% (p=0.000 n=10)
geomean                          9.431        8.549        -9.35%
¹ all samples are equal
```

For more details about the benchmark labeled "malicious_ACRH_vs_preflight",
see https://github.com/rs/cors/issues/170.
