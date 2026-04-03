# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.26.1) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.14.0).

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
CORS/n=one/r=a/o=y-10                531.5n ± 2%   549.3n ± 0%   +3.35% (p=0.000 n=10)
CORS/n=one/r=a/o=n-10                530.8n ± 1%   524.6n ± 0%   -1.18% (p=0.001 n=10)
CORS/n=multiple/r=a/o=y-10           533.1n ± 0%   553.7n ± 0%   +3.85% (p=0.000 n=10)
CORS/n=multiple/r=a/o=n-10           541.1n ± 0%   533.1n ± 0%   -1.49% (p=0.000 n=10)
CORS/n=two/r=a/o=y-10                584.8n ± 0%   559.7n ± 1%   -4.28% (p=0.000 n=10)
CORS/n=two/r=a/o=n-10                588.2n ± 2%   533.7n ± 1%   -9.27% (p=0.000 n=10)
CORS/n=many/r=a/o=y-10              1039.0n ± 1%   570.9n ± 0%  -45.06% (p=0.000 n=10)
CORS/n=many/r=a/o=n-10               641.0n ± 0%   524.1n ± 0%  -18.24% (p=0.000 n=10)
CORS/n=all/r=a/o=y-10                538.6n ± 2%   526.1n ± 0%   -2.33% (p=0.000 n=10)
CORS/n=one/r=p/o=y-10                450.8n ± 0%   433.6n ± 2%   -3.82% (p=0.000 n=10)
CORS/n=one/r=p/o=n-10                381.8n ± 0%   106.6n ± 0%  -72.09% (p=0.000 n=10)
CORS/n=multiple/r=p/o=y-10           461.7n ± 2%   448.1n ± 1%   -2.95% (p=0.000 n=10)
CORS/n=multiple/r=p/o=n-10           388.8n ± 0%   112.6n ± 1%  -71.05% (p=0.000 n=10)
CORS/n=two/r=p/o=y-10                503.6n ± 0%   439.6n ± 6%  -12.71% (p=0.000 n=10)
CORS/n=two/r=p/o=n-10                436.6n ± 0%   110.9n ± 1%  -74.60% (p=0.000 n=10)
CORS/n=many/r=p/o=y-10               930.6n ± 1%   453.8n ± 2%  -51.24% (p=0.000 n=10)
CORS/n=many/r=p/o=n-10               491.2n ± 0%   107.2n ± 0%  -78.17% (p=0.000 n=10)
CORS/n=all/r=p/o=y-10                445.6n ± 1%   445.9n ± 1%        ~ (p=0.956 n=10)
CORS/n=all/r=p/o=y/m=evil_acrh-10    450.6n ± 0%   143.1n ± 0%  -68.25% (p=0.000 n=10)
geomean                              532.3n        342.1n       -35.73%

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
