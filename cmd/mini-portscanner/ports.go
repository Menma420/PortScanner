package main

import (
	"fmt"
	"os"
	"sort"
   "strconv"
	"strings"
) 

func getPorts(ports string) []int {
   parts := strings.Split(ports, ",")

   var Ports []int

   for _, num := range parts{
      if strings.Contains(num, "-"){
         rangeParts := strings.Split(num, "-")
         if len(rangeParts)!=2{
            fmt.Println("Error: invalid port range:", num)
            os.Exit(1)
         
         }
         
         var start, end int

         startStr := strings.TrimSpace(rangeParts[0])
         endStr := strings.TrimSpace(rangeParts[1])

         start, err := strconv.Atoi(startStr)
         if err != nil{
            fmt.Println("Error: invalid port:", rangeParts[0])
            os.Exit(1)
         }
         end, err = strconv.Atoi(endStr)
         if err != nil{
            fmt.Println("Error: invalid port:", rangeParts[1])
            os.Exit(1)
         }
         if start<1 || start>65535 || end<1 || end>65535 || start>end{
            fmt.Println("Error: invalid port range:", num)
            os.Exit(1)
         }
         for i:=start;i<=end;i++{
            Ports = append(Ports,i)
         }

      }else{
         var port int 
         port, err := strconv.Atoi(strings.TrimSpace(num))
         if err != nil{
            fmt.Println("Error: invalid port:", num)
            os.Exit(1)
         }
         if port<1 || port>65535{
            fmt.Println("Error: invalid port:", num)
            os.Exit(1)
         }
         Ports = append(Ports,port)
      }
   }

   unique := make(map[int]struct{})
   deduped := []int{}

   for _, port := range Ports{
      if _, exists := unique[port]; !exists{
         deduped = append(deduped, port)
         unique[port] = struct{}{}
      }
   }

   Ports = deduped
   sort.Ints(Ports)

   return Ports
}