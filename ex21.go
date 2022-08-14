package main

import "fmt"

/*
	Реализовать паттерн «адаптер» на любом примере
*/

type Phone interface {
	ConnectInJack()
}
// использует интерфейс phone для реализации задач
type User struct {
}

func (u *User) ConnectJackToPhone(phone Phone) {
	fmt.Println("User connect headphone in jack")
	phone.ConnectInJack()
}

type Samsung struct {
}

func (s *Samsung) ConnectInJack() {
	fmt.Println("Headphne is connect to jack")
}

// у данной структуры нет юзер-френдли метода для реализации его потребностей
type Xiaomi struct {
}

func (x *Xiaomi) ConnectToTypeC() {
	fmt.Println("Headphne is connect to type-C")
}
// поэтому создаем адаптер который позволяет работать с объектами xiaomi как phone
type XiaomiAdapter struct {
	x *Xiaomi
}

func (xa *XiaomiAdapter) ConnectInJack() {
	fmt.Println("Adapter converts type-c to jack")
	xa.x.ConnectToTypeC()
}


func main() {
	user := &User{}
	sams := &Samsung{}

	user.ConnectJackToPhone(sams)
	xiaomi := &Xiaomi{}
	xAdapter := &XiaomiAdapter{
		x: xiaomi,
	}
	user.ConnectJackToPhone(xAdapter)
}
