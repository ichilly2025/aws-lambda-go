# build
GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go

# package
zip myFunction.zip bootstrap
