.PHONY: start kill run

run: build
	./run.exe
build:
	go mod tidy
	go build -o run.exe	
start: 
	podman system  prune -f
	podman run -p 9324:9324  -p 9325:9325 -d --rm --name sqs -v `pwd`/config/queue.conf:/opt/elasticmq.conf docker.io/softwaremill/elasticmq

# docker run -v `pwd`/application.ini:/opt/docker/conf/application.ini -v `pwd`/logback.xml:/opt/docker/conf/logback.xml -p 9324:9324 softwaremill/elasticmq
kill:
	podman kill sqs
log:
	podman logs -f sqs