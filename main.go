package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type ImprovementSummary struct {
	QuestionId         string `json:"QuestionId"`
	PillarId           string `json:"PillarId"`
	QuestionTitle      string `json:"QuestionTitle"`
	Risk               string `json:"Risk"`
	ImprovementPlanUrl string `json:"ImprovementPlanUrl"`
}

type Workload struct {
	WorkloadId           string               `json:"WorkloadId"`
	LensAlias            string               `json:"LensAlias"`
	LensArn              string               `json:"LensArn"`
	ImprovementSummaries []ImprovementSummary `json:"ImprovementSummaries"`
}

// A helper function that returns the directory where the code is running
func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
	}
	return dir
}

func main() {
	// JSON path
	fmt.Print("Enter the full path to the JSON file (high_risk_improvements.json): ")
	var jsonFilePath string
	fmt.Scanln(&jsonFilePath)

	// open JSON file
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	var workload Workload
	jsonParser := json.NewDecoder(jsonFile)
	err = jsonParser.Decode(&workload)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	csvFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()
	err = writer.Write([]string{"Card Name", "Card Description", "Labels", "List Name", "Checklist", "Checklist item"})
	if err != nil {
		fmt.Println("Error writing CSV header:", err)
		return
	}

	for _, summary := range workload.ImprovementSummaries {
		row := []string{
			summary.PillarId + "-" + summary.QuestionTitle, // Combine PillarId and QuestionTitle for Card Name
			summary.ImprovementPlanUrl,                     // ImprovementPlanUrl for Card Description
			summary.Risk,
			summary.PillarId,
			"", // Empty as there's no corresponding field in your JSON
			"", // Empty as there's no corresponding field in your JSON
		}
		err = writer.Write(row)
		if err != nil {
			fmt.Println("Error writing CSV row:", err)
			return
		}
	}
}
