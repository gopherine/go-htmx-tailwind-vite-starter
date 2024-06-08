.PHONY: build templ css run air

# The run target builds the application and then runs it
run: build
	@./bin/main

# The build target compiles the Go application
build:
	@go build -tags '!dev' -o bin/main .

# The templ target runs the templ generate command
templ:
	@templ generate --watch --proxy=http://localhost:1323

# The css target runs the npx tailwindcss command

# The air target runs air, templ, and npx tailwindcss together
air: css
	@echo "Starting air, templ, and css..."
	npx tailwindcss -i views/css/global.css -o public/assets/global.css  & \
	templ generate --watch --proxy=http://localhost:1323 & \
	air

css: vite
	npx tailwindcss -i views/css/global.css -o public/assets/global.css

vite: clean
	npx vite build

# The clean target removes the tmp directory used by air
clean:
	@rm -rf ./tmpmak