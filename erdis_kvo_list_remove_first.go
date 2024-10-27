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
)


func list_remove_first(msg *nats.Msg){
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

   valueToRemove := string(msg.Data)

   kv, err := js.KeyValue(ctx, bucket)
   if err != nil {
	   println("error is :: ")
	   fmt.Println(err)
	   msg.Respond([]byte(err.Error()))
	   return
   }

   // get value 
   entry, _ := kv.Get(ctx, key)
   if err != nil {
           println("error is :: ")
           fmt.Println(err)
           msg.Respond([]byte(err.Error()))
           return
   }

   // string to string array
   s := strings.Split(string(entry.Value()), ",")

   // get occurences in the slice of the value to remove
   ptr_indices_valueToRemove := list_find(&valueToRemove, s)

   // remove value(s)
   fmt.Println(s)
   s = removeIndicatedValueFromSlice(*ptr_indices_valueToRemove, s)
   //s = &d_s
   fmt.Println(s)

   // back to string
   l_string := strings.Join(s,",")


   // set value     TODO : error if version has changed (use update instead of put)?
   kv.Put(ctx, key, []byte(l_string) )

   // return status
   msg.Respond([]byte("all done!!!"))


}


func removeIndicatedValueFromSlice(ptr_indices_valueToRemove []int, s []string) []string{

	fmt.Printf("%d",ptr_indices_valueToRemove)
	idx := ptr_indices_valueToRemove[0]
	println(idx)
        s = append(s[:idx], s[idx+1:]...)
        return s
}




