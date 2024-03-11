GOCMD=go
GOBUILD=go build -ldflags="-s -w"
GO_FILE= ${PWD}/cmd/main.go
BINARY_FILE=kline

all: remove-binary install building-package

install:
	echo "Generating binary in the pkg folder"
	${GOBUILD} -o ${BINARY_FILE} ${GO_FILE}
	sudo cp ${BINARY_FILE} /usr/local/bin
	mv ${BINARY_FILE} ./kLine_amd64/usr/local/bin/

building-package:
	echo "removing existing package..."
	rm *.deb
	echo "building debian package..."
	dpkg-deb --build kLine_amd64

remove-binary: 
	echo "Removing existing binary..."
	sudo rm -f ${BINARY_FILE}
	sudo rm -f ./kLine_amd64/usr/local/bin/kline
	sudo rm -f /usr/local/bin/${BINARY_FILE}

.PHONY: all remove-binary install building-package
