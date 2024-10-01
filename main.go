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

func (b *MyLibrary) AddBook(books ...SimBook) {

	for _, val := range books {
		var t Book
		t.Author = val.Author
		t.BookName = val.BookName

		//Обычный способ генерирование уникального символа
		/*
			t.Id = b.lastId
			intId, _ := strconv.Atoi(b.lastId)
			b.booksId[intId] = t
			intId++
			lId := strconv.Itoa(intId)
			b.lastId = lId
		*/

		// Способ через UUID
		///*
		t.Id = b.lastId
		var x = uuid.New().String()
		b.lastId = x
		//*/
		b.booksName[val.BookName] = t
		b.books = append(b.books, t)
	}
}

func (b *MyLibrary) Searcher(name string) (*Book, bool) {
	val, ok := b.booksName[name]
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
	var t Library
	t = NewLibrary("some address")
	t.AddBook(SimBook{"1984", "George Orwell"})
	fmt.Println(t.TakeBookByName("1984"))
	fmt.Println(1000000007)
}
