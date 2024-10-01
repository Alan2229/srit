package main

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type Book struct {
	Id       string
	BookName string
	Author   string
}

type TakeBook interface {
	TakeBookById(id string) Book
	TakeBookByName(name string) Book
	TakeBooksByName(name ...string) []Book
	TakeBooksById(id ...string) []Book
}

type Searcher interface {
	Search(id string) (*Book, bool)
}

type SimBook struct {
	BookName string
	Author   string
}

type Library interface {
	Searcher
	TakeBook
	Search(name string) (*Book, bool)
	AddBook(books ...SimBook)
}

type MyLibrary struct {
	books     []Book
	booksName map[string]Book
	booksId   map[string]Book
	lastId    string
	address   string
}

func NewLibrary(Address string) *MyLibrary {
	return &MyLibrary{
		books:     make([]Book, 0),
		booksName: make(map[string]Book),
		booksId:   make(map[string]Book),
		//обычный способ
		//lastId:    "0",
		//через UUID
		lastId:  uuid.NewString(),
		address: Address,
	}
}

func (b *MyLibrary) TakeBookById(Id string) Book {
	_, ok := b.Search(Id)
	if !ok {
		fmt.Println("Book with this Id does not exist")
		fmt.Println("Do you want to try again with the correct Id?")
		fmt.Println("Y/n (Yes/No)")
		k := ""
		fmt.Scan(&k)
		if k == "Y" {
			q := ""
			fmt.Scan(&q)
			return b.TakeBookById(q)
		} else {
			return Book{}
		}
	} else {
		for i := 0; i < len(b.books); i++ {
			if b.books[i].Id == Id {
				w := b.books[i]
				b.books[i] = b.books[len(b.books)-1]
				b.books[len(b.books)-1] = Book{}
				b.books = b.books[:len(b.books)-1]
				delete(b.booksId, Id)
				delete(b.booksName, b.books[i].BookName)
				return w
			}
		}
	}
	return Book{}
}

var tp int = 0

func (b *MyLibrary) AddBook(books ...SimBook) {

	for _, val := range books {
		var t Book
		t.Author = val.Author
		t.BookName = val.BookName
		if tp == 0 {
			//Обычный способ генерирование уникального символа
			t.Id = b.lastId
			intId, _ := strconv.Atoi(b.lastId)
			b.booksId[b.lastId] = t
			intId++
			lId := strconv.Itoa(intId)
			b.lastId = lId
		} else {
			// Способ через UUID
			t.Id = b.lastId
			var x = uuid.New().String()
			b.lastId = x
		}
		b.booksName[val.BookName] = t
		b.books = append(b.books, t)
	}
}

func (b *MyLibrary) Searcher(name string) (*Book, bool) {
	val, ok := b.booksName[name]
	fmt.Println(val, " + ", name)
	if !ok {
		return nil, false
	}
	return &val, true
}

func (b *MyLibrary) Search(id string) (*Book, bool) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, false
	}
	val, ok := b.booksId[strconv.Itoa(intId)]
	if !ok {
		return nil, false
	}
	return &val, true
}

func (b *MyLibrary) TakeBookByName(name string) Book {
	t, ok := b.Searcher(name)
	if ok {
		for i := 0; i < len(b.books); i++ {
			if b.books[i].BookName == name {
				w := t.Id
				b.books[i] = b.books[int(len(b.books))-1]
				b.books[len(b.books)-1] = Book{}
				b.books = b.books[:int(len(b.books))-1]
				delete(b.booksId, w)
				delete(b.booksName, t.BookName)
				break
			}
		}
		return *t
	}
	return Book{}
}

func (b *MyLibrary) TakeBooksByName(names ...string) []Book {
	var ans []Book
	for i := 0; i < len(names); i++ {
		ans = append(ans, b.TakeBookByName(names[i]))
	}
	return ans
}

func (b *MyLibrary) TakeBooksById(ids ...string) []Book {
	var ans []Book
	for i := 0; i < len(ids); i++ {
		ans = append(ans, b.TakeBookById(ids[i]))
	}
	return ans
}

func main() {
	// Создаем слайс книг
	books := []SimBook{
		{"1984", "Джордж Оруэлл"},
		{"Убить пересмешника", "Харпер Ли"},
		{"Великий Гэтсби", "Ф. Скотт Фицджеральд"},
		{"Гордость и предубеждение", "Джейн Остин"},
	}

	// Создаем библиотеку
	lib := NewLibrary("Некоторый адрес")

	// Загружаем книги в библиотеку
	lib.AddBook(books...)

	// Ищем 1-2 книги в библиотеке
	fmt.Println("Поиск книги '1984':")
	book1, found := lib.Searcher("1984")
	if found {
		fmt.Println("Найдена:", *book1)
	} else {
		fmt.Println("Книга не найдена")
	}

	fmt.Println("Берем книгу 'Великий Гэтсби':")
	takenBook := lib.TakeBookByName("Великий Гэтсби")
	if (takenBook != Book{}) {
		fmt.Println("Взята:", takenBook)
	} else {
		fmt.Println("Книга не найдена или уже взята")
	}
	// Заменяем функцию генерации ID на UUID
	fmt.Println("\nЗаменяем функцию генерации ID на UUID")
	lib.lastId = uuid.NewString()
	tp = 1
	// Добавляем больше книг
	newBooks := []SimBook{
		{"Моби Дик", "Герман Мелвилл"},
		{"Война и мир", "Лев Толстой"},
	}
	lib.AddBook(newBooks...)

	// Ищем еще одну книгу в библиотеке
	fmt.Println("\nПоиск книги 'Моби Дик':")
	book2, found := lib.Search("Моби Дик")
	if found {
		fmt.Println("Найдена:", *book2)
	} else {
		fmt.Println("Книга не найдена")
	}
}
