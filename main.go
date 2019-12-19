package main

import "fmt"

func main() {

	backend := Developer{"Bob", "Bilbo", 1000, 32}
	boss := Director{"Bob", "Baggins", 2000, 40}

	backend.FullName()
	boss.FullName()

}

type Employee interface {
	FullName()
}

type Developer struct {
	FirstName string
	LastName  string
	Income    float32
	Age 	  int
}
func (d  Developer) FullName() {
	fmt.Println("Developer ",d.FirstName," ",d.LastName)
}


type Director struct {
	FirstName string
	LastName  string
	Income    float32
	Age       int
}

func (d  Director) FullName() {
	fmt.Println("Director ",d.FirstName," ",d.LastName)
}



