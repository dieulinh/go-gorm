package rest
import (
	"os"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"net/http"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

  "github.com/dieulinh/go-gorm/models"
)

var err error
type Handler struct {
	db *gorm.DB
}

func (dt *Handler) GetProducts(c *gin.Context) {

  var err error
	products := []models.Product{}

	err = dt.db.Model(&models.Product{}).Limit(100).Find(&products).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": products,
	})
	return
}
func Connect(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *gorm.DB {
  var dt *gorm.DB
  log.Printf("Connected to db")
  log.Printf(Dbdriver)
  log.Printf("host=%v port=%v user=%v dbname=%v sslmode=disable password=%v", DbHost, DbPort, DbUser, DbName, DbPassword)
  if Dbdriver == "postgres" {
    DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
    server, err := gorm.Open(Dbdriver, DBURL)
    if err != nil {
      log.Printf("Cannot connect to %s database", Dbdriver)
      log.Fatal("This is the error:", err)
      dt = server
    } else {
      log.Printf("We are connected to the %s database", Dbdriver)
      server.Debug().AutoMigrate(&models.Product{})
      dt = server
    }
  }

  return dt
}


func RunAPI(address string) error{

	err := godotenv.Load()
	if err != nil {
		return err
	}
	route := gin.Default()
	fmt.Println(os.Getenv("DB_DRIVER"))
	h := Handler{}
	h.db = Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), "", os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	route.GET("/", welcome)
	route.GET("/products", (&h).GetProducts)
	return route.Run(address)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Welcome to gorm API",
	})
	return
}
