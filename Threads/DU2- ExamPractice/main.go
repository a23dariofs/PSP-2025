package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

const NUM_WORKERS = 5

type Task struct {
	ExpressionStr string
}

type Result struct {
	WorkerID int
	Expr     string
	Output   string
}

type Stats struct {
	sync.Mutex
	NumTasks        int
	NumOperations   map[string]int
	NumOpePerThread map[int]int
}

func parseExpression(expression string) (int, string, int, error) {
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("Invalid format")
	}

	op1, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Invalid number")
	}

	operator := parts[1]

	if !validateOperator(operator) {
		return 0, "", 0, fmt.Errorf("Unrecognized operator")
	}

	op2, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Invalid number")
	}

	return op1, operator, op2, nil
}

func validateOperator(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

func evaluate(op1 int, operator string, op2 int) (string, error) {
	switch operator {
	case "+":
		return strconv.Itoa(op1 + op2), nil
	case "-":
		return strconv.Itoa(op1 - op2), nil
	case "*":
		return strconv.Itoa(op1 * op2), nil
	case "/":
		if op2 == 0 {
			return "undefined", nil
		}
		return strconv.Itoa(op1 / op2), nil
	default:
		return "", fmt.Errorf("Unrecognized operator")
	}
}

func worker(id int, tasks <-chan Task, results chan<- Result, stats *Stats, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		op1, operator, op2, err := parseExpression(task.ExpressionStr)

		stats.Lock()
		stats.NumTasks++
		stats.NumOpePerThread[id]++
		stats.Unlock()

		if err != nil {
			stats.Lock()
			stats.NumOperations["Error"]++
			stats.Unlock()

			results <- Result{id, task.ExpressionStr, err.Error()}
			continue
		}

		value, err := evaluate(op1, operator, op2)
		if err != nil {
			stats.Lock()
			stats.NumOperations["Error"]++
			stats.Unlock()

			results <- Result{id, task.ExpressionStr, err.Error()}
			continue
		}

		stats.Lock()
		stats.NumOperations[operator]++
		stats.Unlock()

		results <- Result{id, task.ExpressionStr, value}
	}
}

func main() {
	inputFile, err := os.Open("expressions.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("results.txt")
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}
	defer outputFile.Close()

	stats := Stats{
		NumOperations:   make(map[string]int),
		NumOpePerThread: make(map[int]int),
	}

	tasks := make(chan Task, 100)
	results := make(chan Result, 100)

	var wg sync.WaitGroup

	for i := 0; i < NUM_WORKERS; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &stats, &wg)
	}

	go func() {
		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				tasks <- Task{ExpressionStr: line}
			}
		}
		close(tasks)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	writer := bufio.NewWriter(outputFile)

	for res := range results {
		fmt.Fprintf(writer, "(Worker %d) Expression: \"%s\" = %s\n",
			res.WorkerID, res.Expr, res.Output)
	}

	writer.Flush()

	printStats(stats)
	fmt.Println("End of main thread.")
}

func printStats(stats Stats) {
	stats.Lock()
	defer stats.Unlock()

	fmt.Println("\nStats:")
	fmt.Println("------")
	fmt.Printf("Number of expressions processed: %d\n\n", stats.NumTasks)

	fmt.Println("Number of expressions per thread:")
	fmt.Println("--------------------------------")
	for i := 0; i < NUM_WORKERS; i++ {
		fmt.Printf("Thread %d: %d\n", i, stats.NumOpePerThread[i])
	}

	fmt.Println("\nPercentage of operations:")
	fmt.Println("-------------------------")

	for _, op := range []string{"+", "-", "*", "/", "Error"} {
		count := stats.NumOperations[op]
		percent := float64(count) * 100 / float64(stats.NumTasks)
		fmt.Printf("%s: (%.2f%%)\n", op, percent)
	}
}
