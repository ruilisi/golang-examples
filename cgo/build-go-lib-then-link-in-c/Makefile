all: main.go main.c
	go build -buildmode=c-archive -o libsum.a main.go
	gcc -o main main.c -L. -lsum
