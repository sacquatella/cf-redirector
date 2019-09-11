.EXPORT_ALL_VARIABLES:


build :
	go build ${LDFLAGS} -o cf-redirector-local

buildcf:
	GOOS=linux GOARCH=amd64 go build -o cf-redirector