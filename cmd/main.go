package main

import (
	"test/database"
	"test/internal/graph"
	"test/internal/handlers"
	"test/internal/middleware"
	"test/internal/repositories"
	"test/internal/services"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func graphqlHandler(resolver *graph.Resolver) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQl", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()
	db := database.StartPostgres()
	repo := repositories.NewTestRepo(db)
	resolver := graph.NewGraphQlResolver(repo)
	// service := services.NewTestService(repo)

	// middleware := middleware.NewTweetMiddleware(service)
	// handlers.NewTestHandler(r, middleware, service)
	r.POST("/query", graphqlHandler(resolver))
	r.GET("/play", playgroundHandler())
	r.Run(":2400")
}
