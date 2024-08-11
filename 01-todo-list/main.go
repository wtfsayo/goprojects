package main

import (
    "fmt"
	"time"
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
		text: new, completed: false, 
		createdAt: time.Now()})
}

func deleteOne(curr *[]Todo, id int) {
	for _, todo := range *curr {
		if todo.id == id {
			*curr = append((*curr)[:todo.id], (*curr)[todo.id+1:]...)
			return
		}
	}

	fmt.Println("Todo not found for %d", id)

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

func main() {
    
	todos := new([]Todo)

	createNew(todos, "Buy eggs")
	createNew(todos, "Buy milk")
    createNew(todos, "Buy eggs")
    createNew(todos, "Buy bread")
	

	printDivider()

	listTODOs(todos)

	printDivider()

	deleteOne(todos, 2)
	
	printDivider()

	deleteOne(todos, 5)
    
	printDivider()

	listTODOs(todos)
}