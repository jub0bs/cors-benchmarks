# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.27.1) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v1.1.2).

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
CORS/n=one/r=a/o=y-10                512.2n ± 2%   534.0n ± 1%   +4.26% (p=0.000 n=10)
CORS/n=one/r=a/o=n-10                516.2n ± 1%   510.9n ± 0%   -1.05% (p=0.000 n=10)
CORS/n=multiple/r=a/o=y-10           519.5n ± 0%   539.0n ± 0%   +3.74% (p=0.000 n=10)
CORS/n=multiple/r=a/o=n-10           528.0n ± 0%   519.7n ± 0%   -1.56% (p=0.000 n=10)
CORS/n=two/r=a/o=y-10                570.1n ± 0%   546.1n ± 3%   -4.22% (p=0.000 n=10)
CORS/n=two/r=a/o=n-10                571.6n ± 0%   519.7n ± 0%   -9.07% (p=0.000 n=10)
CORS/n=many/r=a/o=y-10              1020.5n ± 0%   557.9n ± 0%  -45.34% (p=0.000 n=10)
CORS/n=many/r=a/o=n-10               629.8n ± 0%   511.8n ± 0%  -18.74% (p=0.000 n=10)
CORS/n=all/r=a/o=y-10                517.9n ± 0%   501.8n ± 0%   -3.11% (p=0.000 n=10)
CORS/n=one/r=p/o=y-10                441.3n ± 1%   382.4n ± 0%  -13.36% (p=0.000 n=10)
CORS/n=one/r=p/o=n-10               369.15n ± 0%   98.03n ± 0%  -73.44% (p=0.000 n=10)
CORS/n=multiple/r=p/o=y-10           443.2n ± 1%   383.2n ± 0%  -13.53% (p=0.000 n=10)
CORS/n=multiple/r=p/o=n-10           374.7n ± 0%   105.5n ± 0%  -71.86% (p=0.000 n=10)
CORS/n=two/r=p/o=y-10                494.1n ± 1%   387.5n ± 1%  -21.57% (p=0.000 n=10)
CORS/n=two/r=p/o=n-10                422.1n ± 0%   103.5n ± 1%  -75.49% (p=0.000 n=10)
CORS/n=many/r=p/o=y-10               924.6n ± 0%   403.6n ± 0%  -56.35% (p=0.000 n=10)
CORS/n=many/r=p/o=n-10              481.10n ± 1%   99.83n ± 0%  -79.25% (p=0.000 n=10)
CORS/n=all/r=p/o=y-10                433.1n ± 0%   380.2n ± 0%  -12.19% (p=0.000 n=10)
CORS/n=all/r=p/o=y/m=evil_acrh-10    436.4n ± 0%   118.0n ± 0%  -72.96% (p=0.000 n=10)
geomean                              518.0n        317.5n       -38.71%

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
