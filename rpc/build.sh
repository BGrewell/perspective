#!/usr/bin/env bash

if [ ! -d grpc/go ]; then
  echo "[!] Creating go output directory"
  mkdir grpc/go;
fi

if [ ! -d grpc/python ]; then
  echo "[!] Creating python output directory"
  mkdir grpc/python;
fi

echo "[+] Building docker container"
docker image build --no-cache -t perspectivegrpc:1.0 .
docker container run --detach --name grpc perspectivegrpc:1.0
docker cp grpc:/go/src/github.com/BGrewell/perspective/rpc/grpc/go grpc/go/.
echo "[+] Updating of go library complete"

docker cp grpc:/go/src/github.com/BGrewell/perspective/rpc/grpc/python/perspective_pb2.py grpc/python/.
docker cp grpc:/go/src/github.com/BGrewell/perspective/rpc/grpc/python/perspective_pb2_grpc.py grpc/python/.
echo "[+] Updating of python library complete"

echo "[+] Removing docker container"
docker rm grpc

#echo "[+] Adding new files to source control"
git add rpc/grpc/go/perspective.pb.go
git add rpc/grpc/python/perspective_pb2.py
git add rpc/grpc/python/perspective_pb2_grpc.py
git commit -m "regenerated grpc libraries"
git push

echo "[+] Done. Everything has been rebuilt and the repository has been updated and pushed"