
all:
	GOOS=windows go build -o i2pchromium.exe
	GOOS=darwin go build -o i2pchromium-darwin
	GOOS=linux go build -o i2pchromium