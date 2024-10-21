package main

import (
  "fmt"
  "strings"
  "os"
  "context"
  "time"
  "github.com/nats-io/nats.go"
  "github.com/nats-io/nats.go/jetstream"
  "reflect"
)


func list_append(msg *nats.Msg){
   // make error
   url := os.Getenv("NATS_URL")
   if url == "" {
        url = nats.DefaultURL
   }

   nc, err := nats.Connect(url)
   if err != nil {
     println("problem here")
   }
   defer nc.Drain()   

   js, _ := jetstream.New(nc)
   ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
   defer cancel()

   cmd_args := strings.Split(msg.Subject, ".")
   bucket := cmd_args[4]
   key_string := cmd_args[5:]
   key := strings.Join(key_string,".") 
   println()
   value := string(msg.Data)

   println(bucket)
   println(key)
   println(value)

   kv, err := js.KeyValue(ctx, bucket)
   if err != nil {
	   println("error is :: ")
	   fmt.Println(err)
   }
   // get or create value
   entry, _ := kv.Get(ctx, key)
   println(" get kv is:")
   println(string(entry.Value()))
   // cast to list
   l := strings.Split(string(entry.Value()), ",")
   fmt.Println(reflect.TypeOf(l))
   // append value
   l = append(l, value)
   // cast to string
   l_string := strings.Join(l,",")
   println(l_string)
   // set value     TODO : error if version has changed ?
   // return status


}



