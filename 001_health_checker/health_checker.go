package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)


var httpClient = &http.Client{
		Timeout: 5 * time.Second,
}

func handleErrors(err error) (int, error){
	if err != nil{
		return 0, err
	}
}

func checkURL(url string) int {
	response, err := httpClient.Get(url)
	handleErrors(err)
	defer response.Body.Close()
	return response.StatusCode
}

func main(){
	fmt.Println("Starting programm...")
	fmt.Println("Reading URLs...")

	
	data, err := os.Open("urls.txt")	

	handleErrors(err)
	defer data.Close()

	var urlsList []string
	var linesCounter int = 0

	scanner := bufio.NewScanner(data)

	for scanner.Scan(){
		linesCounter += 1

		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#"){
			continue
		}

		urlsList = append(urlsList, line)
	}

	for _, url := range urlsList {
		responseCode := checkURL((url))
		fmt.Printf("Response code for %s is %d\n", url, responseCode)
		
	}



	fmt.Printf("Processed %d lines\n", linesCounter)



	

}	






