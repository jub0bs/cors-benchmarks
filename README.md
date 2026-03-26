# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.26.1) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.13.6).

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
CORS/n=one/r=a/o=y-10                530.0n ± 1%   552.2n ± 1%   +4.19% (p=0.000 n=10)
CORS/n=one/r=a/o=n-10                535.1n ± 0%   528.2n ± 0%   -1.30% (p=0.000 n=10)
CORS/n=multiple/r=a/o=y-10           536.6n ± 0%   556.3n ± 0%   +3.68% (p=0.000 n=10)
CORS/n=multiple/r=a/o=n-10           544.5n ± 0%   536.3n ± 0%   -1.49% (p=0.000 n=10)
CORS/n=two/r=a/o=y-10                589.1n ± 1%   562.9n ± 0%   -4.45% (p=0.000 n=10)
CORS/n=two/r=a/o=n-10                589.4n ± 3%   551.4n ± 0%   -6.45% (p=0.000 n=10)
CORS/n=many/r=a/o=y-10              1049.0n ± 1%   575.7n ± 1%  -45.12% (p=0.000 n=10)
CORS/n=many/r=a/o=n-10               646.2n ± 0%   527.0n ± 0%  -18.45% (p=0.000 n=10)
CORS/n=all/r=a/o=y-10                534.7n ± 1%   529.9n ± 0%   -0.90% (p=0.000 n=10)
CORS/n=one/r=p/o=y-10                455.0n ± 1%   436.8n ± 0%   -4.02% (p=0.000 n=10)
CORS/n=one/r=p/o=n-10                384.1n ± 1%   107.5n ± 1%  -72.01% (p=0.000 n=10)
CORS/n=multiple/r=p/o=y-10           457.2n ± 0%   437.4n ± 0%   -4.31% (p=0.000 n=10)
CORS/n=multiple/r=p/o=n-10           391.5n ± 0%   113.5n ± 0%  -71.01% (p=0.000 n=10)
CORS/n=two/r=p/o=y-10                507.3n ± 0%   441.1n ± 1%  -13.04% (p=0.000 n=10)
CORS/n=two/r=p/o=n-10                438.5n ± 1%   110.5n ± 0%  -74.79% (p=0.000 n=10)
CORS/n=many/r=p/o=y-10               934.4n ± 1%   453.9n ± 0%  -51.43% (p=0.000 n=10)
CORS/n=many/r=p/o=n-10               494.9n ± 0%   107.3n ± 0%  -78.31% (p=0.000 n=10)
CORS/n=all/r=p/o=y-10                447.3n ± 0%   446.8n ± 0%        ~ (p=0.254 n=10)
CORS/n=all/r=p/o=y/m=evil_acrh-10    452.9n ± 0%   143.1n ± 0%  -68.41% (p=0.000 n=10)
geomean                              534.6n        343.7n       -35.71%

                                  │   rs-cors    │              jub0bs-cors               │
                                  │     B/op     │     B/op      vs base                  │
CORS/n=one/r=a/o=y-10               1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
CORS/n=one/r=a/o=n-10               1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
CORS/n=multiple/r=a/o=y-10          1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
CORS/n=multiple/r=a/o=n-10          1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
CORS/n=two/r=a/o=y-10               1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
CORS/n=two/r=a/o=n-10               1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
CORS/n=many/r=a/o=y-10              1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
CORS/n=many/r=a/o=n-10              1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
CORS/n=all/r=a/o=y-10               1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
CORS/n=one/r=p/o=y-10                 976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
CORS/n=one/r=p/o=n-10                 960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
CORS/n=multiple/r=p/o=y-10            976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
CORS/n=multiple/r=p/o=n-10            960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
CORS/n=two/r=p/o=y-10                 976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
CORS/n=two/r=p/o=n-10                 960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
CORS/n=many/r=p/o=y-10                976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
CORS/n=many/r=p/o=n-10                960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
CORS/n=all/r=p/o=y-10                 976.0 ± 0%     960.0 ± 0%   -1.64% (p=0.000 n=10)
CORS/n=all/r=p/o=y/m=evil_acrh-10     968.0 ± 0%     224.0 ± 0%  -76.86% (p=0.000 n=10)
geomean                              1016.4          673.9       -33.69%
¹ all samples are equal

                                  │  rs-cors   │             jub0bs-cors              │
                                  │ allocs/op  │ allocs/op   vs base                  │
CORS/n=one/r=a/o=y-10               12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
CORS/n=one/r=a/o=n-10               12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
CORS/n=multiple/r=a/o=y-10          12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
CORS/n=multiple/r=a/o=n-10          12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
CORS/n=two/r=a/o=y-10               12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
CORS/n=two/r=a/o=n-10               12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
CORS/n=many/r=a/o=y-10              12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
CORS/n=many/r=a/o=n-10              12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
CORS/n=all/r=a/o=y-10               12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
CORS/n=one/r=p/o=y-10               8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
CORS/n=one/r=p/o=n-10               9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
CORS/n=multiple/r=p/o=y-10          8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
CORS/n=multiple/r=p/o=n-10          9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
CORS/n=two/r=p/o=y-10               8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
CORS/n=two/r=p/o=n-10               9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
CORS/n=many/r=p/o=y-10              8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
CORS/n=many/r=p/o=n-10              9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
CORS/n=all/r=p/o=y-10               8.000 ± 0%   8.000 ± 0%        ~ (p=1.000 n=10) ¹
CORS/n=all/r=p/o=y/m=evil_acrh-10   9.000 ± 0%   5.000 ± 0%  -44.44% (p=0.000 n=10)
geomean                             9.999        7.626       -23.74%
¹ all samples are equal
```

Nomenclature:
- `n` indicates the number of allowed origins: `one` | `two` | `multiple` | `many` | `all`
- `r` indicates the type of request (actual or preflight): `a` | `p`
- `o` indicates whether the request's origin is allowed: `y` | `n`
- `m` indicates a case involving a malicious request

For more details about the benchmark labeled `m=evil_acrh`,
see https://github.com/rs/cors/issues/170.
