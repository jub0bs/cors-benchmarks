# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.26.0) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.13.3).

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
Middleware/nb=one/req=a/o=y-10                526.2n ± 1%   550.6n ± 1%   +4.65% (p=0.000 n=10)
Middleware/nb=one/req=a/o=n-10                530.7n ± 0%   523.5n ± 0%   -1.37% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=y-10           533.5n ± 0%   556.2n ± 1%   +4.26% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=n-10           539.6n ± 0%   531.9n ± 1%   -1.44% (p=0.000 n=10)
Middleware/nb=two/req=a/o=y-10                582.2n ± 1%   594.5n ± 0%   +2.10% (p=0.000 n=10)
Middleware/nb=two/req=a/o=n-10                580.3n ± 0%   569.9n ± 0%   -1.79% (p=0.000 n=10)
Middleware/nb=many/req=a/o=y-10              1036.0n ± 0%   570.9n ± 0%  -44.89% (p=0.000 n=10)
Middleware/nb=many/req=a/o=n-10               639.6n ± 0%   529.5n ± 1%  -17.21% (p=0.000 n=10)
Middleware/nb=all/req=a/o=y-10                528.2n ± 0%   523.2n ± 0%   -0.95% (p=0.000 n=10)
Middleware/nb=one/req=p/o=y-10                449.8n ± 0%   432.8n ± 1%   -3.78% (p=0.000 n=10)
Middleware/nb=one/req=p/o=n-10                382.5n ± 1%   106.4n ± 1%  -72.19% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=y-10           452.5n ± 0%   434.6n ± 0%   -3.97% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=n-10           388.5n ± 1%   113.0n ± 0%  -70.93% (p=0.000 n=10)
Middleware/nb=two/req=p/o=y-10                500.6n ± 0%   476.4n ± 0%   -4.82% (p=0.000 n=10)
Middleware/nb=two/req=p/o=n-10                434.5n ± 0%   143.7n ± 1%  -66.93% (p=0.000 n=10)
Middleware/nb=many/req=p/o=y-10               929.9n ± 1%   453.8n ± 1%  -51.19% (p=0.000 n=10)
Middleware/nb=many/req=p/o=n-10               492.6n ± 0%   108.0n ± 0%  -78.08% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y-10                444.8n ± 0%   423.6n ± 0%   -4.77% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y/m=evil_acrh-10    449.4n ± 0%   124.5n ± 0%  -72.31% (p=0.000 n=10)
geomean                                       529.8n        346.8n       -34.54%

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
