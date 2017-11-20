# albumMgr
RESTful API to manage a musical Album

albumMgr reads in catalog XML files into memory and provide RESTful API's to search the catalog by artist, song title, and album id. All output will be in JSON.

# API

1. GET  /albummgr
  
    returns "Musical Album management service"
    
2. GET  /albummgr/album

    returns all album sorted by album id

```
$ curl http://localhost:8081/albummgr/album                                        
results: [                                                                         
 {                                                                                 
  "Id": 0,                                                                         
  "Artist": "Kenny Chesney",                                                       
  "Title": "The Boys of Fall"                                                      
 },                                                                                
 {                                                                                 
  "Id": 1,                                                                         
  "Artist": "Kenny Chesney",                                                       
  "Title": "Live a Little"                                                         
 },                                                                                
 {                                                                                 
  "Id": 2,                                                                         
  "Artist": "Kenny Chesney",                                                       
  "Title": "Coastal"                                                               
 },                                                                                
 {                                                                                 
  "Id": 3,                                                                         
  "Artist": "Kenny Chesney feat. Grace Potter",                                    
  "Title": "You and Tequila"                                                       
 },
 ...                                                                                
 }                                                                                 
]                                                                                  
```
    
3. GET  /albummgr/album?artist="Taylor Swift"&title="Look What You Made Me Do"

    returns albums the search criteria

     Example: 


    ```
    $ curl http://localhost:8081/albummgr/album?artist="Kenny%20Chesney"\&title="Live%20a%20Little"
results: [                                                                                     
 {                                                                                             
  "Id": 1,                                                                                     
  "Artist": "Kenny Chesney",                                                                   
  "Title": "Live a Little"                                                                     
 }                                                                                             
]                                                                                              
    ```
    
4. GET  /albummgr/{$albumId}

    returns the album with the albumId

```
$ curl http://localhost:8081/albummgr/album/3     
results: [                                        
 {                                                
  "Id": 3,                                        
  "Artist": "Kenny Chesney feat. Grace Potter",   
  "Title": "You and Tequila"                      
 }                                                
]                                                 
```
    
The service should respond with 404 to all other requests not listed above

# environment & build
 require Go
  
$ go build ../src/github.com/gengwensu/albumMgr/albumMgr.go 

$./albumMgr.exe &

