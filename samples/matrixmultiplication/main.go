package main

import (
	"flag"
	"log"
	"net/http"

	_ "net/http/pprof"

	"gitlab.com/akita/gcn3/benchmarks/amdappsdk/matrixmultiplication"
	"gitlab.com/akita/gcn3/samples/runner"
)

var xFlag = flag.Uint("x", 64, "The height of the first matrix.")
var yFlag = flag.Uint("y", 64, "The width of the first matrix and the height of the second matrix.")
var zFlag = flag.Uint("z", 64, "The width of the second matrix.")

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	flag.Parse()

	runner := new(runner.Runner).ParseFlag().Init()

	benchmark := matrixmultiplication.NewBenchmark(runner.GPUDriver)
	benchmark.X = uint32(*xFlag)
	benchmark.Y = uint32(*yFlag)
	benchmark.Z = uint32(*zFlag)

	runner.AddBenchmark(benchmark)

	runner.Run()
}
