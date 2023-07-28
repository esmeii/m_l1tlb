package main

import (
	"flag"

	"gitlab.com/akita/mgpusim/v3/benchmarks/shoc/fft"
	"gitlab.com/akita/mgpusim/v3/samples/runner"
)

var mb = flag.Int("MB", 64, "data size (in megabytes)") //8
var passes = flag.Int("passes", 16, "number of passes")

func main() {
	flag.Parse()

	runner := new(runner.Runner).ParseFlag().Init()

	benchmark := fft.NewBenchmark(runner.Driver())
	benchmark.Bytes = int32(*mb)
	benchmark.Passes = int32(*passes)

	runner.AddBenchmark(benchmark)

	runner.Run()
}
