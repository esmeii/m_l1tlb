package main

import (
	"flag"

	"gitlab.com/akita/mgpusim/v3/benchmarks/amdappsdk/fastwalshtransform"
	"gitlab.com/akita/mgpusim/v3/samples/runner"
)

var length = flag.Int("length", 524288, "The length of the array that will be transformed") //1024

func main() {
	flag.Parse()

	runner := new(runner.Runner).ParseFlag().Init()

	benchmark := fastwalshtransform.NewBenchmark(runner.Driver())
	benchmark.Length = uint32(*length)

	runner.AddBenchmark(benchmark)

	runner.Run()
}
