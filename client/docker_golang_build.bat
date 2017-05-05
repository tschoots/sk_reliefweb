SET tag=1.7.4

docker run --rm -v %cd%:/usr/src/myapp -w /usr/src/myapp -e GOOS=windows -e GOARCH=amd64 golang:%tag% bash -c ./build_win.sh