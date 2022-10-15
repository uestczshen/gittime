export tag=v1.0
build:
	echo "building httpserver binary"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./httpserver

release: build
	echo "building httpserver container"
	docker build -t uestczshen/httpserver:${tag} .

push: release
	echo "pushing zz/module3/httpserver"
	docker push uestczshen/httpserver:${tag}
