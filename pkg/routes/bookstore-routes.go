package routes
import(
	"github.com/gorilla/mux"
	"github.com/jocelyn-yuchiyuan/Booking-management/pkg/controllers"
)
var RegisterBookStore = func(router *mux.Router){
	router.HandlerFuc("/book/", controllers.Createbook).Methods("POST")
	router.HandlerFuc("/book/", controllers.Getbook).Methods("GET")
	router.HandlerFuc("/book/{bookID}", controllers.GetBookById).Methods("GET")
	router.HandlerFuc("/book/{bookID}", controllers.UpdateBookById).Methods("PUT")
	router.HandlerFuc("/book/{bookID}", controllers.DeleteBookById).Methods("DELETE")
}