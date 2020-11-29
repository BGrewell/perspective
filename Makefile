

build:
	go build -o bin/sensor sensor/main.go

deploy:
	scp bin/sensor root@b.hosts.k3rn3l.io:/root/.