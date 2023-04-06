package main

import (
	"fmt"
	"net/http"
)

type RoutesMap map[string]*struct {
	Get    http.HandlerFunc
	Post   http.HandlerFunc
	Put    http.HandlerFunc
	Patch  http.HandlerFunc
	Delete http.HandlerFunc
}
type Mux struct {
	*http.ServeMux
	routes   RoutesMap
	isParsed bool
}

func NewServeMux() *Mux {
	return &Mux{
		ServeMux: http.NewServeMux(),
		routes:   make(RoutesMap),
		isParsed: false,
	}
}

func (m *Mux) Get(url string, handler http.HandlerFunc) {
	if _, ok := m.routes[url]; !ok {
		m.newRouteStruct(url)
	}

	handlers := m.routes[url]
	handlers.Get = handler
}

func (m *Mux) Post(url string, handler http.HandlerFunc) {
	if _, ok := m.routes[url]; !ok {
		m.newRouteStruct(url)
	}

	handlers := m.routes[url]
	handlers.Post = handler
}

func (m *Mux) Put(url string, handler http.HandlerFunc) {
	if _, ok := m.routes[url]; !ok {
		m.newRouteStruct(url)
	}

	handlers := m.routes[url]
	handlers.Put = handler
}

func (m *Mux) Patch(url string, handler http.HandlerFunc) {
	if _, ok := m.routes[url]; !ok {
		m.newRouteStruct(url)
	}

	handlers := m.routes[url]
	handlers.Patch = handler
}

func (m *Mux) Delete(url string, handler http.HandlerFunc) {
	if _, ok := m.routes[url]; !ok {
		m.newRouteStruct(url)
	}

	handlers := m.routes[url]
	handlers.Delete = handler
}

func (m *Mux) newRouteStruct(url string) {
	m.routes[url] = &struct {
		Get    http.HandlerFunc
		Post   http.HandlerFunc
		Put    http.HandlerFunc
		Patch  http.HandlerFunc
		Delete http.HandlerFunc
	}{
		Get:    nil,
		Post:   nil,
		Put:    nil,
		Patch:  nil,
		Delete: nil,
	}
}

func (m *Mux) parse() {
	if !m.isParsed {
		for route, handlers := range m.routes {
			m.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
				if route == "/" && r.URL.Path != "/" {
					http.NotFound(w, r)
					return
				}

				if r.Method == http.MethodGet {
					if handlers.Get == nil {
						http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
						return
					}

					handlers.Get(w, r)
					return
				}

				if r.Method == http.MethodPost {
					if handlers.Post == nil {
						http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
						return
					}

					handlers.Post(w, r)
					return
				}

				if r.Method == http.MethodPut {
					if handlers.Put == nil {
						http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
						return
					}

					handlers.Put(w, r)
					return
				}

				if r.Method == http.MethodPatch {
					if handlers.Patch == nil {
						http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
						return
					}

					handlers.Patch(w, r)
					return
				}

				if r.Method == http.MethodDelete {
					if handlers.Delete == nil {
						http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
						return
					}

					handlers.Delete(w, r)
					return
				}

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			})
		}
		fmt.Println("routes parsed")
		m.isParsed = true
	}
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.parse()
	m.ServeMux.ServeHTTP(w, r)
}
