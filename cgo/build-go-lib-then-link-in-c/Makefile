all: main.go main.c
	go build -buildmode=c-archive -o libsum.a main.go
	gcc main.c -o main -I./ -L./ -lsum -lpthread

clean:
	rm main
	rm *.a *.h
