package main

import (
	"fmt"
	"sort"
)

type Data struct {
	Name   string
	Age    int
	Salary int
}

type ByNameSalary []Data

func (a ByNameSalary) Len() int {
	return len(a)
}
func (a ByNameSalary) Less(i, j int) bool {

	if a[i].Name < a[j].Name {
		return true
	}
	if a[i].Name > a[j].Name {
		return false
	}
	return a[i].Salary < a[j].Salary

}
func (a ByNameSalary) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {

	n := []Data{}
	n = append(n, Data{Name: "A", Age: 35, Salary: 500})
	n = append(n, Data{Name: "BB", Age: 30, Salary: 400})
	n = append(n, Data{Name: "A", Age: 26, Salary: 9005})
	n = append(n, Data{Name: "D", Age: 29, Salary: 1500})
	n = append(n, Data{Name: "B", Age: 29, Salary: 1000})
	n = append(n, Data{Name: "F", Age: 29, Salary: 3000})

	sort.Sort(ByNameSalary(n))

	for _, v := range n {
		fmt.Printf("Name : %s  Age : %d  Salary : %d\n", v.Name, v.Age, v.Salary)
	}

}
