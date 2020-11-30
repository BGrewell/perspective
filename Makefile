

build:
	go build -o bin/sensor sensor/main.go
	go build -o bin/collector collector/main.go

deploy:
	scp sensor/geoip/geolite2.mmdb root@a.hosts.k3rn3l.io:/root/.
	scp sensor/geoip/geolite2.mmdb root@b.hosts.k3rn3l.io:/root/.
	scp bin/sensor root@a.hosts.k3rn3l.io:/root/.
	scp bin/sensor root@b.hosts.k3rn3l.io:/root/.