package main

import (
	"net/http"

	"net/http/httptrace"
	"net/http/httputil"
	"net/url"

	"flag"
	log "github.com/Sirupsen/logrus"
	loghttp "github.com/motemen/go-loghttp"
	"github.com/vulcand/oxy/forward"
)

var upstreamUrl *url.URL

var clientTracer *httptrace.ClientTrace

func init() {
	upHost := flag.String("u", "http://mockbin.com", "scheme and host of the target upstream server")
	verbose := flag.Bool("v", false, "verbose")
	flag.Parse()

	var ierr error
	upstreamUrl, ierr = url.ParseRequestURI(*upHost)
	if ierr != nil {
		log.Fatalf("%v: %s", ierr, *upHost)
		return
	}

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, ForceColors: true})
	log.SetLevel(log.DebugLevel)

	http.DefaultTransport = &loghttp.Transport{
		Transport: http.DefaultTransport,
		LogRequest: func(req *http.Request) {
			logRequestOut(req, "upstream")
		},
		LogResponse: func(resp *http.Response) {
			logResponse(resp, "upstream")
		},
	}

	if *verbose {
		clientTracer = &httptrace.ClientTrace{
			GetConn: func(hostPort string) {
				log.Debugf("httptrace.ClientTrace.GettingConn: %s", hostPort)
			},
			GotConn: func(connInfo httptrace.GotConnInfo) {
				log.Debugf("httptrace.ClientTrace.GotConn: %+v", connInfo)
			},
			PutIdleConn: func(err error) {
				if err != nil {
					log.Error("httptrace.ClientTrace.PutIdleConn:", err)
				} else {
					log.Debug("httptrace.ClientTrace.PutIdleConn")
				}
			},
			GotFirstResponseByte: func() {
				log.Debug("httptrace.ClientTrace.GotFirstResponseByte")
			},
			Got100Continue: func() {
				log.Debug("httptrace.ClientTrace.Got100Continue")
			},
			DNSStart: func(dnsInfo httptrace.DNSStartInfo) {
				log.Debugf("httptrace.ClientTrace.DNSStart: %+v", dnsInfo)
			},
			DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
				log.Debugf("httptrace.ClientTrace.DNSDone: %+v", dnsInfo)
			},
			ConnectStart: func(network, addr string) {
				log.Debugf("httptrace.ClientTrace.ConnectStart: %s %s", network, addr)
			},
			ConnectDone: func(network, addr string, err error) {
				log.Debugf("httptrace.ClientTrace.ConnectDone: %s %s %v", network, addr, err)
			},
			WroteHeaders: func() {
				log.Debug("httptrace.ClientTrace.WroteHeaders")
			},
			Wait100Continue: func() {
				log.Debug("httptrace.ClientTrace.Wait100Continue")
			},
			WroteRequest: func(info httptrace.WroteRequestInfo) {
				log.Debugf("httptrace.ClientTrace.WroteRequest: %+v", info)
			},
		}
	}
}

func logRequest(req *http.Request, tag string) error {
	dump, err := httputil.DumpRequest(req, true)
	if err == nil {
		log.Infof("%s --> %s\n====", tag, dump)
	}
	return err
}

func logRequestOut(req *http.Request, tag string) error {
	dump, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		log.Infof("%s --> %s\n====", tag, dump)
	}
	return err
}
func logResponse(resp *http.Response, tag string) {
	dump, err := httputil.DumpResponse(resp, true)
	if err == nil {
		log.Infof("%s <-- %s\n====", tag, dump)
	}
}

func main() {
	// Forwards incoming requests to whatever location URL points to, adds proper forwarding headers
	fwd, _ := forward.New()
	redirect := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logRequest(req, "downstream")

		if clientTracer != nil {
			req = req.WithContext(httptrace.WithClientTrace(req.Context(), clientTracer))
		}

		// let us forward this request to another server
		req.URL = upstreamUrl
		fwd.ServeHTTP(w, req)
	})

	// that's it! our reverse proxy is ready!
	s := &http.Server{
		Addr:    ":8080",
		Handler: redirect,
	}
	log.Fatal(s.ListenAndServe())
}
