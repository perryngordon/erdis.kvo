package main

import (
  "fmt"
  "strings"
  "os"
  "context"
  "time"
  "github.com/nats-io/nats.go"
  "github.com/nats-io/nats.go/jetstream"
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
   }


   // get value / TODO create is not present ?
   entry, _ := kv.Get(ctx, key)
   if err != nil {
           println("error is :: ")
           fmt.Println(err)
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



