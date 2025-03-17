package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"git.toolsfdg.net/be/liquid-mesh-svm-standalone/go/util"
)

func main() {
	fmt.Println("Simulating program...")

	// Read accounts.json file
	accountsJSON, err := os.ReadFile("./accounts.json")
	if err != nil {
		log.Fatalf("Failed to read accounts.json file: %v", err)
		os.Exit(1)
	}

	// Read tx.json file
	txJSON, err := os.ReadFile("./tx.json")
	if err != nil {
		log.Fatalf("Failed to read tx.json file: %v", err)
		os.Exit(1)
	}

	// Read program.so file and convert to base64
	programSoPath := "./program.so"
	programSoData, err := os.ReadFile(programSoPath)
	if err != nil {
		log.Fatalf("Failed to read program.so file: %v", err)
		os.Exit(1)
	}
	programSoBase64 := base64.StdEncoding.EncodeToString(programSoData)

	programSoPath2 := "./program2.so"
	programSoData2, err := os.ReadFile(programSoPath2)
	if err != nil {
		log.Fatalf("Failed to read program.so file: %v", err)
		os.Exit(1)
	}
	programSoBase642 := base64.StdEncoding.EncodeToString(programSoData2)

	// Use new API - Create program info array
	programs := []util.ProgramInfo{
		{
			ProgramID:       "HuTkmnrv4zPnArMqpbMbFhfwzTR7xfWQZHH1aQKzDKFZ",
			ProgramSoBase64: programSoBase64,
		},
		{
			ProgramID:       "HuTkmnrv4zPnArMqpbMbFhfwzTR7xfWQZHH1aQKzDKF1",
			ProgramSoBase64: programSoBase642,
		},
	}

	// Call the new simulation function
	result := util.SimulateTransactionWithPrograms(
		programs,
		string(accountsJSON),
		string(txJSON),
	)

	// Process results
	if result.Success {
		fmt.Println("Program simulation succeeded!")
	} else {
		fmt.Println("Program simulation failed!")
	}

	// Process and print logs
	if result.Logs != nil {
		var logs []string
		if err := json.Unmarshal(result.Logs, &logs); err == nil {
			fmt.Println("\nSimulation logs:")
			for _, log := range logs {
				fmt.Println(log)
			}
		} else {
			// If unable to parse as string array, print raw JSON
			fmt.Println("\nSimulation logs (raw):")
			fmt.Println(string(result.Logs))
		}
	}
}
