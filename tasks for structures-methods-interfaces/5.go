package main

import "fmt"

//5. Интерфейс Stringer
//Задача: Реализуйте интерфейс fmt.Stringer для структуры Book (с полями Title, Author),
//чтобы при печати выводилось: "Книга: {Title}, Автор: {Author}".

type Book struct {
	Title  string
	Author string
}

func (b *Book) String() string {
	return fmt.Sprintf("Книга: %s, Автор: %s", b.Title, b.Author)
}
func main() {
	bok := Book{"Go Programming Language", "Marat Vorobei"}
	fmt.Println(bok.String())
}
