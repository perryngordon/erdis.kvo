package main

import (
  "fmt"
  "strings"
  "os"
  "context"
  "time"
  "github.com/nats-io/nats.go"
  "github.com/nats-io/nats.go/jetstream"
  //"container/list"
  //"reflect"
  //"errors"
  "slices"
)


func kv_delete(msg *nats.Msg){
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


   kv, err := js.KeyValue(ctx, bucket)
   if err != nil {
	   println("error is :: ")
	   fmt.Println(err)
	   msg.Respond([]byte(err.Error()))
	   return
   }

   
   // "get value" (below) is segfaulting if the key does not already exist (as of 241221)
   keys, err := kv.Keys(ctx, nil)
   if ! slices.Contains(keys,key) {
      // no op
      msg.Respond([]byte("no op"))
      return
   }else{
     // delete key 
     err := kv.Delete(ctx, key)

     if err != nil {
           println("error / kv delete key :: ")
           fmt.Println(err)
           msg.Respond([]byte(err.Error()))
           return
     }
   }





   // return status
   msg.Respond([]byte("kv delete completed .. "))


}






