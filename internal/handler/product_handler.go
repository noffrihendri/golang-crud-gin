package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noffrihendri/golang-crud-gin.git/internal/usecases"
	"gorm.io/gorm"
)

type ProductHandler struct {
	productUsecase usecases.ProductUsecaseContract
	name           string
	handler        Handler
}
type ProductHandlerContract interface {
	GetProduct(w *gin.Context)
	GetProductByID(w *gin.Context)
}

func NewProductHandler(db *gorm.DB) ProductHandlerContract {
	return &ProductHandler{
		productUsecase: usecases.NewProductUsecase(db),
		name:           "Product Handler",
	}
}

func (p *ProductHandler) GetProduct(c *gin.Context) {
	fmt.Println("masuk sini")

	products, err := p.productUsecase.Read()
	if err != nil {
		log.Println("err ", err)
	}
	c.JSON(http.StatusOK, products)
	//return res
}
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	// id := strconv.Itoa(c.Param("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
	// 	return
	// }

	product, err := h.productUsecase.GetProductById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}
