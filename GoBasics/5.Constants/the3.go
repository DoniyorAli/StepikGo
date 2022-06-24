package main


const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	_  // пропускаем 7
	Add
)

fmt.Println(Sunday)   // вывод: 0
fmt.Println(Saturday) // вывод: 6
fmt.Println(Add) // вывод: 8
