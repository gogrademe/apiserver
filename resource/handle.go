package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

func Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

type Resource struct {
	Table string
}

func main() {
	resources := []types.Resource{
		bitTorrent.BitTorrentPlugin{},
		images.ImagePlugin{},
		discovery.DiscoveryPlugin{},
	}

	// Register a default handler which calls plugin functions.
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		path := req.URL.Path

		for _, resource := range resources {
			pattern := resource.Namespace()
			matched, err := regexp.MatchString(pattern, path)
			if err != nil {
				fmt.Println(err)
			}
			if matched {
				code, res := resource.Handle(*req)
				rw.WriteHeader(code)
				json.NewEncoder(rw).Encode(res)

				return
			}
		}
		// We didn't match a resource so let's 404!

		rw.WriteHeader(404)
		json.NewEncoder(rw).Encode("Not found")
		return

	})

}
