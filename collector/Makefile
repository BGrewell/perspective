
build:
	[ -d bin ] || mkdir bin
	go build -o bin/collectord main.go

deploy:
	ssh root@attackmap.grewelltech.com -C "mkdir -p /opt/collectord/bin || true"
	ssh root@attackmap.grewelltech.com -C "mkdir -p /etc/collectord/ || true"
	scp bin/collectord root@attackmap.grewelltech.com:/opt/collectord/bin/.
	scp config/config.yaml root@attackmap.grewelltech.com:/etc/collectord/.
	scp service/collectord.service root@attackmap.grewelltech.com:/lib/systemd/system/.
	ssh root@attackmap.grewelltech.com -C "systemctl enable collectord"
	ssh root@attackmap.grewelltech.com -C "systemctl start collectord"