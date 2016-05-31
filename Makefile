GOPATH:=$(CURDIR)
export GOPATH

all: build

fmt:
	gofmt -l -w -s src/

dep:fmt
	#go get github.com/aosen/goutils
	#go get github.com/aosen/search
	#go get github.com/astaxie/beego/orm
	#go get github.com/go-sql-driver/mysql
	#go get github.com/gorilla/mux 
	#go get github.com/jakecoffman/cron

build:dep
	go build -o bin/novel src/novel/main.go
	mkdir static
clean:
	#rm -rfv pkg
	rm -rf bin/novel
