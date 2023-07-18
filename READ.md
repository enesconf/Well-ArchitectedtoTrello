# AWS Well-Architected to Trello Tasks

This project automates the creation of Trello tasks from AWS Well-Architected workloads with high risk improvements.

## Prerequisites

1. AWS CLI: You will need to have the AWS CLI installed and configured with your AWS credentials. You can install it following these [instructions](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html).

2. `jq`: You will need `jq` installed to process JSON output from AWS CLI commands. You can install it following these [instructions](https://stedolan.github.io/jq/download/).

3. Go: The main script of this project is written in Go. Make sure you have Go installed on your machine. You can install it following these [instructions](https://golang.org/doc/install).

## Steps to Run the Project

1. List workloads:

    ```bash
    aws wellarchitected list-workloads
    ```

2. List high risk lens-review improvements for a specific workload:

    ```bash
aws wellarchitected list-lens-review-improvements --workload-id $(workload-id) --lens-alias wellarchitected --max-results 100 --output json  > improvements.json
    ```

3. Filter the improvements with high risk and save it to a new JSON file:

    ```bash
    jq '.ImprovementSummaries |= map(select(.Risk == "HIGH"))' improvements.json > high_risk_improvements.json
    ```

4. Run the Go script to generate a CSV file:

    ```bash
    go run main.go
    ```

The Go script (`main.go`) reads from the `high_risk_improvements.json` file, and writes the output to a CSV file named `output.csv`.

Please replace `$(workload-id)` with your actual workload id while executing the command.
