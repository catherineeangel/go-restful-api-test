package sample

type SimpleRepository struct {
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: repository,
	}
}
