VERSION="1.0"
OUTFILE=conv

default: build


build: 
	go build -o ${OUTFILE}	

install:
	go install ${LDFLAGS}

cross:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc-win32 go build -o ${OUTFILE}.exe
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc go build -o ${OUTFILE}-386.exe	

run: 
	go run main.go

zip: build	
	zip ${OUTFILE}-linux-$(shell arch)-${VERSION}.zip ${OUTFILE}
	rm -f ${OUTFILE}

zipcross: cross
	zip ${OUTFILE}-win64-${VERSION}.zip ${OUTFILE}.exe
	zip ${OUTFILE}-i386-${VERSION}.zip ${OUTFILE}-386.exe
	rm -f ${OUTFILE}*.exe


clean:
	rm -f  ${OUTFILE}
	rm -f  ${OUTFILE}*.exe
	rm -f  ${OUTFILE}*.zip 