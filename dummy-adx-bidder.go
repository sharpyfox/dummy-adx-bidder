package main

import (
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sharpyfox/dummy-adx-bidder/realtime_bidding"
	"github.com/sharpyfox/dummy-adx-bidder/utils"
	"log"
	"net/http"
	"strconv"
)

func main() {
	portPtr := flag.Int("p", 7040, "port to start http server")
	showVersion := flag.Bool("version", false, "print version string")
	flag.Parse()

	if *showVersion {
		fmt.Println(utils.Version("dummy-adx-bidder"))
		return
	}

	http.HandleFunc("/auctions", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/octet-stream")

		resp := &realtime_bidding.BidResponse{ProcessingTimeMs: proto.Int32(0)}
		data, err := proto.Marshal(resp)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}

		w.Write(data)
	})
	log.Println("INFO  Starting server on 0.0.0.0:" + strconv.Itoa(*portPtr))
	http.ListenAndServe(":"+strconv.Itoa(*portPtr), nil)
}
