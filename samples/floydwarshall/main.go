package main

import (
	"flag"

	"gitlab.com/akita/mgpusim/v3/benchmarks/amdappsdk/floydwarshall"
	"gitlab.com/akita/mgpusim/v3/samples/runner"
)
//0
var numNodes = flag.Int("node", 16, "The number of nodes in the graph")
var numIterations = flag.Int("iter", 1024, 
	`The number of iterations to run. If this value is set to 0 or a number
	larger than the number of nodes, it will be reset to the number of nodes.`)

func main() {
	flag.Parse()

	runner := new(runner.Runner).ParseFlag().Init()

	benchmark := floydwarshall.NewBenchmark(runner.Driver())
	benchmark.NumNodes = uint32(*numNodes)
	benchmark.NumIterations = uint32(*numIterations)

	runner.AddBenchmark(benchmark)

	runner.Run()
}
