package main

import (
	"fmt"
	"net/http"
)

type MethodHandlers struct {
	Get    http.Handler
	Post   http.Handler
	Put    http.Handler
	Patch  http.Handler
	Delete http.Handler
}
type RoutesMap map[string]*MethodHandlers

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

func (m *Mux) Get(url string, handler http.Handler) {
	if _, ok := m.routes[url]; !ok {
		m.newRouteStruct(url)
	}

	handlers := m.routes[url]
	handlers.Get = handler
}

func (m *Mux) Post(url string, handler http.Handler) {
	if _, ok := m.routes[url]; !ok {
		m.newRouteStruct(url)
	}

	handlers := m.routes[url]
	handlers.Post = handler
}

func (m *Mux) Put(url string, handler http.Handler) {
	if _, ok := m.routes[url]; !ok {
		m.newRouteStruct(url)
	}

	handlers := m.routes[url]
	handlers.Put = handler
}

func (m *Mux) Patch(url string, handler http.Handler) {
	if _, ok := m.routes[url]; !ok {
		m.newRouteStruct(url)
	}

	handlers := m.routes[url]
	handlers.Patch = handler
}

func (m *Mux) Delete(url string, handler http.Handler) {
	if _, ok := m.routes[url]; !ok {
		m.newRouteStruct(url)
	}

	handlers := m.routes[url]
	handlers.Delete = handler
}

func (m *Mux) newRouteStruct(url string) {
	m.routes[url] = &MethodHandlers{}
}

func (m *Mux) parse() {
	if m.isParsed {
		return
	}

	doneCh := make(chan bool)

	for route, handlers := range m.routes {
		go func(route string, handlers *MethodHandlers) {
			m.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
				switch true {
				case route == "/" && r.URL.Path != "/":
					http.NotFound(w, r)

				case r.Method == http.MethodGet:
					if handlers.Get == nil {
						http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
						return
					}
					handlers.Get.ServeHTTP(w, r)

				case r.Method == http.MethodPost:
					if handlers.Post == nil {
						http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
						return
					}
					handlers.Post.ServeHTTP(w, r)

				case r.Method == http.MethodPut:
					if handlers.Put == nil {
						http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
						return
					}
					handlers.Put.ServeHTTP(w, r)

				case r.Method == http.MethodPatch:
					if handlers.Patch == nil {
						http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
						return
					}
					handlers.Patch.ServeHTTP(w, r)

				case r.Method == http.MethodDelete:
					if handlers.Delete == nil {
						http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
						return
					}
					handlers.Delete.ServeHTTP(w, r)

				default:
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			})
			doneCh <- true
		}(route, handlers)
	}

	for range m.routes {
		<-doneCh
	}

	fmt.Println("routes parsed")
	m.isParsed = true
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.parse()
	m.ServeMux.ServeHTTP(w, r)
}
