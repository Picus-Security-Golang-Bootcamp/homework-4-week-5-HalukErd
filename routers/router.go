package routers

import (
	"context"
	routers "github.com/HalukErd/Week5Assignment/routers/api"
	"github.com/HalukErd/Week5Assignment/service"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type ApiController struct {
	bookService   *service.BookService
	authorService *service.AuthorService
}

func NewApiController(bookService *service.BookService, authorService *service.AuthorService) *ApiController {
	return &ApiController{bookService: bookService, authorService: authorService}
}

func (cont *ApiController) InitRouter() {
	r := mux.NewRouter()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	rV1 := r.PathPrefix("/api/v1").Subrouter()

	cont.InitBookRouter(rV1.PathPrefix("/book").Subrouter())
	cont.InitAuthorRouter(rV1.PathPrefix("/author").Subrouter())

	srv := &http.Server{
		Addr:         "0.0.0.0:8090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Listen has failed.", err)
		}
	}()

	ShutdownServer(srv, time.Second*10)
}

func (cont *ApiController) InitAuthorRouter(r *mux.Router) {
	r.HandleFunc("/", routers.GetAllAuthors)
}

func (cont *ApiController) InitBookRouter(r *mux.Router) {
	r.HandleFunc("/", routers.GetAllBooks)
}

func ShutdownServer(srv *http.Server, timeout time.Duration) {
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down book service")
	os.Exit(0)
}
