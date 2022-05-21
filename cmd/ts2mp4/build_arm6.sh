export GOOS=linux
export GOARM=6
export GOARCH=arm
go build -o ../../build/ts2mp4_${GOOS}_${GOARCH} -trimpath -ldflags="-s -w"