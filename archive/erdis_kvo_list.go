package main


import (
	"context"
	"fmt"
	"os"
	"time"


	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)


func main() {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}


	nc, _ := nats.Connect(url)
	defer nc.Drain()


	js, _ := jetstream.New(nc)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	kv, _ := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket: "profiles",
	})

	kv.Put(ctx, "sue.color", []byte("blue"))
	entry, _ := kv.Get(ctx, "sue.color")
	fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))


	kv.Put(ctx, "sue.color", []byte("green"))
	entry, _ = kv.Get(ctx, "sue.color")
	fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

	_, err := kv.Update(ctx, "sue.color", []byte("red"), 1)
	fmt.Printf("expected error: %s\n", err)


	kv.Update(ctx, "sue.color", []byte("red"), 2)
	entry, _ = kv.Get(ctx, "sue.color")
	fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

	name := <-js.StreamNames(ctx).Name()
	fmt.Printf("KV stream name: %s\n", name)


