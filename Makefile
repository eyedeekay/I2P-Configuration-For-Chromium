
VERSION=0.0.01

all:
	GOOS=windows go build -o i2pchromium.exe
	GOOS=darwin go build -o i2pchromium-darwin
	GOOS=linux go build -o i2pchromium

release:
	gothub release -p -u eyedeekay -r "I2P-Configuration-for-Chromium" -t $(VERSION) -n "Launchers" -d "A self-configuring launcher for I2P Browsing with Chromium"; true
	gothub upload -R -u eyedeekay -r "I2P-Configuration-for-Chromium" -t $(VERSION) -n "i2pchromium.exe" -f "i2pchromium.exe"
	gothub upload -R -u eyedeekay -r "I2P-Configuration-for-Chromium" -t $(VERSION) -n "i2pchromium-darwin" -f "i2pchromium-darwin"
	gothub upload -R -u eyedeekay -r "I2P-Configuration-for-Chromium" -t $(VERSION) -n "i2pchromium" -f "i2pchromium"