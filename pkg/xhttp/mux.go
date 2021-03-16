package xhttp

import (
	"net/http"
	"path"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type muxEntrys []muxEntry

func (m muxEntrys) Len() int {
	return len(m)
}

func (m muxEntrys) Less(i, j int) bool {
	return strings.Count(m[i].pattern, "/") > strings.Count(m[j].pattern, "/")
}

func (m muxEntrys) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type RESTfulMux struct {
	mu sync.RWMutex
	me muxEntrys
}

func (mux *RESTfulMux) ServeHTTP(w http.ResponseWriter, r http.Request) {
	// h, _ := mux.Handler(r)
}

func (mux *RESTfulMux) Handler(r http.Request) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()
    path := cleanPath(r.URL.Path)
    mux.handler(r.Host, path)
}

func (mux *RESTfulMux) handler(host, path string) {
	for _, e := range mux.me {
        if e.host {
            e.match(host + path)
        }
        if nil == h {
            h 
        }
	}
}


func (mux *RESTfulMux) getPattern(pattern string) RESTfulHandler {
	for _, e := range mux.me {
		if e.pattern == pattern {
			return e.h
		}
	}
	return nil
}

func (mux *RESTfulMux) Handle(pattern string, handler RESTfulHandler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()
	if pattern == "" {
		panic("http: invalid pattern")
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if h := mux.getPattern(pattern); h != nil {
		panic("http: multiple registrations for " + pattern)
	}

	if mux.me == nil {
		mux.me = muxEntrys(make([]muxEntry, 0))
	}
	e := muxEntry{h: handler, pattern: pattern}
	if pattern[0] != '/' {
		e.host = true
	}
	e.reg = regexp.MustCompile(pattern)
	mux.me = append(mux.me, e)
	sort.Sort(mux.me)

}

type muxEntry struct {
	h       RESTfulHandler
	pattern string
	reg     *regexp.Regexp
	host    bool
}

func (me muxEntry) match(path string) {
    params := me.reg.FindStringSubmatch(path)
}

// cleanPath returns the canonical path for p, eliminating . and .. elements.
func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash except for root;
	// put the trailing slash back if necessary.
	if p[len(p)-1] == '/' && np != "/" {
		// Fast path for common case of p being the string we want:
		if len(p) == len(np)+1 && strings.HasPrefix(p, np) {
			np = p
		} else {
			np += "/"
		}
	}
	return np
}
