package main

import (
	"fmt"
	"unicode"
)

type Person struct {
	Age    int
	Status string
	Name   string
}

func (m *Person) GetAge() int {
	if m.Age < 0 || m.Age > 150 {
		fmt.Println("Возраст должен быть в диапазоне от 0 до 150")
		m.Age = 0
	}
	return m.Age
}
func (m *Person) SetAge(Age int) {
	m.Age = Age
}

func (m *Person) GetName() string {
	for _, char := range m.Name {
		if unicode.IsDigit(char) {
			fmt.Println("В имени не должно быть чисел")
			m.Name = "0"
		}
	}
	return m.Name
}

func (m *Person) SetName(Name string) {
	m.Name = Name
}
func (m *Person) GetStatus() string {
	return m.Status
}
func (m *Person) SetStatusByAge() {
	if m.Age < 10 {
		m.Status = "Ребенок"
	} else if m.Age < 18 {
		m.Status = "Подросток"
	} else if m.Age >= 18 && m.Age < 65 {
		m.Status = "Взрослый"
	} else {
		m.Status = "Пожилой"
	}
}
func NewPerson() []Person {
	var people []Person
	for {
		var name string
		var age int
		fmt.Println("Введите имя или введите exit для выхода:")
		fmt.Scanln(&name)
		if name == "exit" {
			break
		}
		fmt.Println("Введите возраст:")
		fmt.Scanln(&age)
		person := Person{
			Name: name,
			Age:  age,
		}
		people = append(people, person)
	}
	return people
}
func Calculate(people []Person) int {
	sum := 0
	for _, person := range people {
		sum += person.Age
	}
	if len(people) > 0 {
		return sum / len(people)
	}
	return 0
}
func TryAdd(people *[]Person, newPerson Person) bool {
	for _, person := range *people {
		if person.Age == newPerson.Age && person.Name == newPerson.Name {
			return false
		}

	}
	*people = append(*people, newPerson)
	return true
}

func main() {
	people := NewPerson()

	newPerson := Person{Name: "Dave", Age: 35}

	fmt.Println("Список добавленных людей:")
	for _, person := range people {
		fmt.Printf("Имя: %s, Возраст: %d\n", person.Name, person.Age)
	}

	if TryAdd(&people, newPerson) {
		fmt.Println("Люди успешно добавлены!")
	} else {
		fmt.Println("Человек уже существует!")
	}

	averageAge := Calculate(people)
	fmt.Println("Средний:", averageAge)

}
