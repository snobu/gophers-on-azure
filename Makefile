build:
	@echo -------------------------------------
	@echo Compiling app to ELF static binary...
	@echo -------------------------------------
	CGO_ENABLED=0 GOOS=linux go build -o app-static -a -ldflags '-extldflags "-static"' .
	file app-static > version

deps:
	go get .
	@echo Go deps installed.
