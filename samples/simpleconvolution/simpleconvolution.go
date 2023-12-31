package main

import (
	"flag"

	"gitlab.com/akita/mgpusim/v3/benchmarks/amdappsdk/simpleconvolution"
	"gitlab.com/akita/mgpusim/v3/samples/runner"
)

var widthFlag = flag.Uint("width", 512, "The width of the input matrix.")
var heightFlag = flag.Uint("height", 512, "The height of the input matrix.") //254
var maskSizeFlag = flag.Uint("mask-size", 3, "The size of the mask.")

func main() {
	flag.Parse()

	runner := new(runner.Runner).ParseFlag().Init()

	benchmark := simpleconvolution.NewBenchmark(runner.Driver())
	benchmark.Height = uint32(*heightFlag)
	benchmark.Width = uint32(*widthFlag)
	benchmark.SetMaskSize(uint32(*maskSizeFlag))

	runner.AddBenchmark(benchmark)

	runner.Run()
}
