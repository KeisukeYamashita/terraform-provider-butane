version := 99.99.0
provider_macos_path = registry.terraform.io/KeisukeYamashita/butane/$(version)/darwin_arm64/

.PHONY: build
build: 
	@go build

.PHONY: doc
doc:
	@go generate ./...

.PHONY: install_macos
install_macos: build
	@mkdir -p ~/Library/Application\ Support/io.terraform/plugins/$(provider_macos_path)
	@mv terraform-provider-butane ~/Library/Application\ Support/io.terraform/plugins/$(provider_macos_path)
