go mod init example.com/hello

go mod edit -replace example.com/greetings=../greetings

go mod tidy