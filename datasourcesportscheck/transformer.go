package datasourcesportscheck

import (
	"encoding/json"
	"io"
	//"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)


var lastUpdated time.Time

func ParseEvents(path string) chan *Product {
	rv := make(chan *Product)

	go func() {
		defer close(rv)

		var product *Product
		var input io.ReadCloser

		if strings.HasPrefix(path, "http://") ||
			strings.HasPrefix(path, "https://") {
			// HTTP Based
			log.Printf("Loading events from URL: %s", path)
			resp, err := http.Get(path)
			if err != nil {
				log.Println(err)
				return
			}
			if resp.StatusCode != 200 {
				log.Println("bad GET status for "+path, resp.Status)
				return
			}
			lastUpdated = time.Now()
			input = resp.Body
		} else {
			// File Based
			log.Printf("Loading events from file: %s", path)
			file, err := os.Open(path)
			if err != nil {
				log.Println(err)
				return
			}
			input = file
		}
		defer input.Close()

		// Load the JSON
		products, err := getProduct(input)
		if err != nil {
			log.Println(err)
			return
		}
        
        // Loop, and for each item push back onto recieve channel
		for _, item := range products {
            product = new(Product)
			*product = item
            rv <- product
		}
		

	}()

	return rv
}

func getProduct(file io.ReadCloser) ([]Product, error) {
    
    
	var p []Product

	d := json.NewDecoder(file)
	for {
		err := d.Decode(&p)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error decoding: %v", err)
			return nil, err
		}
	}

	return p, nil

}


