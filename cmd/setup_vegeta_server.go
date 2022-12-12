package cmd

import (
	"es_load_test/config"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func StartVegetaServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "attack",
		Short:   "Start HTTP API server",
		Aliases: []string{"vegeta"},
		Run: func(_ *cobra.Command, _ []string) {
			err := config.Load()
			if err != nil {
				log.Fatalf("Config load failed: \n%s", err.Error())
			}

			for _, scenario := range config.Cfg.GetESScenarios() {
				for _, rate := range scenario.Rates {
					name := fmt.Sprintf("%s-%d-%d-%d", scenario.Case, rate, scenario.Duration, scenario.MaxWorkers)
					fmt.Printf("Attacking %s with rate %d - duration %d - max workers %d\n", scenario.Url, rate, scenario.Duration, scenario.MaxWorkers)
					r := vegeta.Rate{Freq: rate, Per: time.Second}
					targeter := vegeta.NewStaticTargeter(vegeta.Target{
						Method: "POST",
						URL:    scenario.Url,
					})
					attacker := vegeta.NewAttacker(vegeta.Timeout(time.Duration(scenario.Duration)*time.Second), vegeta.MaxWorkers(uint64(scenario.MaxWorkers)))

					var metrics vegeta.Metrics
					metrics.Histogram = &vegeta.Histogram{}
					if err := metrics.Histogram.Buckets.UnmarshalText([]byte("[0,100ms,200ms,300ms]")); err != nil {
						panic(err)
					}
					for res := range attacker.Attack(targeter, r, time.Duration(scenario.Duration)*time.Second, name) {
						metrics.Add(res)
					}
					metrics.Close()

					hdrHistogramReporter(metrics, name)
					histogramReporter(metrics, name)
					jsonReporter(metrics, name)
					txtReporter(metrics, name)

					fmt.Printf("Result (Total Request: %d):\n", metrics.Requests)
					fmt.Printf("Success rate: %.0f%% \r\n", metrics.Success*100)
					fmt.Printf("Status Codes[code:count]: %v\r\n\n", metrics.StatusCodes)

					fmt.Printf("\n==========================\n")

					time.Sleep(3 * time.Minute)
				}
			}
		},
	}
}

func hdrHistogramReporter(metrics vegeta.Metrics, name string) {
	plot := vegeta.NewHDRHistogramPlotReporter(&metrics)
	f, err := os.Create(fmt.Sprintf("%s.hdr", name))
	if err != nil {
		panic(err)
	}
	err = plot.Report(f)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func histogramReporter(metrics vegeta.Metrics, name string) {
	hist := vegeta.NewHistogramReporter(metrics.Histogram)
	f, err := os.Create(fmt.Sprintf("%s.hist", name))
	if err != nil {
		panic(err)
	}
	err = hist.Report(f)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func jsonReporter(metrics vegeta.Metrics, name string) {
	hist := vegeta.NewJSONReporter(&metrics)
	f, err := os.Create(fmt.Sprintf("%s.json", name))
	if err != nil {
		panic(err)
	}
	err = hist.Report(f)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func txtReporter(metrics vegeta.Metrics, name string) {
	hist := vegeta.NewTextReporter(&metrics)
	f, err := os.Create(fmt.Sprintf("%s.txt", name))
	if err != nil {
		panic(err)
	}
	err = hist.Report(f)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
