package handler

import (
	"crud-cleancode/usecase"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecaseContract
	name           string
	handler        Handler
}
type ProductHandlerContract interface {
	// Create new product
	// Create(w http.ResponseWriter, r *http.Request)
	// // Update product
	// Update(w http.ResponseWriter, r *http.Request)
	// List of product
	GetProduct(w *gin.Context)
	// Detail of product
	// Detail(w http.ResponseWriter, r *http.Request)
	// // Delete product
	// Delete(w http.ResponseWriter, r *http.Request)
}

func NewProductHandler(db *gorm.DB) ProductHandlerContract {
	return &ProductHandler{
		productUsecase: usecase.NewProductUsecase(db),
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

// func (controller *ProductHandler) DeleteProduct(id string) {
// 	controller.Interactor.Delete(id)
// }
