package util

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type ClosableResponseWriter interface {
	http.ResponseWriter
	Close()
}

type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

func (this gzipResponseWriter) Write(data []byte) (int, error) {
	return this.Writer.Write(data)
}

func (this gzipResponseWriter) Close() {
	this.Writer.Close()
}

//since gzip.writer has an struct attribute called Header, that was causing ambiguity if we don't have this method
func (this gzipResponseWriter) Header() http.Header {
	return this.ResponseWriter.Header()
}

type closableResponseWriter struct {
	http.ResponseWriter
}

func (this closableResponseWriter) Close() {
	//noop
}

func GetResponseWriter(w http.ResponseWriter, req *http.Request) ClosableResponseWriter {
	if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		return gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gzip.NewWriter(w),
		}
	}

	return closableResponseWriter{
		ResponseWriter: w,
	}
}
