# albumMgr
RESTful API to manage a musical Album

albumMgr reads in catalog XML files into memory and provide RESTful API's to search the catalog by artist, song title, and album id. All output will be in JSON.

# API

1. GET  /albummgr
  
    returns "Musical Album management service"
    
2. GET  /albummgr/album

    returns all album sorted by album id
    
3. GET  /albummgr/album?artist="Taylor Swift"&title="Look What You Made Me Do"

    returns albums the search criteria
    
4. GET  /albummgr/{$albumId}

    returns the album with the albumId
    
