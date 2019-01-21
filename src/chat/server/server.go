package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"

	"github.com/99designs/gqlgen/handler"
	"github.com/AskAlexSharov/monorepo/src/chat"
	"github.com/gorilla/websocket"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/rs/cors"
	"sourcegraph.com/sourcegraph/appdash"
	appdashtracer "sourcegraph.com/sourcegraph/appdash/opentracing"
	"sourcegraph.com/sourcegraph/appdash/traceapp"
)

func main() {
	startAppdashServer()

	http.Handle("/", handler.Playground("Todo", "/query"))

	http.Handle("/query", GetHandlerWithMiddlewares())
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func GetHandlerWithMiddlewares() http.Handler {

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	userLoader := chat.NewUserLoader()
	loadersMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(chat.UserLoaderToCtx(r.Context(), userLoader))
			next.ServeHTTP(w, r)
		})
	}

	gqlHandler := handler.GraphQL(chat.NewExecutableSchema(chat.New()),
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}),
		handler.ErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
			return graphql.DefaultErrorPresenter(ctx, e)
		}),
	)

	return c.Handler(loadersMiddleware(gqlHandler))
}

func startAppdashServer() opentracing.Tracer {
	memStore := appdash.NewMemoryStore()
	store := &appdash.RecentStore{
		MinEvictAge: 5 * time.Minute,
		DeleteStore: memStore,
	}

	url, err := url.Parse("http://localhost:8700")
	if err != nil {
		log.Fatal(err)
	}
	tapp, err := traceapp.New(nil, url)
	if err != nil {
		log.Fatal(err)
	}
	tapp.Store = store
	tapp.Queryer = memStore

	go func() {
		log.Fatal(http.ListenAndServe(":8700", tapp))
	}()
	tapp.Store = store
	tapp.Queryer = memStore

	collector := appdash.NewLocalCollector(store)
	tracer := appdashtracer.NewTracer(collector)
	opentracing.InitGlobalTracer(tracer)

	log.Println("Appdash web UI running on HTTP :8700")
	return tracer
}
