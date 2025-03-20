# aws-lambda-go
An AWS Lambda Function in Golang

# How to use
## 1. Package and upload lambda function to s3
```shell
# package
chmod +x package.sh
./package

# upload
aws s3 cp myFunction.zip s3://{your_bucket_name}
```

## 2. Create a Lambda function
1. Function name: "my-go-server"
2. Runtime: "Amazon Linux 2023"
3. Architecture: "arm64"
4. Upload from: s3://{your_bucket_name}/myFunction.zip

## 3. Create an API in API Gateway
1. API name: "GoServerAPI"
2. Resource path: "/v1/users/{name}"
3. Method: "POST"
4. Lambda function: {your_lambda_arn}

## 4. Test API
```shell
curl \
  -d '{}' \
  -H 'Content-Type:application/json' \
  "{invoke_url}/v1/users/chilly?age=20"

{"message":"Hello, chilly! You are 20 years old."}
```

# Files 
## main.go
1. Reads request body containing name and age
2. Replace name with path parameter "name"
3. Replace age with query parameter "age"

## package.sh
1. Build main.go to a binary file called "bootstrap"
2. Zip bootstrap to myFunction.zip
