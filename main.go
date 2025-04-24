package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/neWbie-saby/rss-aggregator/gin_handler"
	"github.com/neWbie-saby/rss-aggregator/internal/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	// feed, err := urlToFeed("http://feeds.bbci.co.uk/news/rss.xml")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(feed)

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT not found in the environment")
	}

	dbString := os.Getenv("DB_URL")

	if dbString == "" {
		log.Fatal("DB_URL not found in the environment")
	}

	conn, err := sql.Open("postgres", dbString)

	if err != nil {
		log.Fatal("Can't connect to database")
	}

	db := database.New(conn)

	go startScraping(db, 10, time.Minute)

	///Using go-chi package
	// router := chi.NewRouter()

	// router.Use(cors.Handler(cors.Options{
	// 	AllowedOrigins:   []string{"https://*", "http://*"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"*"},
	// 	ExposedHeaders:   []string{"Link"},
	// 	AllowCredentials: false,
	// 	MaxAge:           300,
	// }))

	// apiCfg := chi_handler.ApiConfig{
	// 	DB: db,
	// }

	// v1Router := chi.NewRouter()
	// v1Router.Get("/healthz", chi_handler.HandlerReadiness)
	// v1Router.Get("/err", chi_handler.HandlerErr)
	// v1Router.Post("/users", apiCfg.HandlerCreateUser)
	// v1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandleGetUser))

	// v1Router.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	// v1Router.Get("/feeds", apiCfg.HandlerGetFeeds)

	// v1Router.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	// v1Router.Get("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedFollows))
	// v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollow))

	// v1Router.Get("/posts", apiCfg.MiddlewareAuth(apiCfg.HandleGetPostsForUser))

	// router.Mount("/v1", v1Router)

	// srv := &http.Server{
	// 	Handler: router,
	// 	Addr:    ":" + portString,
	// }

	// log.Printf("Server starting on port %v", portString)

	// err = srv.ListenAndServe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	///Using Gin package
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	apiCfg := gin_handler.ApiConfig{
		DB: db,
	}

	v1 := router.Group("v1")
	{
		v1.GET("/healthz", gin_handler.GinHandlerReadiness)
		v1.GET("/err", gin_handler.GinHandlerErr)
		v1.POST("/users", apiCfg.HandlerCreateUser)
		v1.GET("/users", apiCfg.MiddlewareAuth(apiCfg.HandleGetUser))

		v1.POST("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
		v1.GET("/feeds", apiCfg.HandlerGetFeeds)

		v1.POST("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
		v1.GET("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedFollows))
		v1.DELETE("/feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollow))

		v1.GET("/posts", apiCfg.MiddlewareAuth(apiCfg.HandleGetPostsForUser))
	}

	log.Printf("Server starting on port %v", portString)
	if err := router.Run(":" + portString); err != nil {
		log.Fatal(err)
	}
}
