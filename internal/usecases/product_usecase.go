package usecases

import (
	"fmt"
	"log"

	"github.com/noffrihendri/golang-crud-gin.git/internal/domain"
	"gorm.io/gorm"

	infrastructure "github.com/noffrihendri/golang-crud-gin.git/internal/infrastructures"
)

type ProductUsecase struct {
	productRepo infrastructure.ProductRepoContract
	name        string
}

type ProductUsecaseContract interface {
	Read() ([]*domain.Product, error)
	GetProductById(ID string) (*domain.Product, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
}

func NewProductUsecase(db *gorm.DB) ProductUsecaseContract {
	return &ProductUsecase{
		productRepo: infrastructure.NewProductRepository(db),
		name:        "Product Usecase",
	}
}

func (p *ProductUsecase) CreateProduct(product *domain.Product) (*domain.Product, error) {
	log.Printf("[%s][Create] is executed\n", p.name)
	if product.Price > 10000 {
		return product, fmt.Errorf("Nilai price kemahalan")
	}

	if product1, err := p.productRepo.CreateProduct(product); err != nil {
		log.Printf("Error : [%s][Create] %s \n", p.name, err.Error())
		return product1, err
	}

	return product, nil
}

func (p *ProductUsecase) Read() ([]*domain.Product, error) {
	log.Printf("[%s][Read] is executed\n", p.name)

	products, _, err := p.productRepo.ListProduct()
	if err != nil {
		log.Printf("Error : [%s][Read] %s \n", p.name, err.Error())
		return products, err
	}

	return products, nil
}

func (p *ProductUsecase) GetProductById(ID string) (*domain.Product, error) {
	log.Printf("[%s][Detail] is executed\n", p.name)

	product, err := p.productRepo.GetProduct(ID)
	if err != nil {
		log.Printf("Error : [%s][Detail] %s \n", p.name, err.Error())
		return product, err
	}

	return product, nil
}
