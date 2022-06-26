package models

import(
	"github.com/jinshu/gorm"
	"github.com/jocelyn-yuchiyuan/Booking-management/pkg/config"
)

var db *gorm.db

type Book struct{
	gorm.model
	Name string `gorm:""json:"name"`
	Author string`json:"author"`
	Publication string`json:"publication"`
}

func init(){
	config.Connect()
	db = cofig.GetDB()
	db.AutoMigrate(&Book{})
}
func(b *Book)CreateBook()*Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBooks()[]Book{
	var Books []Bookdb.Find(&Books)
	return Books
}
func GetBookById(Id int64)(*book, *gorm.DB){
	var getBook Book
	db :=db.where("ID = ?", Id).Find(&getBook)
	return &getBook.db
}

func DeleteBook(ID int64){
	var book Book
	db.where("ID= ?", ID).Delete(book)
	return book
}

