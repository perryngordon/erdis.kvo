package main

import (
  "fmt"
  "strings"
  "os"
  "context"
  "time"
  "github.com/nats-io/nats.go"
  "github.com/nats-io/nats.go/jetstream"
  "slices"
)


func list_append(msg *nats.Msg){
   // make error
   url := os.Getenv("NATS_URL")
   if url == "" {
        url = nats.DefaultURL
   }

   nc, err := nats.Connect(url)
   if err != nil {
     println("problem connecting to NATS server, does NATS_URL need to be set? ")
   }
   defer nc.Drain()   

   js, _ := jetstream.New(nc)
   ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
   defer cancel()

   cmd_args := strings.Split(msg.Subject, ".")
   bucket := cmd_args[4]
   key_string := cmd_args[5:]
   key := strings.Join(key_string,".") 

   value := string(msg.Data)

   kv, err := js.KeyValue(ctx, bucket)
   if err != nil {
	   println("error is :: ")
	   fmt.Println(err)
           msg.Respond([]byte(err.Error()))
           return
   }

   // "get value" (below) is segfaulting if the key does not already exist (as of 241221)
   // so check it first and create it if it is does not exist
   keys, err := kv.Keys(ctx, nil)
   if ! slices.Contains(keys,key) {
     // create key
     kv.Put(ctx, key, []byte(""))
   } 

   // get value 
   entry, _ := kv.Get(ctx, key)
   if err != nil {
           println("error is :: ")
           fmt.Println(err)
           msg.Respond([]byte(err.Error()))
           return
   }

   // string to list 
   l := strings.Split(string(entry.Value()), ",")

   // append value
   l = append(l, value)

   // back to string
   l_string := strings.Join(l,",")

   // set value     TODO : error if version has changed ?
   kv.Put(ctx, key, []byte(l_string) )

   // return status
   msg.Respond([]byte("all done!!!"))
  

}



