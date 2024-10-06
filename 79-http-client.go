package main

import (
	"bufio"
	"fmt"
	"net/http"
)

/*
---> Http Client <---
The Go standard library comes with excellent support for
"HTTP clients and servers"
in the net/http package.

In this example we’ll use it to issue simple HTTP requests.
*/

func learn_HTTP_client() {
	fmt.Println("\n------------ learn_HTTP_client --------------")

	/*
		Issue an HTTP GET request to a server. http.Get
		is a convenient shortcut around creating an

		http.Client object and calling its Get method;
		it uses the http.DefaultClient object which has useful default settings.
	*/
	resp, err := http.Get("https://bogatimahesh.com.np")
	fmt.Println("response: ", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("Response status:", resp.Status)

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// func main() {
// 	learn_HTTP_client()
// 	println("\n-------------------------------")
// }
