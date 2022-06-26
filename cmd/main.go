package main
import{
	"log"
	"net/http"
	"github.com/gorilla/mux"
	  "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jocelyn-yuchiyuan/Booking-management/pkg/routes"
}
func main(){
	r := mux.NewRouter()
	router.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010"))

}