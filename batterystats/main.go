package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/solarnz/batterystats"
	"os"

	"time"
)

func main() {
	var outfile = flag.String("out", "", "Where to write the output")
	var oneshot = flag.Bool("one", false, "Only collect battery statistics once")
	var period = flag.Int64("period", 10, "How often in minutes to collect battery statistics")
	flag.Parse()

	collect(*outfile)
	if !*oneshot {
		for _ = range time.Tick(time.Duration(*period) * time.Minute) {
			collect(*outfile)
		}
	}
}

func collect(outfile string) {
	batteries, err := batterystats.GetBatteries()
	if err != nil {
		fmt.Errorf("%s", err)
		os.Exit(1)
	}

	var f *os.File

	if outfile == "" {
		f = os.Stdout
	} else {
		f, err = os.OpenFile(outfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Errorf("%s", err)
			os.Exit(1)
		}
	}

	for _, battery := range batteries {
		data, err := json.Marshal(battery)
		if err != nil {
			fmt.Errorf("%s", err)
			os.Exit(1)
		}
		f.Write(data)
		f.WriteString("\n")
	}

	if outfile != "" {
		f.Close()
	}
}
