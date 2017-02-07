// Commission Calculation APP
// Description:
// Calculates the commissions based on an input file sale_data.txt
// 
// History
// 20.01.2017   Nils Buer    Initial Version
//

package main

import (
    "bufio"
    "encoding/csv"
    "os"
    "fmt"
    "io"
    "strconv"
)


// Commission Value
var c_comfactor float64 = 0.15
var v_commission float64

// Function to calculate the commission values
func f_commission (v_commvalue float64, v_comfactor float64) float64 {
	return v_commvalue * v_comfactor
}


func main() {

// Environment Variables "Not yet used"

fmt.Println("Using the Following Varibales:")
fmt.Println("ENV_TYPE = ", os.Getenv("ENV_TYPE"))
fmt.Println("ENV_DATABASE_NAME = ", os.Getenv("ENV_DATABASE_NAME"))
fmt.Println("ENV_DATABASE_PORT = ", os.Getenv("ENV_DATABASE_PORT"))
fmt.Println("ENV_DATABASE_PASSWORD = ", os.Getenv("ENV_DATABASE_PASSWORD"))
fmt.Println("ENV_DATABASE_USER = ", os.Getenv("ENV_DATABASE_USER"))
fmt.Println("ENV_DATABASE_SERVER = ", os.Getenv("ENV_DATABASE_SERVER"))
fmt.Println("")
fmt.Println("############# Commission Report #########################")

    // Load a TXT file.
    f, _ := os.Open("sales_data.txt")

    // Create a new reader.
    r := csv.NewReader(bufio.NewReader(f))
    for {
        record, err := r.Read()
        // Stop at EOF.
        if err == io.EOF {
            break
        }
        // Print each Element in the CSV file
        // Convert input string to float
        v_totalsales, _ := strconv.ParseFloat(record[1], 64)
        v_commission = f_commission(v_totalsales,c_comfactor)
        fmt.Println("Name: ",record[0]," # Number of Machines Sold: ",record[1]," # Commission: ",v_commission,"€")
    }
fmt.Println("############# Commission Report Finished ################")

}
