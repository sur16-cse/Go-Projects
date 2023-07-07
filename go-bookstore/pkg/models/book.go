package models

import (
	"github.com/jinzhu/gorm"
	"github.com/surbhi/go-bookstore/pkg/config"
)

// the variable name db often represents a database connection or session object, which allows you to interact with a database using the 
// GORM library. The asterisk (*) preceding gorm.DB indicates that db is a pointer to an instance of gorm.DB,With db, you can perform various database 
// operations such as querying, inserting, updating, and deleting records in a database. 
var db *gorm.DB

//A struct type Book is defined with four fields. It embeds gorm.Model, which is a struct provided by GORM and contains common fields like ID, 
// CreatedAt, UpdatedAt, and DeletedAt. The Name, Author, and Publication fields are tagged with json tags, indicating how they should be marshaled 
//and unmarshaled when working with JSON.
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

//The init() function is called automatically when the models package is initialized. It establishes a connection to the database by calling 
//config.Connect(). Then, it assigns the database connection object returned by config.GetDB() to the db variable. Finally, db.AutoMigrate(&Book{}) 
//is used to automatically create the corresponding table in the database based on the Book struct.
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//This method, defined on the Book struct, is used to create a new book record in the database. It calls db.NewRecord(b) to indicate that b is a new 
// record, and then db.Create(&b) is used to insert the b object into the database. Finally, it returns the created Book object.
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

//This function retrieves all book records from the database. It declares a variable Books as a slice of Book objects. Then, it uses db.Find(&Books) 
// to query all the records from the database and populate the Books slice with the results. Finally, it returns the Books slice.
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

//This function retrieves a book record from the database based on the provided Id. It declares a variable getBook of type Book. It uses 
//db.Where("ID = ?", Id).Find(&getBook) to query the record from the database and store the result in getBook. It also returns the pointer to the getBook object and the db object, which represents the query result.
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

//This function deletes a book record from the database based on the provided ID. It declares a variable book of type Book. It uses 
//db.Where("ID = ?", ID).Delete(book) to delete the record from the database. Finally, it returns the book object.Delete expects a pointer to the object to be deleted. 
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}
