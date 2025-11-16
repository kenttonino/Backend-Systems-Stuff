#include <iostream>
#include <opencv4/opencv2/core.hpp>
#include <opencv4/opencv2/highgui.hpp>
#include <opencv4/opencv2/imgcodecs.hpp>

using namespace std;

/*
  This example will load an image.
*/
int ex_0001_load_image(void) {
  string image_path = cv::samples::findFile("../files/digital_circuit.png");
  cv::Mat img = imread(image_path, cv::IMREAD_COLOR);

  if (img.empty()) {
    std::cout << "Could not read the image: " << image_path << std::endl;
    return 1;
  }

  imshow("Display window", img);
  int k = cv::waitKey(0); // Wait for a keystroke in the window

  if (k == 's') {
    cv::imwrite("starry_night.png", img);
  }

  return 0;
}
