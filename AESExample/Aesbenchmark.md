Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -coverprofile=/tmp/vscode-goFmbGOA/go-code-cover -bench . test.com/model/Desktop/test/AESExample

goos: linux
goarch: amd64
pkg: test.com/model/Desktop/test/AESExample
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkEncrypt-8   	1000000000	         0.2872 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecrypt-8   	1000000000	         0.05527 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 58.3% of statements
ok  	test.com/model/Desktop/test/AESExample	5.080s
