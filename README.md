# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.26.1) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.13.5).

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
CORS/n=one/r=a/o=y-10                523.7n ± 0%   545.5n ± 1%   +4.16% (p=0.000 n=10)
CORS/n=one/r=a/o=n-10                535.2n ± 2%   524.5n ± 1%   -2.00% (p=0.003 n=10)
CORS/n=multiple/r=a/o=y-10           541.1n ± 2%   552.4n ± 2%   +2.10% (p=0.000 n=10)
CORS/n=multiple/r=a/o=n-10           538.7n ± 0%   530.7n ± 0%   -1.48% (p=0.001 n=10)
CORS/n=two/r=a/o=y-10                583.8n ± 0%   556.6n ± 0%   -4.64% (p=0.000 n=10)
CORS/n=two/r=a/o=n-10                582.2n ± 0%   529.0n ± 0%   -9.14% (p=0.000 n=10)
CORS/n=many/r=a/o=y-10              1039.5n ± 1%   570.0n ± 0%  -45.17% (p=0.000 n=10)
CORS/n=many/r=a/o=n-10               638.2n ± 0%   521.6n ± 0%  -18.28% (p=0.000 n=10)
CORS/n=all/r=a/o=y-10                529.0n ± 0%   523.2n ± 0%   -1.09% (p=0.000 n=10)
CORS/n=one/r=p/o=y-10                448.9n ± 0%   427.8n ± 0%   -4.71% (p=0.000 n=10)
CORS/n=one/r=p/o=n-10                379.2n ± 0%   106.5n ± 0%  -71.93% (p=0.000 n=10)
CORS/n=multiple/r=p/o=y-10           450.1n ± 1%   431.2n ± 0%   -4.20% (p=0.000 n=10)
CORS/n=multiple/r=p/o=n-10           387.3n ± 1%   112.6n ± 1%  -70.93% (p=0.000 n=10)
CORS/n=two/r=p/o=y-10                515.5n ± 0%   447.1n ± 0%  -13.27% (p=0.000 n=10)
CORS/n=two/r=p/o=n-10                434.1n ± 1%   110.4n ± 1%  -74.57% (p=0.000 n=10)
CORS/n=many/r=p/o=y-10               927.5n ± 1%   450.5n ± 0%  -51.43% (p=0.000 n=10)
CORS/n=many/r=p/o=n-10               489.5n ± 0%   106.9n ± 0%  -78.16% (p=0.000 n=10)
CORS/n=all/r=p/o=y-10                441.9n ± 0%   421.1n ± 1%   -4.71% (p=0.001 n=10)
CORS/n=all/r=p/o=y/m=evil_acrh-10    448.0n ± 1%   123.7n ± 1%  -72.39% (p=0.000 n=10)
geomean                              530.2n        336.9n       -36.46%

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
CORS/n=all/r=p/o=y-10                 976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
CORS/n=all/r=p/o=y/m=evil_acrh-10     968.0 ± 0%     208.0 ± 0%  -78.51% (p=0.000 n=10)
geomean                              1016.4          670.7       -34.01%
¹ all samples are equal

                                  │  rs-cors   │            jub0bs-cors             │
                                  │ allocs/op  │ allocs/op   vs base                │
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
CORS/n=all/r=p/o=y-10               8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
CORS/n=all/r=p/o=y/m=evil_acrh-10   9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
geomean                             9.999        7.484       -25.16%
```

Nomenclature:
- `n` indicates the number of allowed origins: `one` | `two` | `multiple` | `many` | `all`
- `r` indicates the type of request (actual or preflight): `a` | `p`
- `o` indicates whether the request's origin is allowed: `y` | `n`
- `m` indicates a case involving a malicious request

For more details about the benchmark labeled `m=evil_acrh`,
see https://github.com/rs/cors/issues/170.
