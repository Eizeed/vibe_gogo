package main

import (
	"log"
	"os"

	"github.com/Eizeed/vibe_gogo/db"
	"github.com/Eizeed/vibe_gogo/handlers"
	"github.com/Eizeed/vibe_gogo/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
    err := godotenv.Load(".env")
    if err != nil {
		log.Fatal("error: failed to load the env file")
	}

    db.InitDB();

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
        r.POST("/playlists", playlists.Create);
        r.PATCH("/playlists/:uuid/addtrack", playlists.AddTrack);
        r.PATCH("/playlists/:uuid/deletetrack", playlists.Delete);
        r.PUT("/playlists/:uuid", playlists.Update);
        r.PUT("/playlists/:uuid/visibility", playlists.ChangeVisibility);
        r.DELETE("/playlists/:uuid", playlists.Delete);
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
