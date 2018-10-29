image:
	@echo -------------------------------------
	@echo Compiling app to ELF static binary...
	@echo -------------------------------------
	GO_ENABLED=0 GOOS=linux go build -o app-static -a -ldflags '-extldflags "-static"' .
	file app-static > version
	@echo ------------------------
	@echo BUILDING DOCKER IMAGE...
	@echo ------------------------
	docker build . -t shipyard.azurecr.io/adcaline/gophers-on-azure

push:
	docker push adcaline/gophers-on-azure

deps:
	go get github.com/gorilla/mux
	go get github.com/sirupsen/logrus
	@echo Go deps installed.
