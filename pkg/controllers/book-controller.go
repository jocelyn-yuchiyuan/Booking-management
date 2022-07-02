package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/jocelyn-yuchiyuan/Booking-management/pkg/utils"
	"github.com/jocelyn-yuchiyuan/Booking-management/pkg/models"

)

var NewBook models.Book
func GetBook(w http.Responsewriter, r *http.Request){
	newbooks := models.GetAllBooks()
	res, _:=json.Marshall(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.statusOK)
	w.Write(res)
}
func GetBookById(w http.Responsewriter, r *http.Request){
	vars :=mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:=strconv.ParsInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _:= models.GetBookById(ID)
	res, _ :=json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Writeheader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.Responsewriter, r *http.Request){
	CreateBook := &models.Book{}
	Utils.ParseBody(r, CreateBook)
	b := CreateBook.Createbook()
	res, _:=json.Marshal(b)
	w.Writeheader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.Responsewriter, r *http.Request){
	vars := mux.vars(r)
	bookId := vars["bookId"]
	ID,err :=strconv.ParsInt(bookId,0 ,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

