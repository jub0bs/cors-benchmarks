# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.26.1) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.13.4).

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
Middleware/nb=one/req=a/o=y-10                526.8n ± 1%   548.4n ± 0%   +4.10% (p=0.000 n=10)
Middleware/nb=one/req=a/o=n-10                528.9n ± 1%   538.9n ± 1%   +1.90% (p=0.001 n=10)
Middleware/nb=multiple/req=a/o=y-10           541.9n ± 1%   555.5n ± 1%   +2.52% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=n-10           539.4n ± 0%   531.9n ± 0%   -1.39% (p=0.000 n=10)
Middleware/nb=two/req=a/o=y-10                582.8n ± 0%   558.0n ± 0%   -4.25% (p=0.000 n=10)
Middleware/nb=two/req=a/o=n-10                579.9n ± 0%   529.7n ± 0%   -8.66% (p=0.000 n=10)
Middleware/nb=many/req=a/o=y-10              1038.0n ± 2%   571.1n ± 2%  -44.98% (p=0.000 n=10)
Middleware/nb=many/req=a/o=n-10               637.6n ± 0%   527.6n ± 0%  -17.25% (p=0.000 n=10)
Middleware/nb=all/req=a/o=y-10                527.9n ± 1%   523.8n ± 0%   -0.78% (p=0.002 n=10)
Middleware/nb=one/req=p/o=y-10                451.4n ± 0%   434.7n ± 0%   -3.70% (p=0.000 n=10)
Middleware/nb=one/req=p/o=n-10                383.0n ± 2%   106.3n ± 1%  -72.26% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=y-10           453.9n ± 0%   435.1n ± 0%   -4.14% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=n-10           389.4n ± 0%   112.9n ± 0%  -71.01% (p=0.000 n=10)
Middleware/nb=two/req=p/o=y-10                501.9n ± 0%   439.6n ± 0%  -12.40% (p=0.000 n=10)
Middleware/nb=two/req=p/o=n-10                436.4n ± 0%   110.9n ± 0%  -74.58% (p=0.000 n=10)
Middleware/nb=many/req=p/o=y-10               929.1n ± 1%   457.0n ± 0%  -50.81% (p=0.000 n=10)
Middleware/nb=many/req=p/o=n-10               490.6n ± 0%   107.3n ± 1%  -78.12% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y-10                445.3n ± 0%   424.4n ± 0%   -4.69% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y/m=evil_acrh-10    448.9n ± 0%   124.9n ± 0%  -72.18% (p=0.000 n=10)
geomean                                       530.4n        338.8n       -36.13%

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
