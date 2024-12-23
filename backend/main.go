package main

import (
	"log"
	"os"

	"github.com/Eizeed/vibe_gogo/db"
	"github.com/Eizeed/vibe_gogo/handlers"
	"github.com/Eizeed/vibe_gogo/middleware"
	"github.com/Eizeed/vibe_gogo/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
    err := godotenv.Load(".env")
    if err != nil {
		log.Fatal("error: failed to load the env file")
	}

    db.InitDB();

    // Change in prod
    db.GetDB().AutoMigrate(&models.User{})
    db.GetDB().AutoMigrate(&models.Playlist{})

	//Start the default gin server
	app := gin.Default()

    r := app.Group("/api");
	{
        users := handlers.UserHandler {};

        r.GET("/users", users.GetAll);
        r.POST("/users/register", users.Register);
        r.POST("/users/login", users.Login);
        r.PUT("/users/:uuid", middleware.AuthMiddleware(), users.Update);
        r.DELETE("/users/:uuid", middleware.AuthMiddleware(), users.Delete);

        playlists := handlers.PlaylistHandler {};

        r.GET("/playlists/user/:uuid", playlists.GetByUserUuid);
        r.GET("/playlists/:uuid", playlists.GetByUuid);
        r.POST("/playlists", middleware.AuthMiddleware(), playlists.Create);
        r.PATCH("/playlists/:uuid/addtrack", middleware.AuthMiddleware(), playlists.AddTrack);
        r.PATCH("/playlists/:uuid/deletetrack", middleware.AuthMiddleware(), playlists.DeleteTrack);
        r.PUT("/playlists/:uuid", middleware.AuthMiddleware(), playlists.Update);
        r.PUT("/playlists/:uuid/visibility", middleware.AuthMiddleware(), playlists.ChangeVisibility);
        r.DELETE("/playlists/:uuid", middleware.AuthMiddleware(), playlists.Delete);

        deezer := handlers.DeezerHandler {};

        r.GET("/deezer/search", deezer.Search);
    }


	// app.NoRoute(func(c *gin.Context) {
    //     c.HTML(404, gin.H{"error": "No path implemented"})
	// })

	port := os.Getenv("PORT")

	if os.Getenv("SSL") == "TRUE" {

		//Generated using sh generate-certificate.sh
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{
			CERT: "./cert/myCA.cer",
			KEY:  "./cert/myCA.key",
		}

		app.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
	    app.Run(":" + port)
	}
}
