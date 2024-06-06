.PHONY: build templ css run air

# The run target builds the application and then runs it
run: build
	@./bin/app

# The build target compiles the Go application
build:
	@go build -tags '!dev' -o bin/app .

# The templ target runs the templ generate command
templ:
	@templ generate --watch --proxy=http://localhost:1323

# The css target runs the npx tailwindcss command
css:
	npx tailwindcss -I views/css/global.css -o public/global.css --watch

# The air target runs air, templ, and npx tailwindcss together
air: clean
	@echo "Starting air, templ, and css..."
	templ generate --watch --proxy=http://localhost:1323 & \
	npx tailwindcss -i views/css/global.css -o public/global.css --watch & \
	air

# The clean target removes the tmp directory used by air
clean:
	@rm -rf ./tmpmak