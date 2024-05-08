# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.22.2) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.0), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.2.0).

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
single_vs_actual                643.7n ± 1%   693.8n ± 1%   +7.78% (p=0.000 n=10)
multiple_vs_actual              647.4n ± 1%   690.8n ± 1%   +6.70% (p=0.000 n=10)
pathological_vs_actual          716.0n ± 0%   809.0n ± 1%  +12.99% (p=0.000 n=10)
many_vs_actual                  676.8n ± 1%   661.3n ± 1%   -2.28% (p=0.000 n=10)
any_vs_actual                   632.9n ± 1%   643.3n ± 1%   +1.65% (p=0.001 n=10)
all_CORS_headers_vs_actual      677.4n ± 1%   764.3n ± 1%  +12.83% (p=0.000 n=10)
single_vs_preflight             517.1n ± 1%   477.1n ± 1%   -7.75% (p=0.000 n=10)
multiple_vs_preflight           520.2n ± 1%   479.1n ± 1%   -7.92% (p=0.000 n=10)
pathological_vs_preflight       528.8n ± 1%   578.5n ± 0%   +9.41% (p=0.000 n=10)
many_vs_preflight               489.7n ± 1%   434.4n ± 2%  -11.27% (p=0.000 n=10)
any_vs_preflight                508.8n ± 0%   465.0n ± 2%   -8.61% (p=0.000 n=10)
ACRH_vs_preflight               477.8n ± 2%   440.9n ± 1%   -7.72% (p=0.000 n=10)
all_CORS_headers_vs_preflight   490.9n ± 2%   457.6n ± 1%   -6.77% (p=0.000 n=10)
malicious_ACRH_vs_preflight     482.4n ± 1%   440.3n ± 4%   -8.73% (p=0.000 n=10)
geomean                         566.1n        560.1n        -1.06%

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
