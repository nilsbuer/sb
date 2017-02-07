package main

import (
    "testing"
    "fmt"
)

// from commission.go
func BenchmarkCommission(b *testing.B) {
        // run the commission function b.N times
        fmt.Println("############# Benchmark Testing in progress  #########################")
        for n := 0; n < b.N; n++ {
                f_commission(100,0.1)
}

}

 
