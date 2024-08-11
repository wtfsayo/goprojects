package main

import (
    "fmt"
	"os"
	"time"
	"encoding/csv"
	"github.com/mergestat/timediff"
)

type Todo struct {
    id int
    text string
	completed bool
	createdAt time.Time
}


func createNew(curr *[]Todo, new string) {
	*curr = append(*curr, Todo{
		id: len(*curr) + 1, 
		text: new, 
		completed: false, 
		createdAt: time.Now()})
}

func deleteOne(curr *[]Todo, id int) {
	for _, todo := range *curr {
		if todo.id == id {
			*curr = append((*curr)[:todo.id], (*curr)[todo.id+1:]...)
			return
		}
	}

	fmt.Println("Todo not found for", id)

}

func listTODOs(curr *[]Todo) {
    for _, todo := range *curr {
		status := "pending"
		if todo.completed {
			status = "completed"
		}
        fmt.Println(todo.id, todo.text, status, timediff.TimeDiff(todo.createdAt))
    }
}

func printDivider() {
	fmt.Println("--------------------------------------------------------------------------------")
}


func loadCSV(fileName string) *os.File {
	file, err := os.Open(fileName)
    if err != nil {
        file, err = os.Create(fileName)
		return file
    }

	return file
	
}

func saveCSV(file *os.File, todos *[]Todo) {

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"id", "text", "completed", "createdAt"}
	w.Write(header)
	w.Flush()

	for _, todo := range *todos {
		w.Write([]string{fmt.Sprint(todo.id), todo.text, fmt.Sprint(todo.completed), fmt.Sprint(todo.createdAt)})
		w.Flush()
	}
}

func main() {
    
	todos := new([]Todo)

	csv := loadCSV("todos.csv")

	createNew(todos, "Buy eggs")
	createNew(todos, "Buy milk")
    createNew(todos, "Buy eggs")
    createNew(todos, "Buy bread")
	createNew(todos, "Buy raspberries")
	createNew(todos, "Buy bananas")
	

	// todo: fix existing todos file etc

	saveCSV(csv, todos)


	listTODOs(todos)
}