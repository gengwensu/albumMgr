/*
Albummgr: a http server that reads in XML files and accepts HTTP GET requests
to display album info.
*/

package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type XMLProducts struct {
	XMLName xml.Name   `xml:"PRODUCTS"`
	Tracks  []XMLAlbum `xml:"TRACKBUNDLE>TRACKS>TRACK"`
}

type XMLAlbum struct {
	XMLName xml.Name `xml:"TRACK"`
	Artist  string   `xml:"PRODUCTINFO>LANGUAGE>ARTIST"`
	Title   string   `xml:"PRODUCTINFO>LANGUAGE>TITLE"`
}

type Album struct {
	Id     int
	Artist string
	Title  string
}

type Albums struct {
	AlbumCollection []Album
}

const MAXOUTPUT = 50
const EXT = "xml" // file extension to read in

func main() {
	fileDirPath := flag.String("dir", "./", "directory path for input XML files")
	flag.Parse()
	//fmt.Printf("directory for input is %s\n", *fileDirPath)
	files := []string{}
	filepath.Walk(*fileDirPath, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(EXT+"$", f.Name()) // match ext at end of line
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})

	// albums := Albums{}
	data := XMLProducts{}
	for _, f := range files {
		//fmt.Printf("file to read %s\n", f)
		dat, err := ioutil.ReadFile(f)
		if err != nil {
			panic(err)
		}

		err = xml.Unmarshal(dat, &data)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
	}

	ds := Albums{}
	ds.AlbumCollection = []Album{}
	for i, t := range data.Tracks {
		album := Album{}
		album.Id = i
		album.Artist = t.Artist
		album.Title = t.Title
		ds.AlbumCollection = append(ds.AlbumCollection, album)
	}
	log.Fatal(http.ListenAndServe("localhost:8081", ds))

}

func (db Albums) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/albummgr", "/albummgr/": //GET
		if req.Method == "GET" {
			fmt.Fprint(w, "Musical Album management service\n") // return signature of the service
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	case "/albummgr/album", "/albummgr/album/": //GET album
		if req.Method == "GET" {
			artist := req.URL.Query().Get("artist")
			title := req.URL.Query().Get("title")

			out := []Album{}
			count := 0
			for _, r := range db.AlbumCollection {

				if (r.Artist == artist || artist == "") &&
					(r.Title == title || title == "") {
					out = append(out, r)
					count++
					if count >= MAXOUTPUT { //limit the number of output
						break
					}
				}
			}

			dataout, err := json.MarshalIndent(out, "", " ")
			if err != nil {
				log.Fatalf("JSON marshaling failed: %s", err)
			}
			fmt.Fprintf(w, "results: %s\n", string(dataout))
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	default:
		if strings.Contains(req.URL.Path, "/albummgr/album/") {
			if req.Method == "GET" {
				sid := strings.TrimPrefix(req.URL.Path, "/albummgr/album/")
				id, err := strconv.Atoi(sid)
				if err != nil || id < 0 || id >= len(db.AlbumCollection) {
					http.Error(w, "Id not valid", http.StatusBadRequest)
				} else {
					out := []Album{}
					out = append(out, db.AlbumCollection[id])
					dataout, err := json.MarshalIndent(out, "", " ")
					if err != nil {
						log.Fatalf("JSON marshaling failed: %s", err)
					}
					fmt.Fprintf(w, "results: %s\n", string(dataout))
				}
			} else {
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			}
		} else {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "http %d, %s invalid. Only albummgr, albummgr/album, albummgr/album/{$Id) are allowed.\n",
				http.StatusMethodNotAllowed, req.URL)
		}
	}
}
