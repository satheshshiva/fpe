generatenative:
	go build -buildmode=c-shared -o bin/lib-fpe.so

test:
	go test -v github.com/capitalone/fpe/ff1

benchmark:
	go test -v -bench=. -run=NONE github.com/capitalone/fpe/ff1
