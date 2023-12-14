package usecase

import (
	"crud-cleancode/domain"
	"crud-cleancode/infrastructure"
	"log"

	"gorm.io/gorm"
)

type ProductUsecase struct {
	productRepo infrastructure.ProductRepoContract
	name        string
}

type ProductUsecaseContract interface {
	// Create new product
	//Create(name, description string, stock int) (domain.Product, error)
	// List of product
	Read() ([]domain.Product, error)
	// Detail of product
	// Detail(ID int) (domain.Product, error)
	// // Update existing product
	// Update(ID int, name, description string, stock int) (domain.Product, error)
	// // Delete product
	// Delete(ID int) error
}

func NewProductUsecase(db *gorm.DB) ProductUsecaseContract {
	return &ProductUsecase{
		productRepo: infrastructure.NewProductRepository(db),
		name:        "Product Usecase",
	}
}

// func (p *ProductUsecase) Create(name, description string, stock int) (domain.Product, error) {
// 	log.Printf("[%s][Create] is executed\n", p.name)
// 	// product := domain.Product{
// 	// 	Name:        name,
// 	// 	Description: description,
// 	// 	Stock:       stock,
// 	// }

// 	if err := p.productRepo.Create(&product); err != nil {
// 		log.Printf("Error : [%s][Create] %s \n", p.name, err.Error())
// 		return product, err
// 	}

// 	return product, nil
// }

func (p *ProductUsecase) Read() ([]domain.Product, error) {
	log.Printf("[%s][Read] is executed\n", p.name)

	products, _, err := p.productRepo.ListProduct()
	if err != nil {
		log.Printf("Error : [%s][Read] %s \n", p.name, err.Error())
		return products, err
	}

	return products, nil
}

func (p *ProductUsecase) Detail(ID int) (domain.Product, error) {
	log.Printf("[%s][Detail] is executed\n", p.name)

	product, err := p.productRepo.GetProduct(ID)
	if err != nil {
		log.Printf("Error : [%s][Detail] %s \n", p.name, err.Error())
		return product, err
	}

	return product, nil
}

// func (p *ProductUsecase) Update(ID int, name, description string, stock int) (domain.Product, error) {
// 	log.Printf("[%s][Update] is executed\n", p.name)
// 	product := domain.Product{
// 		ID:          ID,
// 		Name:        name,
// 		Description: description,
// 		Stock:       stock,
// 	}

// 	if err := p.productRepo.Update(&product); err != nil {
// 		log.Printf("Error : [%s][Update] %s \n", p.name, err.Error())
// 		return product, err
// 	}

// 	return product, nil
// }

// func (p *ProductUsecase) Delete(ID int) error {
// 	log.Printf("[%s][Delete] is executed\n", p.name)

// 	err := p.productRepo.Delete(ID)
// 	if err != nil {
// 		log.Printf("Error : [%s][Delete] %s \n", p.name, err.Error())
// 		return err
// 	}

// 	return nil
// }
