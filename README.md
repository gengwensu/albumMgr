# musicAlbum
RESTful API to manage a music Album

musicAlbum reads in catalog XML files into memory and provide RESTful API's to search the catalog by artist, song title, and album id. All output will be in JSON.

# API

1. GET  /musicalbum
  
    returns "Music Album management service"
    
2. GET  /musicalbum/artist

    returns all artists sorted alphabetically
    
3. GET  /musicalbum/artist?name="Taylor Swift"

    returns all albums authored by the artist
    
