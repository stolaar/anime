package routes

import (
	"fmt"

	"anime/server/handlers"

	s "anime/server"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *s.Server) {
	postHandler := handlers.NewPostHandlers(server)
	authHandler := handlers.NewAuthHandler(server)
	animeHandler := handlers.NewAnimeHandlers(server)
	registerHandler := handlers.NewRegisterHandler(server)

	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST("/login", authHandler.Login)
	server.Echo.POST("/register", registerHandler.Register)
	server.Echo.POST("/refresh", authHandler.RefreshToken)

	fmt.Println(server.Config.Auth.AccessSecret)

	r := server.Echo.Group("/api")
	// Configure middleware with the custom claims type
	/* config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(token.JwtCustomClaims)
		},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	r.Use(echojwt.WithConfig(config))
	*/

	r.GET("/posts", postHandler.GetPosts)
	r.POST("/posts", postHandler.CreatePost)
	r.DELETE("/posts/:id", postHandler.DeletePost)
	r.PUT("/posts/:id", postHandler.UpdatePost)
	r.GET("/animes", animeHandler.GetAnimes)
	r.GET("/animes/:id", animeHandler.GetAnimeEpisodes)
	r.GET("/episode/:id/src", animeHandler.GetAnimeEpisodeSrc)
}
