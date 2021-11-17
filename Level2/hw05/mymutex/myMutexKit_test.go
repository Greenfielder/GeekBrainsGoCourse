package mymutex

import (
	"math/rand"
	"sync"
	"testing"
)

func BenchmarkMutextKitR10W90(t *testing.B) {
	s := NewMutexKit()
	MutexKitRWTest(s, t.N, 10, 90)
}

func BenchmarkRWMutextKitR10W90(t *testing.B) {
	s := NewRWMutexKit()
	MutexKitRWTest(s, t.N, 10, 90)
}

func BenchmarkMutextSetR50W50(t *testing.B) {
	s := NewMutexKit()
	MutexKitRWTest(s, t.N, 50, 50)
}

func BenchmarkRWMutextSetR50W50(t *testing.B) {
	s := NewRWMutexKit()
	MutexKitRWTest(s, t.N, 50, 50)
}

func BenchmarkMutextSetR90W10(t *testing.B) {
	s := NewMutexKit()
	MutexKitRWTest(s, t.N, 90, 10)
}

func BenchmarkRWMutextSetR90W10(t *testing.B) {
	s := NewRWMutexKit()
	MutexKitRWTest(s, t.N, 90, 10)
}

func MutexKitRWTest(s Kit, iterNum int, readOps int, writeOps int) {

	wg := &sync.WaitGroup{}

	for i := 0; i < iterNum; i++ {
		for writeOp := 0; writeOp < writeOps; writeOp++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, number float64) {
				defer wg.Done()
				s.Add(number)
			}(wg, rand.Float64())
		}

		for readOp := 0; readOp < readOps; readOp++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, number float64) {
				defer wg.Done()
				s.Has(number)
			}(wg, rand.Float64())
		}
	}
	wg.Wait()

}
