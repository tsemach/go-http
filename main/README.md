1. create a module by: go mod init github.com/tsemach/go-http-mai
2. go mod edit -replace github.com/tsemach/go-http=../http
3. go mod tidy

### Call the service
curl --cert certs/client.crt --key certs/client.key --cacert certs/ca.crt https://localhost:8080/health