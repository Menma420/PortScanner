package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
   "time"
) 

func main(){
   fmt.Println("mini-portscanner:ready")

   target := flag.String("target","","target IP address")
   ports := flag.String("ports","80","ports to check")
   concurrency := flag.Int("concurrency", 100, "threads")
   timeout := flag.Float64("timeout", 1.5, "timeout")
   output := flag.String("output","table","table or json")
   confirm := flag.Bool("confirm", false, "confirm")

   flag.Parse()

   if *target == "" {
      fmt.Println("Error: --target required")
      flag.Usage()
      os.Exit(1)
   } 

   if *output!="table" && *output!="json"{
      fmt.Println("Error: --type needs to be table or json")
      os.Exit(1)
   }

   fmt.Println("Config:")
   fmt.Println("  target:", *target)
   fmt.Println("  ports:", *ports)
   fmt.Println("  concurrency:", *concurrency)
   fmt.Println("  timeout:", *timeout)
   fmt.Println("  output:", *output)
   fmt.Println("  confirm:", *confirm)
   fmt.Println("  tail:", flag.Args())

   trimmedPorts := strings.TrimSpace(*ports)
   Ports := getPorts(trimmedPorts)

   fmt.Println("Ports to scan:", Ports)

   targets, err := resolveTargets(*target)
   if err != nil {
      fmt.Println("Error resolving target:", err)
      os.Exit(1)
   }
   fmt.Println("Targets to scan:", targets)


   for _, p := range Ports{
      checkTCP(*target, p, time.Duration(*timeout*float64(time.Second)))
   }

}
