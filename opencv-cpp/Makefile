CXX = g++
CXXFLAGS = -I/usr/local/include/opencv4/ -std=c++17
LDFLAGS = -L/usr/local/lib
OPENCV_LIBS = -lopencv_core -lopencv_imgcodecs -lopencv_highgui -lopencv_imgproc

SRC = ./src/main.cpp
TARGET = ./bin/main

build:
	$(CXX) $(CXXFLAGS) $(SRC) -o $(TARGET) $(LDFLAGS) $(OPENCV_LIBS)

start: build
	./bin/main
