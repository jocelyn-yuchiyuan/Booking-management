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
	res,   :=json.Marshall(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.statusOK)
	w.Write(res)
}