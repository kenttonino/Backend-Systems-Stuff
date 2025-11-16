## Description

> - Sample OpenCV computer programs in C++.

<br />
<br />
<br />



## Local Setup

> - Check the official [installation](https://docs.opencv.org/4.x/d7/d9f/tutorial_linux_install.html) (Debian 12).

```sh
# Install minimal prerequisites (Ubuntu 18.04 as reference)
sudo apt update && sudo apt install -y cmake g++ wget unzip

# Download and unpack sources
wget -O opencv.zip https://github.com/opencv/opencv/archive/4.x.zip
unzip opencv.zip

# Create build directory
mkdir -p build && cd build

# Configure
cmake  ../opencv-4.x

# Build
cmake --build .
```

> - These steps will move the executables and headers in there desired location.

```sh
sudo make install
```

> - Follow the steps below to run the program.

```sh
# Build the program.
make build

# Run the program with GNU Debugger.
make debug

# Run the program with GNU C++ Compiler.
make start
```
