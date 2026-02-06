# Benchmarks comparing rs/cors and jub0bs/cors

This repo contains benchmarks (run with Go v1.25.7) that compare the
performance of two CORS middleware libraries:

- the more popular [rs/cors](https://github.com/rs/cors) (v1.11.1), and
- the more user-friendly [jub0bs/cors](https://github.com/jub0bs/cors) (v0.11.0).

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
Middleware/nb=one/req=a/o=y-10                           495.9n ± 1%   528.6n ± 1%   +6.59% (p=0.000 n=10)
Middleware/nb=one/req=a/o=n-10                           509.7n ± 2%   509.8n ± 0%        ~ (p=0.382 n=10)
Middleware/nb=multiple/req=a/o=y-10                      519.0n ± 1%   538.1n ± 0%   +3.69% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=n-10                      525.9n ± 1%   524.0n ± 0%   -0.37% (p=0.006 n=10)
Middleware/nb=two/req=a/o=y-10                           556.8n ± 1%   562.4n ± 1%   +1.01% (p=0.000 n=10)
Middleware/nb=two/req=a/o=n-10                           555.1n ± 1%   543.5n ± 1%   -2.10% (p=0.000 n=10)
Middleware/nb=many/req=a/o=y-10                         1012.5n ± 1%   549.2n ± 0%  -45.76% (p=0.000 n=10)
Middleware/nb=many/req=a/o=n-10                          607.9n ± 1%   508.5n ± 1%  -16.34% (p=0.000 n=10)
Middleware/nb=all/req=a/o=y-10                           516.4n ± 1%   513.5n ± 1%   -0.57% (p=0.002 n=10)
Middleware/nb=one/req=p/o=y-10                           443.8n ± 1%   430.5n ± 1%   -3.01% (p=0.000 n=10)
Middleware/nb=one/req=p/o=n-10                           381.4n ± 1%   110.2n ± 2%  -71.11% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=y-10                      447.8n ± 1%   431.7n ± 0%   -3.61% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=n-10                      388.5n ± 0%   115.4n ± 1%  -70.31% (p=0.000 n=10)
Middleware/nb=two/req=p/o=y-10                           477.1n ± 1%   454.4n ± 1%   -4.77% (p=0.000 n=10)
Middleware/nb=two/req=p/o=n-10                           420.4n ± 1%   133.4n ± 1%  -68.27% (p=0.000 n=10)
Middleware/nb=many/req=p/o=y-10                          918.1n ± 1%   432.6n ± 1%  -52.89% (p=0.000 n=10)
Middleware/nb=many/req=p/o=n-10                          460.8n ± 1%   114.7n ± 0%  -75.12% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y-10                           439.8n ± 0%   419.2n ± 1%   -4.68% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y/special=malicious_acrh-10    433.5n ± 1%   123.0n ± 1%  -71.64% (p=0.000 n=10)
geomean                                                  514.1n        339.5n       -33.96%

                                                      │   rs-cors    │              jub0bs-cors               │
                                                      │     B/op     │     B/op      vs base                  │
Middleware/nb=one/req=a/o=y-10                          1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
Middleware/nb=one/req=a/o=n-10                          1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=y-10                     1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
Middleware/nb=multiple/req=a/o=n-10                     1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
Middleware/nb=two/req=a/o=y-10                          1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
Middleware/nb=two/req=a/o=n-10                          1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
Middleware/nb=many/req=a/o=y-10                         1.047Ki ± 0%   1.047Ki ± 0%        ~ (p=1.000 n=10) ¹
Middleware/nb=many/req=a/o=n-10                         1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
Middleware/nb=all/req=a/o=y-10                          1.047Ki ± 0%   1.031Ki ± 0%   -1.49% (p=0.000 n=10)
Middleware/nb=one/req=p/o=y-10                            976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
Middleware/nb=one/req=p/o=n-10                            960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=y-10                       976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=n-10                       960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
Middleware/nb=two/req=p/o=y-10                            976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
Middleware/nb=two/req=p/o=n-10                            960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
Middleware/nb=many/req=p/o=y-10                           976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
Middleware/nb=many/req=p/o=n-10                           960.0 ± 0%     208.0 ± 0%  -78.33% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y-10                            976.0 ± 0%     944.0 ± 0%   -3.28% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y/special=malicious_acrh-10     968.0 ± 0%     208.0 ± 0%  -78.51% (p=0.000 n=10)
geomean                                                  1016.4          670.7       -34.01%
¹ all samples are equal

                                                      │  rs-cors   │            jub0bs-cors             │
                                                      │ allocs/op  │ allocs/op   vs base                │
Middleware/nb=one/req=a/o=y-10                          12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=one/req=a/o=n-10                          12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=y-10                     12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=multiple/req=a/o=n-10                     12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=two/req=a/o=y-10                          12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=two/req=a/o=n-10                          12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=many/req=a/o=y-10                         12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=many/req=a/o=n-10                         12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=all/req=a/o=y-10                          12.00 ± 0%   11.00 ± 0%   -8.33% (p=0.000 n=10)
Middleware/nb=one/req=p/o=y-10                          8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
Middleware/nb=one/req=p/o=n-10                          9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=y-10                     8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
Middleware/nb=multiple/req=p/o=n-10                     9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
Middleware/nb=two/req=p/o=y-10                          8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
Middleware/nb=two/req=p/o=n-10                          9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
Middleware/nb=many/req=p/o=y-10                         8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
Middleware/nb=many/req=p/o=n-10                         9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y-10                          8.000 ± 0%   7.000 ± 0%  -12.50% (p=0.000 n=10)
Middleware/nb=all/req=p/o=y/special=malicious_acrh-10   9.000 ± 0%   4.000 ± 0%  -55.56% (p=0.000 n=10)
geomean                                                 9.999        7.484       -25.16%
```

Nomenclature:
- `nb` indicates the number of allowed origins: `one` | `two` | `multiple` | `many` | `all`
- `req` indicates the type of request (actual or preflight): `a` | `p`
- `o` indicates whether the request's origin is allowed: `y` | `n`
- `special` indicates a case of interest.

For more details about the benchmark labeled `special=malicious_acrh`,
see https://github.com/rs/cors/issues/170.
