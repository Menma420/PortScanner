package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"github.com/olekukonko/tablewriter"
)

type Job struct {
	IP   string
	Port int
}

type Result struct {
	IP      string
	Port    int
	Open    bool
	Latency time.Duration
	Err     error
}

func main() {
	fmt.Println("mini-portscanner:ready")

	target := flag.String("target", "", "target IP address")
	ports := flag.String("ports", "80", "ports to check")
	concurrency := flag.Int("concurrency", 100, "threads")
	timeout := flag.Float64("timeout", 1.5, "timeout")
	output := flag.String("output", "table", "table or json")
	confirm := flag.Bool("confirm", false, "confirm")

	flag.Parse()

	if *target == "" {
		fmt.Println("Error: --target required")
		flag.Usage()
		os.Exit(1)
	}

	if *output != "table" && *output != "json" {
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

	//fmt.Println("Ports to scan:", Ports)

	targets, err := resolveTargets(*target)
	if err != nil {
		fmt.Println("Error resolving target:", err)
		os.Exit(1)
	}
	//fmt.Println("Targets to scan:", targets)

	if !*confirm {
		fmt.Println("pls confirm with --confirm to start scan")
		os.Exit(0)
	}

	// convert timeout float to time.Duration once
	timeoutDur := time.Duration(*timeout * float64(time.Second))

	//making Jobs slice
	Jobs := []Job{}
	for _, t := range targets {
		for _, p := range Ports {
			Jobs = append(Jobs, Job{IP: t, Port: p})
		}
	}
	var totalJobs = len(Jobs)
	fmt.Println("Total jobs to scan:", totalJobs)

	//Job and results channels
	jobs := make(chan Job, totalJobs)
	results := make(chan Result, totalJobs)

	//starting workers
	var wg sync.WaitGroup

	for w := 1; w <= *concurrency; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				open, lat, err := checkTCP(job.IP, job.Port, time.Duration(timeoutDur))
				results <- Result{IP: job.IP, Port: job.Port, Open: open, Latency: lat, Err: err}
			}
		}(w)
	}

	go func() {
		for _, job := range Jobs {
			jobs <- job
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	//collecting results
	Results := make([]Result, 0, totalJobs)
	for result := range results {
		Results = append(Results, result)
	}

	// At this point Results contains all scanned results
	fmt.Printf("Scan complete: %d results\n", len(Results))

	if *output == "json" {
		jsonData, err := json.MarshalIndent(Results, "", "  ")
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
	}

	if *output == "table" {
		table := tablewriter.NewWriter(os.Stdout)
		// Append header row as first row
		table.Append([]string{"IP", "Port", "Open", "Latency", "Error"})

		for _, r := range Results {
			table.Append([]string{
				r.IP,
				fmt.Sprintf("%d", r.Port),
				fmt.Sprintf("%v", r.Open),
				r.Latency.String(),
				fmt.Sprintf("%v", r.Err),
			})
		}
		table.Render()
	}
}
