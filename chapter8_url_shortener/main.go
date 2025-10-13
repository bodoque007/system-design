package main

import "fmt"


type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
}

type truck struct {
	vehicle
	fourWheel bool
}

type person struct {
	fname string
	sname string
	favFood []string
}

type transport interface {
	yorke()
}

func (t truck) yorke() {
	fmt.Println("Let down")
}

func (s sedan) yorke() {
	fmt.Println("Hanging around")
}

func letdown(t transport) {
	t.yorke()
}

func (p person) walk() string {
	return fmt.Sprintln(p.fname, "is walking")
}


func main() {
	/*
	1)
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	for i, val := range slice {
		fmt.Println(i, val)
	}
	2)
	mapita := map[string]int{"h" : 1,}
	fmt.Println(mapita)
	for _, v := range mapita {
		fmt.Println(v)
	}*/
	p1 := person{
		"James",
		"Sunderland",
		[]string{"Pizza"},
	}
	t1 := truck{
		vehicle{
			4,
			"red",
		},
		false,
	}
	t2 := sedan{
		vehicle{
			4,
			"red",
		},
	}
	letdown(t1)
	letdown(t2)
	fmt.Println(t1.doors)
	fmt.Println(p1.walk())
}