# Install the go dependencies
install:
	go install

# Remove previous build
clean:
	if [ -f neondb ]; then rm neondb; fi

# Build the binary
build:
	go build -o neondb .

# Run all the commands in order
all: install clean build copy
