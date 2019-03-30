package matrixmultiplication

import (
	"log"
	"math"
	"math/rand"

	"gitlab.com/akita/gcn3/driver"
)

type Benchmark struct {
	driver  *driver.Driver
	context *driver.Context
	gpus    []int

	X, Y, Z                   uint32
	MatrixA, MatrixB, MatrixC *Matrix
}

func NewBenchmark(driver *driver.Driver) *Benchmark {
	b := new(Benchmark)
	b.driver = driver
	b.context = driver.Init()
	return b
}

func (b *Benchmark) SelectGPU(gpus []int) {
	b.gpus = gpus
}

func (b *Benchmark) Run() {
	b.initMem()
	b.exec()
}

func (b *Benchmark) initMem() {
	rand.Seed(0)

	b.MatrixA = NewMatrix(b.X, b.Y)
	for i := uint32(0); i < b.X; i++ {
		for j := uint32(0); j < b.Y; j++ {
			b.MatrixA.Data[j*b.X+i] = rand.Float32()
			//b.MatrixA.Data[j*b.X+i] = float32(j*b.X + i)
		}
	}

	b.MatrixB = NewMatrix(b.Z, b.X)
	for i := uint32(0); i < b.Z; i++ {
		for j := uint32(0); j < b.X; j++ {
			b.MatrixB.Data[j*b.Z+i] = rand.Float32()
			//b.MatrixB.Data[j*b.Z+i] = float32(j*b.Z + i)
		}
	}
}

func (b *Benchmark) exec() {
	m := NewGPUMatrixMultiplier(b.driver, b.context)
	m.SelectGPU(b.gpus)
	b.MatrixC = m.Multiply(b.MatrixA, b.MatrixB)
}

func (b *Benchmark) Verify() {
	m := CPUMatrixMultiplier{}
	mCPU := m.Multiply(b.MatrixA, b.MatrixB)
	for i := uint32(0); i < mCPU.Width; i++ {
		for j := uint32(0); i < mCPU.Width; i++ {
			index := i + j*mCPU.Width

			if math.Abs(float64(mCPU.Data[index]-b.MatrixC.Data[index])) > 1e-3 {
				log.Panicf("mismatch at [%d, %d]: expected %f, but get %f",
					i, j, mCPU.Data[index], b.MatrixC.Data[index])
			}
		}
	}

	log.Print("Passed!")
}