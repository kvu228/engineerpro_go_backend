package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ResponseData struct {
	Status  string     `json:"status"`
	Data    []Employee `json:"data"`
	Message string     `json:"message"`
}

type Employee struct {
	Id             int    `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
	EmployeeAge    int    `json:"employee_age"`
	ProfilePicture string `json:"profile_picture"`
}

func getEmployees(url string) ([]Employee, error) {
	var response ResponseData
	var resp *http.Response
	var err error

	// Set initial delay and max retry count for exponential backoff
	delay := 1 * time.Second
	maxRetries := 5

	for retries := 0; retries < maxRetries; retries++ {
		// Make a GET request
		resp, err = http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to make GET request: %w", err)
		}

		if resp.StatusCode == http.StatusOK {
			// Successful response
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("failed to read response body: %w", err)
			}

			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
			}
			var employees []Employee
			for i := range response.Data {
				employees = append(employees, response.Data[i])
			}

			return employees, nil

		} else if resp.StatusCode == http.StatusTooManyRequests {
			// Handle HTTP 429 by implementing exponential backoff
			fmt.Println("Received 429 Too Many Requests. Retrying...")
			time.Sleep(delay)
			delay *= 2 // Exponential backoff: double the delay time
			continue

		} else {
			// Handle other HTTP errors
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
		resp.Body.Close()

	}

	return nil, fmt.Errorf("max retries reached; unable to fetch employees")
}

func worker(id int, jobs <-chan Employee, results chan<- Employee) {
	for j := range jobs {
		fmt.Printf("Worker %v: Processing Employee id:%d\tname: %v\tSalary/age: %f\n", id, j.Id, j.EmployeeName, float64(j.EmployeeSalary/j.EmployeeAge))
		fmt.Printf("Finished calculate for employee_id: %d\n", j.Id)
		results <- j
	}
}

func main() {
	// Fetch data from URL
	url := "https://dummy.restapiexample.com/api/v1/employees"
	employees, err := getEmployees(url)
	if err != nil {
		fmt.Printf("error fetching employees: %v", err)
	}

	// print data
	for _, employee := range employees {
		fmt.Printf("ID: %d, Name: %s, Salary: %d, Age: %d, Profile Picture: %s\n", employee.Id, employee.EmployeeName, employee.EmployeeSalary, employee.EmployeeAge, employee.ProfilePicture)
	}

	jobs := make(chan Employee, len(employees))
	results := make(chan Employee, len(employees))
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	for _, employee := range employees {
		jobs <- employee
	}
	close(jobs)

	for a := 1; a <= len(employees); a++ {
		<-results
	}
}
