.PHONY: build templ templ-build css css-watch assets run air dev gen clean tools tools-templ tools-air

GO ?= go
NODE ?= node
TOOLBIN ?= $(shell $(GO) env GOPATH)/bin
TEMPL_BIN ?= $(TOOLBIN)/templ
AIR_BIN ?= $(TOOLBIN)/air
AIR_CONFIG ?= .air.toml
TAILWIND ?= $(NODE) ./node_modules/.bin/tailwindcss -c tailwind.config.cjs
ASSETS_BUILD ?= $(NODE) scripts/esbuild.mjs --minify
ASSETS_DEV ?= $(NODE) scripts/esbuild.mjs --watch

# The run target builds the application and then runs it
run: build
	@./bin/main

# The build target compiles the Go application
build: gen templ-build assets css
	@go build -tags '!dev' -o bin/main .

# The templ target runs the templ generate command
templ:
	@$(MAKE) tools-templ
	@$(TEMPL_BIN) generate --watch --proxy=http://localhost:1323

templ-build:
	@$(MAKE) tools-templ
	@$(TEMPL_BIN) generate

# The css target runs the tailwindcss command
css:
	$(TAILWIND) -i views/css/global.css -o public/assets/global.css

css-watch:
	$(TAILWIND) -i views/css/global.css -o public/assets/global.css --watch

assets:
	$(ASSETS_BUILD)

# The air target runs air, templ, and css together
air: css
	@echo "Starting air, templ, css, and assets..."
	@$(MAKE) tools
	$(ASSETS_DEV) & \
	$(TAILWIND) -i views/css/global.css -o public/assets/global.css --watch & \
	$(TEMPL_BIN) generate --watch --proxy=http://localhost:1323 & \
	$(AIR_BIN) -c $(AIR_CONFIG)

dev:
	@echo "Starting assets, css, templ (if installed), and server (air if installed)..."
	@$(MAKE) tools
	$(ASSETS_DEV) & \
	$(TAILWIND) -i views/css/global.css -o public/assets/global.css --watch & \
	$(TEMPL_BIN) generate --watch --proxy=http://localhost:1323 & \
	$(AIR_BIN) -c $(AIR_CONFIG) & \
	wait

tools:
	@$(MAKE) tools-templ tools-air

tools-templ:
	@test -x $(TEMPL_BIN) || { echo "Installing templ..."; $(GO) install github.com/a-h/templ/cmd/templ@latest; }

tools-air:
	@test -x $(AIR_BIN) || { echo "Installing air..."; $(GO) install github.com/air-verse/air@latest; }

gen:
	go generate ./...

# The clean target removes the tmp directory used by air
clean:
	@rm -rf ./tmp
