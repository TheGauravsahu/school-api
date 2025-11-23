run: build
      @./bin/school-api

build:
	@go build -o @./bin/school-api cmd/school-api/main.go