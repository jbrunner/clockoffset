package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpserver := flag.String(
		"ntpserver",
		"",
		"ntp server hostname",
	)
	format := flag.String(
		"format",
		"ms",
		"output formats:\n"+
			"  s    seconds\n"+
			"  ms   miliseconds\n"+
			"  us   microseconds\n"+
			"  h    human readable\n",
	)
	quiet := flag.Bool(
		"quiet",
		false,
		"suppress output to standard output",
	)
	limit := flag.Int64(
		"limit",
		0,
		"exit(2) if offset diff is greather than <n> ms",
	)
	sleep := flag.Int(
		"sleep",
		0,
		"sleep for <n> ms if the limit was exceeded",
	)

	flag.Parse()

	if *format != "s" && *format != "ms" && *format != "us" && *format != "h" {
		log.Fatal("format must be one of: {s,ms,us,h}")
	}
	if *ntpserver == "" {
		log.Fatal("ntpserver is required")
	}

	response, err := ntp.Query(*ntpserver)
	if err != nil {
		log.Fatal(err)
	}

	if *quiet == false {
		switch {
		case *format == "s":
			fmt.Println(response.ClockOffset.Seconds())
		case *format == "ms":
			fmt.Println(response.ClockOffset.Milliseconds())
		case *format == "us":
			fmt.Println(response.ClockOffset.Microseconds())
		case *format == "h":
			fmt.Println("Clock offset is:", response.ClockOffset.String())
		}
	}

	ms := response.ClockOffset.Milliseconds()
	if ms < 0 {
		ms = -ms
	}
	if *limit > 0 && ms > *limit {
		if *sleep > 0 {
			time.Sleep(time.Duration(*sleep) * time.Millisecond)
		}
		os.Exit(2)
	}
}
