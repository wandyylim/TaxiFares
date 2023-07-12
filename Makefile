
# Generate mocks target
genmocks:
	mockgen -source=internal\usecase\TaxiFares.go -destination=internal\mocks\usecase\mock_TaxiFares.go -package=usecase
