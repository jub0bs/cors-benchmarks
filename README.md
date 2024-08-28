# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.23.0) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.0), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.3.0).

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
                              │   rs-cors   │             jub0bs-cors              │
                              │   sec/op    │    sec/op     vs base                │
single_vs_actual                631.9n ± 2%   710.8n ±  2%  +12.48% (p=0.000 n=10)
multiple_vs_actual              673.5n ± 1%   727.3n ±  1%   +8.00% (p=0.000 n=10)
pathological_vs_actual          753.4n ± 1%   839.2n ±  1%  +11.40% (p=0.000 n=10)
many_vs_actual                  716.4n ± 1%   703.2n ±  0%   -1.84% (p=0.000 n=10)
any_vs_actual                   675.5n ± 1%   679.6n ±  1%        ~ (p=0.052 n=10)
all_CORS_headers_vs_actual      721.0n ± 3%   819.6n ±  1%  +13.68% (p=0.000 n=10)
single_vs_preflight             550.4n ± 2%   528.8n ±  2%   -3.92% (p=0.000 n=10)
multiple_vs_preflight           552.0n ± 1%   534.2n ±  2%   -3.22% (p=0.000 n=10)
pathological_vs_preflight       568.2n ± 1%   603.0n ±  1%   +6.12% (p=0.000 n=10)
many_vs_preflight               522.8n ± 1%   468.5n ±  4%  -10.39% (p=0.000 n=10)
any_vs_preflight                543.2n ± 3%   557.2n ± 13%        ~ (p=0.739 n=10)
ACRH_vs_preflight               529.2n ± 8%   488.6n ±  2%   -7.67% (p=0.000 n=10)
all_CORS_headers_vs_preflight   541.3n ± 3%   518.8n ±  3%   -4.16% (p=0.002 n=10)
malicious_ACRH_vs_preflight     503.5n ± 2%   501.9n ±  2%        ~ (p=0.869 n=10)
geomean                         600.4n        608.8n         +1.40%

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
