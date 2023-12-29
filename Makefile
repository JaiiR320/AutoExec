BINARY_NAME=autoexec.exe
STARTUP_PATH="C:\Users\MSI\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup"

all: copy run

build: 
	go build -o ${BINARY_NAME} .

run: build
	./${BINARY_NAME} help

copy: build
	cp ${BINARY_NAME} ${STARTUP_PATH}

clean:
	go clean
	rm ${BINARY_NAME}