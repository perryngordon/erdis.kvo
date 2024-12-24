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


func list_remove_all(msg *nats.Msg){
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

   s := []string{}
   // "get value" (below) is segfaulting if the key does not already exist (as of 241221)
   keys, err := kv.Keys(ctx, nil)
   if ! slices.Contains(keys,key) {
      // no op

   }else{
     // get value 
     entry, _ := kv.Get(ctx, key)
     // string to list 
     s = strings.Split(string(entry.Value()), ",")

     if err != nil {
           println("error is :: ")
           fmt.Println(err)
           msg.Respond([]byte(err.Error()))
           return
     }
   }

   // get occurences in the slice of the value to remove
   ptr_indices_valueToRemove := list_find(&valueToRemove, s)

   // remove value(s)
   fmt.Println(s)
   s = removeAllOcurrencesOfTheIndicatedValueFromSlice(ptr_indices_valueToRemove, s)
   fmt.Println(s)

   // back to string
   l_string := strings.Join(s,",")

   if l_string == "" {
           // delete this key if it is now empty
           kv.Delete(ctx, key)
   }else{
     // set value     TODO : error if version has changed (use update instead of put)?
     kv.Put(ctx, key, []byte(l_string) )
   }


   // return status
   msg.Respond([]byte("all done!!!"))


}


func removeAllOcurrencesOfTheIndicatedValueFromSlice(ptr_indices_valueToRemove *[]int, s []string) []string{

  for i,v := range *ptr_indices_valueToRemove {

	  // if index is not 0 then decrement v
	  // v is the value of the index of the value to be removed
	  // once the first occurence of the value is removed, the indices
	  // of the remaining occurences has to be decremented, b/c
	  // the items shift left in the slice when an item is removed
	  if( i != 0){
           v = v - i
	  }


          s = append(s[:v], s[v+1:]...)

  }


  return  s
}




