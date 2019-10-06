package main

import (
	"flag"
	"os"

	"github.com/umiheikin/authlog-exporter/exporter"
)

var (
	promPort      = flag.String("promethues.port", "9100", "Port for prometheus metrics to listen on")
	promEndpoint  = flag.String("prometheus.endpoint", "/metrics", "Endpoint used for metrics")
	authPath      = flag.String("auth.path", "/var/log/auth.log", "Path to auth.log")
	debug         = flag.Bool("debug", false, "Run full scan on test logs file")
)

func main() {
	flag.Parse()

	exporter.SetDebugging(*debug)
	exporter.SetPrometheusEndpointAndPort(*promEndpoint, *promPort)

	if *authPath != "" {
		if _, err := exporter.LoadAuthLog(*authPath); err != nil {
			panic(err)
		}
	}

	// Start the file listeners
	exporter.Start()

	// Wait for kill signal
	k := make(chan os.Signal, 2)
	<-k

	// Shutdown all exporters gracefully
	exporter.Shutdown()
}
