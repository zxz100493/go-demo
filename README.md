# go-demo-test

go run ./server/server.go -l="127.0.0.1:5000" -k="1234567890abcdef"

go run ./forward.go -l="127.0.0.1:4000" -t="127.0.0.1:5000" -k1="abcdefghigklmnop" -k2="1234567890abcdef"

go run ./client/client.go -t="127.0.0.1:4000" -c="hello word" -k="abcdefghigklmnop"
