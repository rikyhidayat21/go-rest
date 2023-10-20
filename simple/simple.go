package simple

type SimpleRepository struct {
}

// Provide a constructor for the SimpleRepository
func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

type SimpleService struct {
	*SimpleRepository
}

// provide a constructor for the SimpleService
func NewSimpleService(simpleRepository *SimpleRepository) *SimpleService {
	return &SimpleService{SimpleRepository: simpleRepository}
}
