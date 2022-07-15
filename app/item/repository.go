package item

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Item, error)
	FindByID(ID int) (Item, error)
	Create(item Item) (Item, error)
	Update(item Item) (Item, error)
	Delete(item Item) (Item, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Item, error) {
	var items []Item

	err := r.db.Find(&items).Error

	return items, err
}

func (r *repository) FindByID(ID int) (Item, error) {
	var item Item

	err := r.db.Find(&item, ID).Error

	return item, err
}

func (r *repository) Create(item Item) (Item, error) {
	err := r.db.Create(&item).Error

	return item, err
}

func (r *repository) Update(item Item) (Item, error) {
	err := r.db.Save(&item).Error

	return item, err
}
func (r *repository) Delete(item Item) (Item, error) {
	err := r.db.Delete(&item).Error

	return item, err
}
