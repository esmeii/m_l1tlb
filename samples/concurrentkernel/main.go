package main

import (
	"flag"

	_ "net/http/pprof"

	"gitlab.com/akita/mgpusim/v3/benchmarks/amdappsdk/bitonicsort"
	"gitlab.com/akita/mgpusim/v3/benchmarks/heteromark/fir"
	"gitlab.com/akita/mgpusim/v3/samples/runner"
)

func main() {
	flag.Parse()

	runner := new(runner.Runner).ParseFlag().Init()

	firBenchmark := fir.NewBenchmark(runner.Driver())
	firBenchmark.Length = 102400 //10240
	firBenchmark.SelectGPU([]int{1})

	bsBenchmark := bitonicsort.NewBenchmark(runner.Driver())
	bsBenchmark.Length = 512 //64
	bsBenchmark.SelectGPU([]int{1})

	runner.AddBenchmarkWithoutSettingGPUsToUse(firBenchmark)
	runner.AddBenchmarkWithoutSettingGPUsToUse(bsBenchmark)

	runner.Run()
}
