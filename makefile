dll:
	go build -ldflags="-w -s" -buildmode=c-shared -o nhs-number-validation.dll main.go
