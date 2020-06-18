package main
import "github.com/dieulinh/go-gorm/rest"
import "log"
func main(){
	log.Fatal(rest.RunAPI(":8081"))
}
