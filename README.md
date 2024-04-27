# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.22.2) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.0), and
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
                              │   rs-cors   │             jub0bs-cors             │
                              │   sec/op    │   sec/op     vs base                │
single_vs_actual                649.1n ± 1%   680.7n ± 8%   +4.87% (p=0.000 n=10)
multiple_vs_actual              656.1n ± 1%   685.6n ± 2%   +4.49% (p=0.000 n=10)
pathological_vs_actual          717.2n ± 1%   752.5n ± 1%   +4.92% (p=0.002 n=10)
many_vs_actual                  684.7n ± 1%   648.1n ± 1%   -5.34% (p=0.000 n=10)
any_vs_actual                   640.3n ± 1%   637.5n ± 1%        ~ (p=0.271 n=10)
all_CORS_headers_vs_actual      682.5n ± 1%   752.9n ± 2%  +10.31% (p=0.000 n=10)
single_vs_preflight             519.0n ± 2%   465.9n ± 1%  -10.23% (p=0.000 n=10)
multiple_vs_preflight           521.5n ± 1%   471.4n ± 1%   -9.61% (p=0.000 n=10)
pathological_vs_preflight       524.8n ± 1%   523.1n ± 1%        ~ (p=0.739 n=10)
many_vs_preflight               488.8n ± 1%   426.7n ± 3%  -12.71% (p=0.000 n=10)
any_vs_preflight                511.8n ± 2%   458.2n ± 2%  -10.47% (p=0.000 n=10)
ACRH_vs_preflight               480.8n ± 2%   439.1n ± 2%   -8.68% (p=0.000 n=10)
all_CORS_headers_vs_preflight   492.8n ± 3%   452.8n ± 1%   -8.11% (p=0.000 n=10)
malicious_ACRH_vs_preflight     484.7n ± 1%   431.6n ± 2%  -10.94% (p=0.000 n=10)
geomean                         569.0n        546.3n        -4.00%

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
