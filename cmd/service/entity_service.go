package service

// Interface genérica para definir os métodos CRUD para qualquer tipo de entidade
type EntityService interface {
	GetAll() ([]interface{}, error)
	GetByID(id int) (interface{}, error)
	Update(id int, data interface{}) error
	Delete(id int) error
	DeleteAll() error
}
