package bench

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"

	"github.com/jub0bs/cors"
	rsCors "github.com/rs/cors"
)

const (
	headerOrigin = "Origin"

	headerACRM = "Access-Control-Request-Method"
	headerACRH = "Access-Control-Request-Headers"
)

const hostMaxLen = 253

func BenchmarkMiddleware(b *testing.B) {
	// preliminary sanity checks
	if len(multipleOrigins) == 0 {
		b.Fatalf("multipleOrigins is empty")
	}
	if slices.Contains(multipleOrigins, disallowedOrigin) {
		b.Fatalf("multipleOrigins contains %q but should not", disallowedOrigin)
	}
	if len(manyOrigins) == 0 {
		b.Fatalf("manyOrigins is empty")
	}
	if slices.Contains(manyOrigins, disallowedOrigin) {
		b.Fatalf("manyOrigins contains %q but should not", disallowedOrigin)
	}

	type BenchmarkCase struct {
		desc    string // type=a|p/o=y/n/nallowed=one|two|multiple|many
		handler http.Handler
		// CORS config
		allowedOrigins    []string
		credentialed      bool
		allowPNA          bool
		allowedMethods    []string
		allowedReqHeaders []string
		maxAge            int
		exposedResHeaders []string
		// request
		reqMethod  string
		reqHeaders http.Header
	}
	cases := []BenchmarkCase{
		{
			desc:      "", // no CORS middleware
			handler:   dummyHandler,
			reqMethod: http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
			},
		}, {
			desc:              "nb=one/req=a/o=y",
			handler:           dummyHandler,
			allowedOrigins:    []string{"https://example.com"},
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
			},
		}, {
			desc:              "nb=one/req=a/o=n",
			handler:           dummyHandler,
			allowedOrigins:    []string{"https://example.com"},
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.org"},
			},
		}, {
			desc:              "nb=multiple/req=a/o=y",
			handler:           dummyHandler,
			allowedOrigins:    multipleOrigins,
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{last(multipleOrigins)},
			},
		}, {
			desc:              "nb=multiple/req=a/o=n",
			handler:           dummyHandler,
			allowedOrigins:    multipleOrigins,
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{disallowedOrigin},
			},
		}, {
			desc:    "nb=two/req=a/o=y",
			handler: dummyHandler,
			allowedOrigins: []string{
				"https://a" + strings.Repeat(".a", hostMaxLen/2),
				"https://b" + strings.Repeat(".a", hostMaxLen/2),
			},
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://a" + strings.Repeat(".a", hostMaxLen/2)},
			},
		}, {
			desc:    "nb=two/req=a/o=n",
			handler: dummyHandler,
			allowedOrigins: []string{
				"https://a" + strings.Repeat(".a", hostMaxLen/2),
				"https://b" + strings.Repeat(".a", hostMaxLen/2),
			},
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://c" + strings.Repeat(".a", hostMaxLen/2)},
			},
		}, {
			desc:              "nb=many/req=a/o=y",
			handler:           dummyHandler,
			allowedOrigins:    manyOrigins,
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{last(manyOrigins)},
			},
		}, {
			desc:              "nb=many/req=a/o=n",
			handler:           dummyHandler,
			allowedOrigins:    manyOrigins,
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{disallowedOrigin},
			},
		}, {
			desc:              "nb=all/req=a/o=y",
			handler:           dummyHandler,
			allowedOrigins:    []string{"*"},
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
			},
		}, {
			desc:              "nb=one/req=p/o=y",
			handler:           dummyHandler,
			allowedOrigins:    []string{"https://example.com"},
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
				headerACRM:   []string{http.MethodPut},
			},
		}, {
			desc:              "nb=one/req=p/o=n",
			handler:           dummyHandler,
			allowedOrigins:    []string{"https://example.com"},
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.org"},
				headerACRM:   []string{http.MethodPut},
			},
		}, {
			desc:              "nb=multiple/req=p/o=y",
			handler:           dummyHandler,
			allowedOrigins:    multipleOrigins,
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{last(multipleOrigins)},
				headerACRM:   []string{http.MethodPut},
			},
		}, {
			desc:              "nb=multiple/req=p/o=n",
			handler:           dummyHandler,
			allowedOrigins:    multipleOrigins,
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{disallowedOrigin},
				headerACRM:   []string{http.MethodPut},
			},
		}, {
			desc:    "nb=two/req=p/o=y",
			handler: dummyHandler,
			allowedOrigins: []string{
				"https://a" + strings.Repeat(".a", hostMaxLen/2),
				"https://b" + strings.Repeat(".a", hostMaxLen/2),
			},
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://a" + strings.Repeat(".a", hostMaxLen/2)},
				headerACRM:   []string{http.MethodPut},
			},
		}, {
			desc:    "nb=two/req=p/o=n",
			handler: dummyHandler,
			allowedOrigins: []string{
				"https://a" + strings.Repeat(".a", hostMaxLen/2),
				"https://b" + strings.Repeat(".a", hostMaxLen/2),
			},
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://c" + strings.Repeat(".a", hostMaxLen/2)},
				headerACRM:   []string{http.MethodPut},
			},
		}, {
			desc:              "nb=many/req=p/o=y",
			handler:           dummyHandler,
			allowedOrigins:    manyOrigins,
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{last(manyOrigins)},
				headerACRM:   []string{http.MethodPut},
			},
		}, {
			desc:              "nb=many/req=p/o=n",
			handler:           dummyHandler,
			allowedOrigins:    manyOrigins,
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{disallowedOrigin},
				headerACRM:   []string{http.MethodPut},
			},
		}, {
			desc:              "nb=all/req=p/o=y",
			handler:           dummyHandler,
			allowedOrigins:    []string{"*"},
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
				headerACRM:   []string{http.MethodPut},
			},
		}, {
			desc:              "nb=all/req=p/o=y/special=malicious_acrh",
			handler:           dummyHandler,
			allowedOrigins:    []string{"*"},
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
				headerACRM:   []string{http.MethodPut},
				headerACRH:   []string{strings.Repeat(",", 1024)},
			},
		},
	}

	for _, bc := range cases {
		req := newRequest(bc.reqMethod, bc.reqHeaders)

		var handler http.Handler = bc.handler
		if bc.allowedOrigins == nil { // no CORS
			b.Run("mw=none", subBenchmark(handler, req))
			continue
		}

		// rs/cors
		rsMw := rsCors.New(rsCors.Options{
			AllowedOrigins:   bc.allowedOrigins,
			AllowCredentials: bc.credentialed,
			AllowedMethods:   bc.allowedMethods,
			AllowedHeaders:   bc.allowedReqHeaders,
			MaxAge:           bc.maxAge,
			ExposedHeaders:   bc.exposedResHeaders,
		})
		desc := "mw=rs-cors/" + bc.desc
		b.Run(desc, subBenchmark(rsMw.Handler(handler), req))

		// jub0bs/cors
		jub0bsMw, err := cors.NewMiddleware(cors.Config{
			Origins:         bc.allowedOrigins,
			Credentialed:    bc.credentialed,
			Methods:         bc.allowedMethods,
			RequestHeaders:  bc.allowedReqHeaders,
			MaxAgeInSeconds: bc.maxAge,
			ResponseHeaders: bc.exposedResHeaders,
		})
		if err != nil {
			b.Fatal(err)
		}
		desc = "mw=jub0bs-cors/" + bc.desc
		b.Run(desc, subBenchmark(jub0bsMw.Wrap(handler), req))
	}
}

var dummyHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
})

func newRequest(method string, hdrs http.Header) *http.Request {
	const dummyEndpoint = "https://example.com/whatever"
	req := httptest.NewRequest(method, dummyEndpoint, nil)
	req.Header = hdrs
	return req
}

func subBenchmark(handler http.Handler, req *http.Request) func(*testing.B) {
	return func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rec := httptest.NewRecorder()
				handler.ServeHTTP(rec, req)
			}
		})
	}
}

const disallowedOrigin = "https://example.org:6060"

var multipleOrigins = []string{
	"https://*.example.net",
	"https://example.net:8080",
	"https://example.net",
	"https://*.example.org",
	"https://example.org:8080",
	"https://example.org",
	"https://*.example.com",
	"https://example.com:8080",
	"https://example.com",
}

var manyOrigins []string

func init() { // populates manyOrigins
	const n = 1000
	manyOrigins = make([]string, n)
	// Make all origins the same length.
	width := int(math.Ceil(math.Log10(n))) // max digits of numbers in [1, n)
	tmpl := fmt.Sprintf("https://%%0%dd.example.com", width)
	for i := range n {
		manyOrigins[i] = fmt.Sprintf(tmpl, i)
	}
}

var reqHeadersInDefaultRsCORS = []string{
	"Accept",
	"Content-Type",
	"X-Requested-With",
}

// Precondition: s is not empty
func last(s []string) string {
	return s[len(s)-1]
}
