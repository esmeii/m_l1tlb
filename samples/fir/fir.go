package main

import (
	"flag"

	"gitlab.com/akita/mgpusim/v3/benchmarks/heteromark/fir"
	"gitlab.com/akita/mgpusim/v3/samples/runner"
)

var numData = flag.Int("length", 524288, "The number of samples to filter.") //4096

func main() {
	flag.Parse()

	runner := new(runner.Runner).ParseFlag().Init()

	benchmark := fir.NewBenchmark(runner.Driver())
	benchmark.Length = *numData

	runner.AddBenchmark(benchmark)

	runner.Run()
}
