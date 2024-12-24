package main

import(
 "fmt"
 //"slices"
)

func list_find(valueToFind *string, s []string) *[]int{

   fmt.Println(s)
   foundValue := []int{}

   for i, v := range s {
      if (v == *valueToFind){
              foundValue = append(foundValue,i)
      }
      //fmt.Printf("%d",i)
      //fmt.Printf(v)
  } 

  fmt.Println(foundValue)
  return &foundValue

}
