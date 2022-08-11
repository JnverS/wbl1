package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) printName() {
	fmt.Println(h.name)
}

func (h Human) sleep() {
	fmt.Println(h.name, "спит 8 часов в сутки")
}

type Action struct {
	Human
	name string
}

func (a Action) printName() {
	fmt.Println(a.name)
}

func main() {
	action := Action{Human{name: "Иван", age: 30}, "Job"}

	fmt.Printf("Value: %#v\n", action)

	// выводим имя встроенной структуры Human
	fmt.Println(action.Human.name)

	// выводим поле имя, принадлежащее самой структуре action
	fmt.Println(action.name)

	// если есть переопределение метода, для того чтобы
	// обратиться к встроенному методу нужно прописать "полный" путь
	action.printName()
	action.Human.printName()

	// если переопределения нет, возможна скоращенная запись
	action.sleep()
}
