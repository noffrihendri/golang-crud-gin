package handler

import (
	"crud-cleancode/internal/domain"
	"crud-cleancode/internal/usecases"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	productUsecase usecases.ProductUsecaseContract
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
	GetProductByID(w *gin.Context)
	CreateProduct(w *gin.Context)
	// Detail of product
	// Detail(w http.ResponseWriter, r *http.Request)
	// // Delete product
	// Delete(w http.ResponseWriter, r *http.Request)
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
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := h.productUsecase.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

// func (controller *ProductHandler) DeleteProduct(id string) {
// 	controller.Interactor.Delete(id)
// }
