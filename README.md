# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.26.0) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.11.0).

## Running the benchmarks

Run the following commands in your shell (preferably on an idle machine):

```shell
git clone https://github.com/jub0bs/cors-benchmarks
cd cors-benchmarks
go test -run ^$ -bench . -count 10 > bench.out
benchstat -col "/mw@(rs-cors jub0bs-cors)" bench.out && rm bench.out
```

## Results

```text
goos: darwin
goarch: arm64
pkg: github.com/jub0bs/cors-benchmarks
cpu: Apple M4
                                           │   rs-cors    │             jub0bs-cors             │
                                           │    sec/op    │   sec/op     vs base                │
Middleware/nb=one/req=a/o=y-10                529.4n ± 1%   553.6n ± 0%   +4.57% (p=0.000 n=10)
Middleware/nb=one/req=a/o=n-10                528.6n ± 1%   524.3n ± 0%   -0.81% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=y-10           531.5n ± 1%   559.8n ± 0%   +5.32% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=n-10           539.5n ± 0%   537.0n ± 3%        ~ (p=0.517 n=10)
Middleware/nb=two/req=a/o=y-10                596.2n ± 0%   588.7n ± 0%   -1.25% (p=0.002 n=10)
Middleware/nb=two/req=a/o=n-10                580.0n ± 0%   559.1n ± 1%   -3.60% (p=0.000 n=10)
Middleware/nb=many/req=a/o=y-10              1034.5n ± 0%   595.4n ± 3%  -42.45% (p=0.000 n=10)
Middleware/nb=many/req=a/o=n-10               654.3n ± 0%   551.4n ± 0%  -15.73% (p=0.000 n=10)
Middleware/nb=all/req=a/o=y-10                527.7n ± 3%   538.4n ± 0%        ~ (p=0.138 n=10)
Middleware/nb=one/req=p/o=y-10                465.2n ± 0%   447.2n ± 1%   -3.87% (p=0.000 n=10)
Middleware/nb=one/req=p/o=n-10                381.8n ± 0%   110.3n ± 0%  -71.11% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=y-10           454.1n ± 0%   442.4n ± 1%   -2.55% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=n-10           389.5n ± 0%   117.1n ± 1%  -69.94% (p=0.000 n=10)
Middleware/nb=two/req=p/o=y-10                504.3n ± 1%   473.4n ± 0%   -6.11% (p=0.000 n=10)
Middleware/nb=two/req=p/o=n-10                436.1n ± 0%   138.0n ± 1%  -68.34% (p=0.000 n=10)
Middleware/nb=many/req=p/o=y-10               926.1n ± 0%   456.2n ± 0%  -50.73% (p=0.000 n=10)
Middleware/nb=many/req=p/o=n-10               491.6n ± 0%   114.0n ± 0%  -76.81% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y-10                445.6n ± 0%   426.6n ± 0%   -4.26% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y/m=evil_acrh-10    450.6n ± 0%   126.6n ± 1%  -71.90% (p=0.000 n=10)
geomean                                       532.3n        351.7n       -33.93%

                                           │   rs-cors    │              jub0bs-cors               │
                                           │     B/op     │     B/op      vs base                  │
Middleware/nb=one/req=a/o=y-10               1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
Middleware/nb=one/req=a/o=n-10               1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=y-10          1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
Middleware/nb=multiple/req=a/o=n-10          1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
Middleware/nb=two/req=a/o=y-10               1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
Middleware/nb=two/req=a/o=n-10               1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
Middleware/nb=many/req=a/o=y-10              1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
Middleware/nb=many/req=a/o=n-10              1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
Middleware/nb=all/req=a/o=y-10               1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
Middleware/nb=one/req=p/o=y-10                 976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
Middleware/nb=one/req=p/o=n-10                 960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=y-10            976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=n-10            960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
Middleware/nb=two/req=p/o=y-10                 976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
Middleware/nb=two/req=p/o=n-10                 960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
Middleware/nb=many/req=p/o=y-10                976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
Middleware/nb=many/req=p/o=n-10                960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y-10                 976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y/m=evil_acrh-10     968.0 ± 0%     208.0 ± 0%  -78.51% (p=0.000 n=10)
geomean                                       1016.4          670.7       -34.01%
¹ all samples are equal

                                           │  rs-cors   │            jub0bs-cors             │
                                           │ allocs/op  │ allocs/op   vs base                │
Middleware/nb=one/req=a/o=y-10               12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=one/req=a/o=n-10               12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=y-10          12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=n-10          12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=two/req=a/o=y-10               12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=two/req=a/o=n-10               12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=many/req=a/o=y-10              12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=many/req=a/o=n-10              12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=all/req=a/o=y-10               12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=one/req=p/o=y-10               8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
Middleware/nb=one/req=p/o=n-10               9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=y-10          8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=n-10          9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
Middleware/nb=two/req=p/o=y-10               8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
Middleware/nb=two/req=p/o=n-10               9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
Middleware/nb=many/req=p/o=y-10              8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
Middleware/nb=many/req=p/o=n-10              9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y-10               8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y/m=evil_acrh-10   9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
geomean                                      9.999        7.484       -25.16%
```

Nomenclature:
- `nb` indicates the number of allowed origins: `one` | `two` | `multiple` | `many` | `all`
- `req` indicates the type of request (actual or preflight): `a` | `p`
- `o` indicates whether the request's origin is allowed: `y` | `n`
- `m` indicates a case involving a malicious request

For more details about the benchmark labeled `m=evil_acrh`,
see https://github.com/rs/cors/issues/170.
