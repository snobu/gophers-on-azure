build:
	@echo -------------------------------------
	@echo Compiling app to ELF static binary...
	@echo -------------------------------------
	CGO_ENABLED=0 GOOS=linux go build -o app-static -a -ldflags '-extldflags "-static"' .
	file app-static > version

deps:
	go get github.com/gorilla/mux
	go get github.com/sirupsen/logrus
	@echo Go deps installed.
