package week1

import (
	"fmt"
	"github.com/santiaago/caltechx.go/pla"
	"runtime"
	"time"
)

type experiment struct {
	NRuns             int     // number of times we repeat the experiment
	NPoints           int     // number of in sample points to learn
	SumOfIterations   int     // sum of iterations taken to converge through all the runs of the experiment
	SumOfDisagreement float64 // sum of disagreement between learned function g and target function f though all the runs of the experiment
}

// print will show:
// The average of iteration taken by the experiment to converge.
// The average of disagreement between g the learned function (h function with Wn vector) and f the target function.
func (exp *experiment) print() {
	avgIterations := exp.SumOfIterations / exp.NRuns
	avgDisagreement := exp.SumOfDisagreement / float64(exp.NRuns)
	fmt.Printf("average for PLA to converge for N = %v is %v\n", exp.NPoints, avgIterations)
	fmt.Printf("average of disagreement between the hypothesis function and the random function for N=%v is %4.2f\n", exp.NPoints, avgDisagreement)
}

func (exp *experiment) avgIterations() int {
	return exp.SumOfIterations / exp.NRuns
}

func (exp *experiment) avgDisagreement() float64 {
	return exp.SumOfDisagreement / float64(exp.NRuns)
}

// measure will measure the time taken by function f to run and display it.
func measure(f func() experiment, name string) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("%s took %4.2f seconds\n", name, elapsed.Seconds())
}

// Take N = 10. How many iterations does it take on average for the PLA to
// converge for N = 10 training points?
func q7() experiment {
	exp := experiment{NRuns: 1000, NPoints: 10}
	pla := pla.NewPLA()

	for run := 0; run < exp.NRuns; run++ {
		pla.Initialize()
		iterations := pla.Converge()
		exp.SumOfDisagreement += pla.Disagreement()
		exp.SumOfIterations += iterations
	}
	exp.print()
	return exp
}

func q7cc() experiment {
	exp := experiment{NRuns: 1000, NPoints: 10}
	pla := pla.NewPLA()

	chanDisagreement := make(chan float64, exp.NRuns)
	chanIterations := make(chan int, exp.NRuns)

	for run := 0; run < exp.NRuns; run++ {
		go func() {
			pla.Initialize()
			iterations := pla.Converge()
			chanDisagreement <- pla.Disagreement()
			chanIterations <- iterations
		}()
	}
	for i := 0; i < exp.NRuns; i++ {
		exp.SumOfDisagreement += <-chanDisagreement
	}
	for i := 0; i < exp.NRuns; i++ {
		exp.SumOfIterations += <-chanIterations
	}

	exp.print()
	return exp
}

func q9() experiment {
	exp := experiment{NRuns: 1000, NPoints: 100}
	pla := pla.NewPLA()

	for run := 0; run < exp.NRuns; run++ {
		pla.N = exp.NPoints
		pla.Initialize()
		iterations := pla.Converge()
		exp.SumOfDisagreement += pla.Disagreement()
		exp.SumOfIterations += iterations
	}
	exp.print()
	return exp
}

func q9cc() experiment {
	exp := experiment{NRuns: 1000, NPoints: 100}
	pla := pla.NewPLA()

	chanDisagreement := make(chan float64, exp.NRuns)
	chanIterations := make(chan int, exp.NRuns)

	for run := 0; run < exp.NRuns; run++ {
		go func() {
			pla.N = exp.NPoints
			pla.Initialize()
			iterations := pla.Converge()
			chanDisagreement <- pla.Disagreement()
			chanIterations <- iterations
		}()
	}

	for i := 0; i < exp.NRuns; i++ {
		exp.SumOfDisagreement += <-chanDisagreement
	}
	for i := 0; i < exp.NRuns; i++ {
		exp.SumOfIterations += <-chanIterations
	}

	exp.print()
	return exp
}

func main() {
	fmt.Println("Num CPU: ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("week 1")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	measure(q7, "q7")
	measure(q7cc, "q7 concurrent")
	fmt.Println("8")
	fmt.Println("9")
	measure(q9, "q9")
	measure(q9cc, "q9 concurrent")
	fmt.Println("10")
}
