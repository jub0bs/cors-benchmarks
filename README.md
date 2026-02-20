# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.26.0) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.13.0).

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
Middleware/nb=one/req=a/o=y-10                524.6n ± 1%   547.8n ± 0%   +4.43% (p=0.000 n=10)
Middleware/nb=one/req=a/o=n-10                528.1n ± 1%   520.5n ± 0%   -1.42% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=y-10           530.9n ± 1%   559.2n ± 1%   +5.32% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=n-10           537.1n ± 0%   536.5n ± 0%        ~ (p=0.138 n=10)
Middleware/nb=two/req=a/o=y-10                581.1n ± 3%   587.6n ± 1%   +1.13% (p=0.014 n=10)
Middleware/nb=two/req=a/o=n-10                580.7n ± 0%   556.9n ± 0%   -4.10% (p=0.000 n=10)
Middleware/nb=many/req=a/o=y-10              1032.5n ± 0%   573.5n ± 0%  -44.45% (p=0.000 n=10)
Middleware/nb=many/req=a/o=n-10               638.8n ± 0%   533.2n ± 0%  -16.53% (p=0.000 n=10)
Middleware/nb=all/req=a/o=y-10                541.8n ± 2%   538.0n ± 0%        ~ (p=0.052 n=10)
Middleware/nb=one/req=p/o=y-10                463.4n ± 0%   449.1n ± 3%   -3.10% (p=0.000 n=10)
Middleware/nb=one/req=p/o=n-10                380.2n ± 0%   107.8n ± 1%  -71.65% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=y-10           452.8n ± 0%   440.6n ± 0%   -2.69% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=n-10           388.1n ± 0%   116.9n ± 0%  -69.87% (p=0.000 n=10)
Middleware/nb=two/req=p/o=y-10                503.6n ± 0%   467.7n ± 1%   -7.12% (p=0.000 n=10)
Middleware/nb=two/req=p/o=n-10                435.3n ± 0%   135.2n ± 1%  -68.94% (p=0.000 n=10)
Middleware/nb=many/req=p/o=y-10               928.4n ± 0%   451.9n ± 1%  -51.33% (p=0.000 n=10)
Middleware/nb=many/req=p/o=n-10               493.8n ± 0%   110.8n ± 0%  -77.55% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y-10                456.6n ± 3%   426.1n ± 0%   -6.70% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y/m=evil_acrh-10    448.9n ± 0%   126.5n ± 0%  -71.82% (p=0.000 n=10)
geomean                                       531.4n        348.1n       -34.50%

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
