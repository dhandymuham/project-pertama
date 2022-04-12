package main

import (
	"fmt"
	"sync"
)

type employee struct {
	name     string
	division string
}

type display interface {
	getName() string
	getDivision() string
}

func newEmployee(name, division string) display {
	return &employee{
		name:     name,
		division: division,
	}
}

func (e *employee) getName() string {
	return e.name
}
func (e *employee) getDivision() string {
	return e.division
}

func main() {
	emps := []display{
		newEmployee("Tumpal", "PM"),
		newEmployee("Parhehean", "FE"),
		newEmployee("Parlindungan", "QA"),
		newEmployee("Karpidol", "BE"),
		newEmployee("Juminten", "BE"),
	}

	var wg sync.WaitGroup

	wg.Add(len(emps))
	for _, emp := range emps {
		go show(emp, &wg)
	}

	wg.Wait()

}

func show(emp display, wg *sync.WaitGroup) {
	fmt.Println(emp.getName(), "\t", emp.getDivision())
	wg.Done()
}
