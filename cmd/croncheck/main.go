package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/user/croncheck/internal/describer"
	"github.com/user/croncheck/internal/parser"
	"github.com/user/croncheck/internal/simulator"
	"github.com/user/croncheck/internal/visualizer"
)

func main() {
	count := flag.Int("n", 5, "number of next runs to preview")
	showTimeline := flag.Bool("timeline", false, "show ASCII timeline")
	showTable := flag.Bool("table", true, "show run table")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: croncheck [flags] \"<cron expression>\"\n\n")
		fmt.Fprintf(os.Stderr, "Examples:\n")
		fmt.Fprintf(os.Stderr, "  croncheck \"0 9 * * 1-5\"\n")
		fmt.Fprintf(os.Stderr, "  croncheck -n 10 -timeline \"*/15 * * * *\"\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	expr := flag.Arg(0)

	sched, err := parser.Parse(expr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid cron expression: %v\n", err)
		os.Exit(1)
	}

	desc, err := describer.Describe(expr)
	if err == nil {
		fmt.Printf("Expression : %s\n", expr)
		fmt.Printf("Description: %s\n\n", desc)
	}

	now := time.Now()
	runs := simulator.NextRuns(sched, now, *count)

	if *showTable {
		fmt.Println(visualizer.Table(runs, now))
	}

	if *showTimeline {
		window := time.Duration(0)
		if len(runs) > 1 {
			window = runs[len(runs)-1].Sub(now) + 5*time.Minute
		} else {
			window = time.Hour
		}
		fmt.Println(visualizer.Timeline(runs, window))
	}
}
