package main

import "fmt"

type Books struct {
	title string
	bookId uint
}
func main() {
	var myBook Books
	myBook.title = "标题"
	myBook.bookId = 1
	//fmt.Printf( "myBook title : %s\n", myBook.title)
	//fmt.Printf( "myBook id : %d\n", myBook.book_id)
	printBook(&myBook)
	print(myBook)
}

//打印
func printBook(book *Books)  {
	fmt.Print(book)
	fmt.Print("\n")
	fmt.Print(&book)
	fmt.Print("\n")
	fmt.Print(&book.title)
	fmt.Print("\n")
	fmt.Printf( "myBook title : %s\n", book.title)
	fmt.Printf( "myBook id : %d\n", book.bookId)
}

func print(book Books)  {
	fmt.Print(book)
	fmt.Print("\n")
	fmt.Print(&book)
	fmt.Print("\n")
	fmt.Print(&book.title)
	fmt.Print("\n")
	fmt.Printf( "myBook title : %s\n", book.title)
	fmt.Printf( "myBook id : %d\n", book.bookId)
}