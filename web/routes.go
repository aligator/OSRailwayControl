package web

import (
	"OSRailwayControl/bindata"
	"fmt"
	"mime"
	"net/http"
	"net/url"
	"path"
)

func (w *web) setupRoutes() {
	w.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveSpecificEmbeddedFile("/client/index.html", w, r)
	})
	w.router.PathPrefix("/ws").HandlerFunc(func(writer http.ResponseWriter, reader *http.Request) {
		w.Socket().ServeHTTP(writer, reader)
	})
	w.router.PathPrefix("/").HandlerFunc(serveEmbeddedFile)
}

func serveEmbeddedFile(w http.ResponseWriter, r *http.Request) {
	var urlPath string

	u, err := url.Parse(r.URL.Path)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	urlPath = "/client" + u.String()

	serveSpecificEmbeddedFile(urlPath, w, r)
}

func serveSpecificEmbeddedFile(filePath string, w http.ResponseWriter, _ *http.Request) {
	if data, err := bindata.Asset(filePath); err == nil {
		mimeType := mime.TypeByExtension(path.Ext(filePath))

		fmt.Println(mimeType, path.Ext(filePath))

		if mimeType == "" {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("unknown content type: " + path.Ext(filePath)))
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		w.Header().Add("Content-Type", mimeType)
		_, err := w.Write(data)
		if err != nil {
			fmt.Println(err)
		}
		return
	} else {
		if err.Error() == "file does not exist" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}
