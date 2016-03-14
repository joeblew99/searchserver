package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/blevesearch/bleve"
	bleveHttp "github.com/blevesearch/bleve/http"
    
    data "bitbucket.org/apparently_different/backend/nodes/searchserver/datasourcesportscheck"
)

var indexPath = flag.String("indexpath", "fosdem.bleve", "index path")
var indexName = flag.String("indexname", "items", "index name")
var dataSourcePath = flag.String("datasourcepath", "fosdem.ical", "fosdem events ical path")
var bindAddr = flag.String("addr", ":8099", "http listen address")
var update = flag.Duration("update", 0, "update every")

var lastUpdated time.Time

func main() {

	flag.Parse()

	// turn on http request logging
	bleveHttp.SetLog(log.New(os.Stderr, "bleve.http", log.LstdFlags))

	// open index
	index, err := bleve.Open(*indexPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		// or create new if it doesn't exist
        
        // ! Here is the Mapper
		mapping := data.BuildMapping()
		index, err = bleve.New(*indexPath, mapping)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	// insert/update index in background
	go batchIndexEvents(index, *dataSourcePath)

	// update data periodically
	if *update > 0 {
		log.Printf("Updating every: %s", *update)
		ticker := time.NewTicker(*update)
		go func() {
			for _ = range ticker.C {
				log.Printf("Updating index now")
				go batchIndexEvents(index, *dataSourcePath)
			}
		}()
	}

	// start server
	startServer(index, *bindAddr)
}

func startServer(index bleve.Index, addr string) {
	// create a router to serve static files
	router := staticFileRouter()

	// add the API
    // ! "fosdem". Need a different one per Data Structure i guess.
	bleveHttp.RegisterIndexName(*indexName, index)
    
	searchHandler := bleveHttp.NewSearchHandler(*indexName)
	router.Handle("/api/search", searchHandler).Methods("POST")
    
	lastHandler := new(lastUpdatedHandler)
	router.Handle("/api/lastUpdated", lastHandler).Methods("GET")

	http.Handle("/", router)
	log.Printf("Listening on %v", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
    
}

func batchIndexEvents(index bleve.Index, path string) {
	count := 0
	batch := index.NewBatch()
    
    // ! Here is the Transformer
	for event := range data.ParseProduct(path) {
        // ! Data Type MUST have UID
		batch.Index(event.UID, event)
		if batch.Size() >= 100 {
			err := index.Batch(batch)
			if err != nil {
				log.Println(err)
				return
			}
			count += batch.Size()
			log.Printf("Indexed %d Items\n", count)
			batch = index.NewBatch()
		}
	}
	if batch.Size() > 0 {
		err := index.Batch(batch)
		if err != nil {
			log.Println(err)
			return
		}
		count += batch.Size()
	}
	log.Printf("Indexed %d Items\n", count)
}
