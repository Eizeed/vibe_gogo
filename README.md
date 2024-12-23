A little pet project for me to learn GO and Gin with some utils like gorm orm

vibe gogo is a music service where you can create playlists with your favorite tracks

I can't integrate player for music because it's not free and also there are copyrights

In my API i have authorization with jwt tokens and cookies.
For tracks i'm using https://api.deezer.com
It gives metadata for tracks for free. With it's functions i don't have to store all tracks' metadata in my database.
There is visibility for playlist. User can set if other people will be able to see playlist.

Soon I'll create an alorithm for recommended tracks in playlists based on tracks' genre, related artists etc.

API:
    base_url: "http://localhost:8080/api"

    users:
        GET:    "/users"
        POST:   "/users/login"
        POST:   "/users/register"
        POST:   "/users/logout"
        PUT:    "/users/:uuid"
        DELTE:  "/users/:uuid"

    playlists:
        GET:    "/playlists/user/:uuid"
        GET:    "/playlists/:uuid"
        POST:   "/playlists"
        PATCH:  "/playlists/:uuid/addtrack"
        PATCH:  "/playlists/:uuid/deletetrack"
        PUT:    "/playlists/:uuid/visibility"
        PUT:    "/playlists/:uuid"
        DELETE: "/playlists/:uuid"

    search:
        GET:    "/deezer/search"
