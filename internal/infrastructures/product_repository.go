package infrastructures

import (
	"crud-cleancode/internal/domain"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db   *gorm.DB
	name string
}

type ProductRepoContract interface {
	// Create a new Product
	CreateProduct(product *domain.Product) (*domain.Product, error)
	// Update Product
	UpdateProduct(product *domain.Product) error
	// List Product
	ListProduct() ([]*domain.Product, int, error)
	// Get Product
	GetProduct(ID string) (*domain.Product, error)
	// Delete Product
	DeleteProduct(ID string) error
}

// Create new role product instance
func NewProductRepository(db *gorm.DB) ProductRepoContract {
	return &ProductRepo{
		db:   db,
		name: "Product Repository",
	}
}

func (r *ProductRepo) CreateProduct(product *domain.Product) (*domain.Product, error) {
	log.Printf("[%s][Create] is executed\n", r.name)

	sql := fmt.Sprintf("insert into product(name,price,quantity) values('%s',%f,%d) RETURNING id,name,price,quantity", product.Name, product.Price, product.Quantity)

	if err := r.db.Raw(sql).Scan(&product).Error; err != nil {
		log.Printf("Error : [%s][Create] %s\n", r.name, err.Error())
		return product, nil
	}
	fmt.Println("model created", product)
	return product, nil
}

func (r *ProductRepo) GetProduct(ID string) (*domain.Product, error) {
	log.Printf("[%s][Get] is executed\n", r.name)

	db := r.db
	var product domain.Product

	if err := db.Debug().Table("product").Where("id = ?", ID).First(&product).Error; err != nil {
		log.Printf("Error : [%s][GET] %s", r.name, err.Error())
		return &product, err
	}

	return &product, nil
}

func (r *ProductRepo) ListProduct() ([]*domain.Product, int, error) {
	log.Printf("[%s][List] is executed\n", r.name)

	var count int64
	var products []*domain.Product

	db := r.db

	db.Table("product").Find(&products)

	return products, int(count), nil
}

func (r *ProductRepo) UpdateProduct(product *domain.Product) error {
	log.Printf("[%s][Update] is executed\n", r.name)

	if err := r.db.Model(&product).Updates(&product).Error; err != nil {
		log.Printf("Error : [%s][Update] %s", r.name, err.Error())
		return err
	}

	return nil
}

func (r *ProductRepo) DeleteProduct(ID string) error {
	log.Printf("[%s][Delete] is executed\n", r.name)

	var product domain.Product

	if err := r.db.Delete(&product, ID).Error; err != nil {
		log.Printf("Error : [%s][Delete] %s", r.name, err.Error())
		return err
	}

	return nil
}
