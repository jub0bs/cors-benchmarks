# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.26.0) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.13.2).

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
Middleware/nb=one/req=a/o=y-10                532.2n ± 2%   554.5n ± 0%   +4.18% (p=0.000 n=10)
Middleware/nb=one/req=a/o=n-10                533.0n ± 1%   525.2n ± 1%   -1.47% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=y-10           536.5n ± 1%   561.1n ± 0%   +4.58% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=n-10           542.3n ± 0%   540.4n ± 0%   -0.35% (p=0.005 n=10)
Middleware/nb=two/req=a/o=y-10                584.3n ± 0%   590.3n ± 1%   +1.03% (p=0.000 n=10)
Middleware/nb=two/req=a/o=n-10                583.8n ± 0%   561.0n ± 0%   -3.91% (p=0.000 n=10)
Middleware/nb=many/req=a/o=y-10              1035.0n ± 1%   578.8n ± 1%  -44.08% (p=0.000 n=10)
Middleware/nb=many/req=a/o=n-10               641.5n ± 0%   532.2n ± 0%  -17.04% (p=0.000 n=10)
Middleware/nb=all/req=a/o=y-10                544.2n ± 2%   529.2n ± 0%   -2.76% (p=0.000 n=10)
Middleware/nb=one/req=p/o=y-10                456.2n ± 0%   439.7n ± 2%   -3.61% (p=0.000 n=10)
Middleware/nb=one/req=p/o=n-10                381.8n ± 2%   107.9n ± 1%  -71.74% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=y-10           457.8n ± 0%   443.2n ± 0%   -3.19% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=n-10           391.4n ± 0%   117.8n ± 1%  -69.92% (p=0.000 n=10)
Middleware/nb=two/req=p/o=y-10                506.4n ± 0%   471.8n ± 0%   -6.82% (p=0.000 n=10)
Middleware/nb=two/req=p/o=n-10                437.2n ± 0%   135.3n ± 1%  -69.04% (p=0.000 n=10)
Middleware/nb=many/req=p/o=y-10               936.3n ± 1%   454.2n ± 1%  -51.48% (p=0.000 n=10)
Middleware/nb=many/req=p/o=n-10               495.9n ± 0%   110.8n ± 1%  -77.66% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y-10                449.2n ± 0%   427.8n ± 0%   -4.78% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y/m=evil_acrh-10    454.2n ± 0%   126.2n ± 0%  -72.22% (p=0.000 n=10)
geomean                                       534.0n        348.9n       -34.66%

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
