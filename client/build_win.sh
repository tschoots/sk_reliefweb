rm -rf sk_reliefweb.exe;CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o sk_reliefweb.exe