package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/KnightHacks/knighthacks_hackathon/graph"
	"github.com/KnightHacks/knighthacks_hackathon/graph/generated"
	"github.com/KnightHacks/knighthacks_hackathon/repository"
	"github.com/KnightHacks/knighthacks_shared/auth"
	"github.com/KnightHacks/knighthacks_shared/azure_blob"
	"github.com/KnightHacks/knighthacks_shared/database"
	"github.com/KnightHacks/knighthacks_shared/pagination"
	"github.com/KnightHacks/knighthacks_shared/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"os"
	"runtime/debug"
)

const defaultPort = "8080"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	pool, err := database.ConnectWithRetries(utils.GetEnvOrDie("DATABASE_URI"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	newAuth, err := auth.NewAuthWithEnvironment()
	if err != nil {
		log.Fatalf("An error occured when trying to create an instance of Auth: %s\n", err)
	}

	var client *azure_blob.AzureBlobClient
	_, exists := os.LookupEnv("AZURE_ACCOUNT_NAME")
	if exists {
		credential, err := azure_blob.NewSharedCredentialFromEnv()
		if err != nil {
			log.Fatalf("error occured while making azure shared credential, err = %v", err)
		}
		client, err = azure_blob.NewAzureBlobClient(credential)
		if err != nil {
			log.Fatalf("error occured while making azure blob client, err = %v", err)
		}
	}

	ginRouter := gin.Default()
	ginRouter.Use(auth.AuthContextMiddleware(newAuth))
	ginRouter.Use(utils.GinContextMiddleware())

	ginRouter.POST("/query", graphqlHandler(newAuth, pool, client))
	ginRouter.GET("/", playgroundHandler())

	log.Fatal(ginRouter.Run(":" + port))
}

func graphqlHandler(a *auth.Auth, pool *pgxpool.Pool, client *azure_blob.AzureBlobClient) gin.HandlerFunc {
	// TODO: Sponsor doesn't have a sense of ownership, maybe we should have sponsor linked users?

	hasRoleDirective := auth.HasRoleDirective{GetUserId: auth.DefaultGetUserId, Queryable: pool}

	config := generated.Config{
		Resolvers: &graph.Resolver{
			Repository:      repository.NewDatabaseRepository(pool, client),
			AzureBlobClient: client,
			Auth:            a,
		},
		Directives: generated.DirectiveRoot{
			HasRole:    hasRoleDirective.Direct,
			Pagination: pagination.Pagination,
		},
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		log.Println("Error presented: ", err)
		debug.PrintStack()
		return graphql.DefaultErrorPresenter(ctx, err)
	})
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
