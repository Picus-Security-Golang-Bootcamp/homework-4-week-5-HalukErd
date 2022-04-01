.PHONY: models generate

# ==============================================================================
# Swagger Models
models:
	$(call print-target)
	find ./models/generated -type f -not -name '*_test.go' -delete
	swagger generate model -m api -f ./swagger/book.yml -t ./models/generated

generate: models
