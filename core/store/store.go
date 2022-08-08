package stores


func NewStore() Store {
	return &store{}
}

type Store interface {
	//todo:add needed methods for your database
}
type store struct {
}
