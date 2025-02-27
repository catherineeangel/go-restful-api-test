mock:
	mockgen -source=controller/category_controller.go -destination=controller/mocks/category_controller_mock.go -package=mocks
	mockgen -source=repository/category_repository.go -destination=repository/mocks/category_repository_mock.go -package=mocks
	mockgen -source=service/category_service.go -destination=service/mocks/category_service_mock.go -package=mocks

	mockgen -source=controller/customer_controller.go -destination=controller/mocks/customer_controller_mock.go -package=mocks
	mockgen -source=repository/customer_repository.go -destination=repository/mocks/customer_repository_mock.go -package=mocks
	mockgen -source=service/customer_service.go -destination=service/mocks/customer_repository_mock.go -package=mocks

