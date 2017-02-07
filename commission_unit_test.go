package main
import (
        "testing"
        "fmt"
)

func TestAverage(t *testing.T) {
var v float64

fmt.Println ("## Automatic Unit Testing using GO testing package")
v = f_commission(100,0.1)
if v != 10 {
t.Error("Expected 10, got ", v)
}
}
