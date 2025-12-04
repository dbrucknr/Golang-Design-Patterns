package examples

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Enables hot-swapping of backend services
func SampleAdapterPatternUsage() {
	// Starting with No Adapter Pattern Usage
	todo := getRemoteTodo()
	fmt.Println("Todo without adapter pattern:\t", todo)

	// With JSON Adapter Pattern Usage
	jsonAdapter := &RemoteService{Remote: &JsonBackend{}}
	jsonTodo, err := jsonAdapter.CallRemoteService()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Todo with adapter pattern:\t", jsonTodo)

	// With XML Adapter Pattern Usage
	xmlAdapter := &RemoteService{Remote: &XmlBackend{}}
	xmlTodo, err := xmlAdapter.CallRemoteService()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Todo with XML adapter pattern:\t", xmlTodo)
}

// Enables hot-swapping of backend services
type DataInterface interface {
	GetData() (*Todo, error)
}

// Wrapper struct - Anything that satisfies the DataInterface, can be used as a backend service
type RemoteService struct {
	Remote DataInterface
}

func (rs *RemoteService) CallRemoteService() (*Todo, error) {
	// Since Remote is an interface we can put anything in that
	// satisfies the interface
	return rs.Remote.GetData()
}

type Todo struct {
	UserId    int    `json:"userId" xml:"userId"`
	Id        int    `json:"id" xml:"id"`
	Title     string `json:"title" xml:"title"`
	Completed bool   `json:"completed" xml:"completed"`
}

// What if this service notifies us about switching from JSON to XML?
func getRemoteTodo() *Todo {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Unexpected status code:", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)

	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Fatal(err)
	}
	return &todo
}

// Adapter Pattern:
type JsonBackend struct {
}

func (jb *JsonBackend) GetData() (*Todo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Unexpected status code:", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)

	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

// XML Backend:
type XmlBackend struct {
}

func (xb *XmlBackend) GetData() (*Todo, error) {
	// Fake the HTTP call with an XML stub
	xmlFile := `
		<?xml version="1.0" encoding="UTF-8"?>
		<root>
			<userId>1</userId>
			<id>1</id>
			<title>delectus aut autem</title>
			<completed>false</completed>
		</root>
	`
	var todo Todo
	err := xml.Unmarshal([]byte(xmlFile), &todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}
