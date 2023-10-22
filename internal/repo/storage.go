package repo

type UserRepository interface {
	Update()
	Get()
	Delete()
	Create()
	Search()
}
