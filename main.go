package main

import (
	"log"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

func contextWithTotalTimeout(total_timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), total_timeout)
}
func main() {
	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)

	key := "/foo"
	opts := &client.GetOptions{Recursive: true}
	log.Print("Getting '" + key + "' key value")
	// opts = nil
	resp, err := kapi.Get(context.Background(), key, opts)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
}
