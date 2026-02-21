package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Creating file
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	// Writing in the file
	f.WriteString("Testing os.Create")

	// Reading a file
	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	// Opening file
	file2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	// Reader of the file
	reader := bufio.NewReader(file2)
	// Creating buffer
	buffer := make([]byte, 10)
	for {
		// Reading from a buffer
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}

	// Remove file
	os.Remove("test.txt")

	// --------------------------------

	// GET request
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	// Defer is a reserved word to close resources when it's done. You can put it "first" as it
	// will clone when the operation is done
	defer req.Body.Close()
	fmt.Printf("\nreq is of type %T\n", req)
	//
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// --------------------------------

	// Struct with json tags to "bind" the attribute to the key in the JSON
	type Account struct {
		Name    string `json:"Name"`
		Balance int    `json:"Balance"`
	}

	account := Account{
		Name:    "Jane",
		Balance: 150,
	}
	// Marshalling to JSON
	acc, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(acc))
	// Encoder that will throw the output in the standard output (in this case the terminal)
	// Then Encode the account variable
	if err := json.NewEncoder(os.Stdout).Encode(account); err != nil {
		panic(err)
	}

	pureJson := []byte(`{"Name": "Jane","Balance": 150}`)
	var accountBack Account
	// Unmarshal (from JSON to Account type)
	json.Unmarshal(pureJson, &accountBack)
	fmt.Println(accountBack.Balance)
}
