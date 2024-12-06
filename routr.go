package routr

import (
	"fmt"
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

type Routr struct {
	mw  []Middleware
	mux *http.ServeMux
}

func (r *Routr) wrap(handler http.HandlerFunc) http.Handler  {
	var final http.Handler
	for i := len(r.mw)-1; i >= 0; i-- {
		final = r.mw[i](handler)
	}
	return final
}

func (r *Routr) Use(m Middleware) {
	r.mw = append(r.mw, m)
}

func (r *Routr) Get(path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("%v %v", http.MethodGet, path)
	chainedHandler := r.wrap(handler)
	
	r.mux.HandleFunc(pattern, chainedHandler.ServeHTTP)
}

func (r *Routr) Post(path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("%v %v", http.MethodPost, path)
	chainedHandler := r.wrap(handler)
	r.mux.HandleFunc(pattern, chainedHandler.ServeHTTP)
}

func (r *Routr) Patch(path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("%v %v", http.MethodPatch, path)
	chainedHandler := r.wrap(handler)
	r.mux.HandleFunc(pattern, chainedHandler.ServeHTTP)
}

func (r *Routr) Put(path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("%v %v", http.MethodPut, path)
	chainedHandler := r.wrap(handler)
	r.mux.HandleFunc(pattern, chainedHandler.ServeHTTP)
}

func (r *Routr) Delete(path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("%v %v", http.MethodDelete, path)
	chainedHandler := r.wrap(handler)
	r.mux.HandleFunc(pattern, chainedHandler.ServeHTTP)
}

func (r *Routr) Connect(path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("%v %v", http.MethodConnect, path)
	chainedHandler := r.wrap(handler)
	r.mux.HandleFunc(pattern, chainedHandler.ServeHTTP)
}

func (r *Routr) Options(path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("%v %v", http.MethodOptions, path)
	chainedHandler := r.wrap(handler)
	r.mux.HandleFunc(pattern, chainedHandler.ServeHTTP)
}

func (r *Routr) Trace(path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("%v %v", http.MethodTrace, path)
	chainedHandler := r.wrap(handler)
	r.mux.HandleFunc(pattern, chainedHandler.ServeHTTP)
}

func (r *Routr) Head(path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("%v %v", http.MethodHead, path)
	chainedHandler := r.wrap(handler)
	r.mux.HandleFunc(pattern, chainedHandler.ServeHTTP)
}

func (r *Routr) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func NewRouter() *Routr {
	return &Routr{
		mux: http.NewServeMux(),
	}
}
