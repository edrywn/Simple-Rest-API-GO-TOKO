package item

type Service interface {
	FindAll() ([]Item, error)
	FindByID(ID int) (Item, error)
	Create(itemRequest ItemRequest) (Item, error)
	Update(ID int, itemRequest ItemRequest) (Item, error)
	Delete(ID int) (Item, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Item, error) {
	items, err := s.repository.FindAll()
	return items, err
}

func (s *service) FindByID(ID int) (Item, error) {

	item, err := s.repository.FindByID(ID)

	return item, err
}

func (s *service) Create(itemRequest ItemRequest) (Item, error) {

	price, _ := itemRequest.Price.Int64()
	stock, _ := itemRequest.Stock.Int64()

	item := Item{
		Name:  itemRequest.Name,
		Price: int(price),
		Stock: int(stock),
	}

	newItem, err := s.repository.Create(item)

	return newItem, err
}
func (s *service) Update(ID int, itemRequest ItemRequest) (Item, error) {

	item, _ := s.repository.FindByID(ID)

	price, _ := itemRequest.Price.Int64()
	stock, _ := itemRequest.Stock.Int64()

	item.Name = itemRequest.Name
	item.Price = int(price)
	item.Stock = int(stock)

	newItem, err := s.repository.Update(item)

	return newItem, err
}
func (s *service) Delete(ID int) (Item, error) {

	item, _ := s.repository.FindByID(ID)

	newItem, err := s.repository.Delete(item)

	return newItem, err
}
