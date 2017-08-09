all:
	export GOPATH=~/go
	mkdir -p ~/go/bin
	export GOBIN=$GOPATH/bin
	go get .
	echo All done.
