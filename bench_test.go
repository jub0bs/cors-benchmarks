package bench

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
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
	type BenchmarkCase struct {
		desc    string
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
			desc:              "single vs actual",
			handler:           dummyHandler,
			allowedOrigins:    []string{"https://example.com"},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
			},
		}, {
			desc:              "multiple vs actual",
			handler:           dummyHandler,
			allowedOrigins:    multipleOrigins,
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
			},
		}, {
			desc:    "pathological vs actual",
			handler: dummyHandler,
			allowedOrigins: []string{
				"https://a" + strings.Repeat(".a", hostMaxLen/2),
				"https://b" + strings.Repeat(".a", hostMaxLen/2),
			},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://c" + strings.Repeat(".a", hostMaxLen/2)},
			},
		}, {
			desc:              "many vs actual",
			handler:           dummyHandler,
			allowedOrigins:    manyOrigins,
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
			},
		}, {
			desc:              "any vs actual",
			handler:           dummyHandler,
			allowedOrigins:    []string{"*"},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodGet,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
			},
		}, {
			desc:              "all CORS headers vs actual",
			handler:           dummyHandler,
			allowedOrigins:    []string{"https://example.com"},
			credentialed:      true,
			allowPNA:          true,
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			maxAge:            30,
			reqMethod:         http.MethodPut,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
			},
		}, {
			desc:              "single vs preflight",
			handler:           dummyHandler,
			allowedOrigins:    []string{"https://example.com"},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
				headerACRM:   []string{http.MethodGet},
			},
		}, {
			desc:              "multiple vs preflight",
			handler:           dummyHandler,
			allowedOrigins:    multipleOrigins,
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
				headerACRM:   []string{http.MethodGet},
			},
		}, {
			desc:    "pathological vs preflight",
			handler: dummyHandler,
			allowedOrigins: []string{
				"https://a" + strings.Repeat(".a", hostMaxLen/2),
				"https://b" + strings.Repeat(".a", hostMaxLen/2),
			},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://c" + strings.Repeat(".a", hostMaxLen/2)},
				headerACRM:   []string{http.MethodGet},
			},
		}, {
			desc:              "many vs preflight",
			handler:           dummyHandler,
			allowedOrigins:    manyOrigins,
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
				headerACRM:   []string{http.MethodGet},
			},
		}, {
			desc:              "any vs preflight",
			handler:           dummyHandler,
			allowedOrigins:    []string{"*"},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
				headerACRM:   []string{http.MethodGet},
			},
		}, {
			desc:              "ACRH vs preflight",
			handler:           dummyHandler,
			allowedOrigins:    []string{"*"},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
				headerACRM:   []string{http.MethodGet},
				headerACRH:   []string{"content-length"},
			},
		}, {
			desc:              "all CORS headers vs preflight",
			handler:           dummyHandler,
			allowedOrigins:    []string{"https://example.com"},
			credentialed:      true,
			allowPNA:          true,
			allowedMethods:    []string{http.MethodPut},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			maxAge:            30,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
				headerACRM:   []string{http.MethodPut},
				headerACRH:   []string{"content-length"},
			},
		}, {
			desc:              "malicious ACRH vs preflight",
			handler:           dummyHandler,
			allowedOrigins:    []string{"*"},
			allowedReqHeaders: reqHeadersInDefaultRsCORS,
			reqMethod:         http.MethodOptions,
			reqHeaders: http.Header{
				headerOrigin: []string{"https://example.com"},
				headerACRM:   []string{http.MethodGet},
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
		desc := "mw=rs-cors/req=" + bc.desc
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
		desc = "mw=jub0bs-cors/req=" + bc.desc
		b.Run(desc, subBenchmark(jub0bsMw.Wrap(handler), req))
	}
}

var dummyHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
})

func newRequest(method string, hdrs http.Header) *http.Request {
	const dummyEndpoint = "https://example.com/whatever"
	req := httptest.NewRequest(method, dummyEndpoint, nil)
	for name, value := range hdrs {
		req.Header[name] = value
	}
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
	const n = 10
	for i := 0; i < n; i++ {
		manyOrigins = append(
			manyOrigins,
			// https
			fmt.Sprintf("https://%d.example.com", i),
			fmt.Sprintf("https://%d.example.com:7070", i),
			fmt.Sprintf("https://%d.example.com:8080", i),
			fmt.Sprintf("https://%d.example.com:9090", i),
			// one subdomain deep
			fmt.Sprintf("https://%d.foo.example.com", i),
			fmt.Sprintf("https://%d.foo.example.com:6060", i),
			fmt.Sprintf("https://%d.foo.example.com:7070", i),
			fmt.Sprintf("https://%d.foo.example.com:9090", i),
			// two subdomains deep
			fmt.Sprintf("https://%d.foo.bar.example.com", i),
			fmt.Sprintf("https://%d.foo.bar.example.com:6060", i),
			fmt.Sprintf("https://%d.foo.bar.example.com:7070", i),
			fmt.Sprintf("https://%d.foo.bar.example.com:9090", i),
			// arbitrary subdomains
			"https://*.foo.bar.example.com",
			"https://*.foo.bar.example.com:6060",
			"https://*.foo.bar.example.com:7070",
			"https://*.foo.bar.example.com:9090",
		)
	}
}

var reqHeadersInDefaultRsCORS = []string{
	"Accept",
	"Content-Type",
	"X-Requested-With",
}
