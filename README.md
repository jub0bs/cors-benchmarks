# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.27.3) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v1.3.0).

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
CORS/n=one/r=a/o=y-10                514.9n ± 1%   538.3n ± 1%   +4.55% (p=0.001 n=10)
CORS/n=one/r=a/o=n-10                521.3n ± 1%   511.3n ± 1%   -1.92% (p=0.000 n=10)
CORS/n=multiple/r=a/o=y-10           522.6n ± 2%   541.1n ± 0%   +3.54% (p=0.001 n=10)
CORS/n=multiple/r=a/o=n-10           529.4n ± 1%   521.2n ± 0%   -1.55% (p=0.000 n=10)
CORS/n=two/r=a/o=y-10                572.4n ± 1%   546.2n ± 0%   -4.56% (p=0.000 n=10)
CORS/n=two/r=a/o=n-10                570.5n ± 1%   520.5n ± 1%   -8.76% (p=0.000 n=10)
CORS/n=many/r=a/o=y-10              1020.0n ± 0%   562.6n ± 2%  -44.84% (p=0.000 n=10)
CORS/n=many/r=a/o=n-10               626.8n ± 1%   513.1n ± 1%  -18.15% (p=0.000 n=10)
CORS/n=all/r=a/o=y-10                516.9n ± 1%   501.2n ± 1%   -3.05% (p=0.000 n=10)
CORS/n=one/r=p/o=y-10                441.3n ± 1%   382.6n ± 1%  -13.30% (p=0.000 n=10)
CORS/n=one/r=p/o=n-10               372.70n ± 0%   98.64n ± 1%  -73.53% (p=0.000 n=10)
CORS/n=multiple/r=p/o=y-10           444.0n ± 0%   384.0n ± 1%  -13.51% (p=0.000 n=10)
CORS/n=multiple/r=p/o=n-10           380.3n ± 0%   105.0n ± 1%  -72.38% (p=0.000 n=10)
CORS/n=two/r=p/o=y-10                493.8n ± 1%   388.2n ± 0%  -21.38% (p=0.000 n=10)
CORS/n=two/r=p/o=n-10                426.6n ± 0%   103.4n ± 1%  -75.76% (p=0.000 n=10)
CORS/n=many/r=p/o=y-10               926.4n ± 1%   405.1n ± 0%  -56.28% (p=0.000 n=10)
CORS/n=many/r=p/o=n-10              479.95n ± 1%   99.70n ± 0%  -79.23% (p=0.000 n=10)
CORS/n=all/r=p/o=y-10                433.9n ± 1%   382.6n ± 0%  -11.83% (p=0.000 n=10)
CORS/n=all/r=p/o=y/m=evil_acrh-10    439.6n ± 1%   116.5n ± 1%  -73.50% (p=0.000 n=10)
geomean                              519.7n        318.0n       -38.82%

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
