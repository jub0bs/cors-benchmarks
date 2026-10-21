# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.27.3) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v1.1.3).

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
CORS/n=one/r=a/o=y-10                516.7n ± 0%   534.9n ± 1%   +3.50% (p=0.000 n=10)
CORS/n=one/r=a/o=n-10                517.9n ± 0%   510.5n ± 1%   -1.43% (p=0.000 n=10)
CORS/n=multiple/r=a/o=y-10           519.6n ± 0%   539.4n ± 1%   +3.80% (p=0.000 n=10)
CORS/n=multiple/r=a/o=n-10           527.6n ± 0%   518.8n ± 1%   -1.69% (p=0.002 n=10)
CORS/n=two/r=a/o=y-10                569.9n ± 0%   543.7n ± 1%   -4.60% (p=0.000 n=10)
CORS/n=two/r=a/o=n-10                569.4n ± 0%   517.0n ± 0%   -9.19% (p=0.000 n=10)
CORS/n=many/r=a/o=y-10              1021.5n ± 0%   558.5n ± 0%  -45.32% (p=0.000 n=10)
CORS/n=many/r=a/o=n-10               626.2n ± 0%   509.6n ± 0%  -18.63% (p=0.000 n=10)
CORS/n=all/r=a/o=y-10                523.6n ± 2%   524.4n ± 0%        ~ (p=0.754 n=10)
CORS/n=one/r=p/o=y-10                461.8n ± 1%   401.1n ± 0%  -13.14% (p=0.000 n=10)
CORS/n=one/r=p/o=n-10                383.6n ± 2%   103.3n ± 0%  -73.07% (p=0.000 n=10)
CORS/n=multiple/r=p/o=y-10           464.1n ± 0%   402.7n ± 1%  -13.24% (p=0.000 n=10)
CORS/n=multiple/r=p/o=n-10           393.6n ± 0%   109.1n ± 4%  -72.29% (p=0.000 n=10)
CORS/n=two/r=p/o=y-10                490.1n ± 0%   385.8n ± 0%  -21.27% (p=0.000 n=10)
CORS/n=two/r=p/o=n-10                423.0n ± 0%   102.9n ± 0%  -75.67% (p=0.000 n=10)
CORS/n=many/r=p/o=y-10               919.4n ± 1%   403.9n ± 0%  -56.06% (p=0.000 n=10)
CORS/n=many/r=p/o=n-10              478.75n ± 0%   98.83n ± 1%  -79.36% (p=0.000 n=10)
CORS/n=all/r=p/o=y-10                431.9n ± 0%   379.8n ± 0%  -12.06% (p=0.000 n=10)
CORS/n=all/r=p/o=y/m=evil_acrh-10    436.1n ± 1%   116.1n ± 1%  -73.37% (p=0.000 n=10)
geomean                              522.7n        320.5n       -38.69%

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
