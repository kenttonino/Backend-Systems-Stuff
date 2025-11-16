# Compiler used.
CXX=g++

# Compiler flags.
CXXFLAGS = -I/usr/local/include/opencv4/ -std=c++11

# OpenCV include and library paths.
OPENCV_INCLUDE = $(shell pkg-config opencv4 --cflags -DOPENCV_GENERATE_PKGCONFIG=ON)
OPENCV_LIBS = $(shell pkg-config opencv4 --libs)

# Source file.
SRC = ./src/main.cpp

# Executable file.
TARGET = ./bin/main

build:
	$(CXX) $(CXXFLAGS) -o $(TARGET)  $(SRC)
