package server

import (
	"fmt"
	"http-fileserver/internal/middleware"
	"http-fileserver/internal/qrcode"
	"http-fileserver/internal/utils"
	"net/http"
	"strconv"
)

type FileServer struct {
	port   int
	folder string
	server *http.Server
}

func New(port int, folder string) *FileServer {
	return &FileServer{
		port:   port,
		folder: folder,
	}
}

func (fs *FileServer) Start() error {

	mux := http.NewServeMux()

	fs.registerRoutes(mux)

	fs.server = &http.Server{
		Addr:    ":" + strconv.Itoa(fs.port),
		Handler: mux,
	}

	url := fmt.Sprintf("http://%s:%d/local", utils.GetLocalIpAddress(), fs.port)

	qrcode.Generate(url)

	fmt.Println("Server is listening on " + url)

	return fs.server.ListenAndServe()
}

func (fs *FileServer) registerRoutes(mux *http.ServeMux) {

	handle := func(w http.ResponseWriter, r *http.Request) {
		fileHandler := http.StripPrefix("/local/", http.FileServer(http.Dir(fs.folder)))
		fileHandler.ServeHTTP(w, r)
	}

	mux.Handle("/local/", middleware.NoCache(handle))

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

}
