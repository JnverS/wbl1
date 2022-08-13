package main

import "fmt"

/*
	Реализовать паттерн «адаптер» на любом примере
*/

type Admin interface {
	AddProgram()
}
// использует интерфейс admin для реализации задач
type User struct {
}

func (u *User) AddNewProgramOnPC(admin Admin) {
	fmt.Println("User tried install program on pc")
	admin.AddProgram()
}

type Mac struct {
}

func (m *Mac) AddProgram() {
	fmt.Println("programm installed on mac")
}

// у данной структуры нет юзер-френдли метода для реализации его потребностей
type Linux struct {
}

func (l *Linux) AddProgramOnLinux() {
	fmt.Println("Tadam!")
}
// поэтому создаем адаптер который позволяет работать с объектами linux как admin
type LinuxAdapter struct {
	l *Linux
}

func (la *LinuxAdapter) AddProgram() {
	fmt.Println("Adapter что-то там шаманит")
	la.l.AddProgramOnLinux()
}


func main() {
	user := &User{}
	mac := &Mac{}

	user.AddNewProgramOnPC(mac)
	linux := &Linux{}
	lAdapter := &LinuxAdapter{
		l: linux,
	}
	user.AddNewProgramOnPC(lAdapter)
}
