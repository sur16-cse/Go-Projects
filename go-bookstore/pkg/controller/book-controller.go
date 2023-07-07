package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/surbhi/go-bookstore/pkg/models"
	"github.com/surbhi/go-bookstore/pkg/utils"
)

func GetBook(w http.ResponseWriter, r *http.Request)  {
	newBooks:=models.GetAllBooks()
	res,_:=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request)  {
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		log.Fatal("error while parsing")
	}
	bookDetails,_:=models.GetBookById(ID)
	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	//createBook := &models.Book{} creates a new instance of models.Book struct using the & operator to obtain its memory address. 
	//This will be used to hold the book data extracted from the request.
	createBook:=&models.Book{}
	utils.ParseBody(r,createBook)
	fmt.Println(createBook.Author,createBook.Name,createBook.Publication)
	b:=createBook.CreateBook()
	res,_:=json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter,r *http.Request)  {
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		log.Fatal("error while parsing")
	}
	book:=models.DeleteBook(ID)
	res,_:=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter,r *http.Request)  {
	var updateBook=&models.Book{}
	utils.ParseBody(r,updateBook)
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	booksDetails,db:=models.GetBookById(ID)
	if updateBook.Name!=""{
		booksDetails.Name=updateBook.Name
	}
	if updateBook.Author!=""{
		booksDetails.Author=updateBook.Author
	}
	if updateBook.Publication!=""{
		booksDetails.Publication=updateBook.Publication
	}
	db.Save(&booksDetails)
	res,_:=json.Marshal(booksDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}