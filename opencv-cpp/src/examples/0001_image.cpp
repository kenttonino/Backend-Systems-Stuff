#include <filesystem>
#include <iostream>
#include <opencv4/opencv2/core.hpp>
#include <opencv4/opencv2/highgui.hpp>
#include <opencv4/opencv2/imgcodecs.hpp>

using namespace std;
using namespace cv;
using namespace std::filesystem;

/*
  This example will load an image.
*/
int ex_0001_load_image(void) {
  string absolute_path = current_path().string();
  string image_path = "/src/files/digital_circuit.png";
  Mat img = imread(absolute_path + image_path, IMREAD_COLOR);

  if (img.empty()) {
    cout << "Could not read the image: " << image_path << endl;
    return 1;
  }

  imshow("Display window", img);
  int k = waitKey(0); // Wait for a keystroke in the window

  if (k == 's') {
    imwrite("../files/digital_circuit_saved.png", img);
  }

  return 0;
}
