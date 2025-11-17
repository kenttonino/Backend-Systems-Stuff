import cv2
import os

absolute_path = os.path.abspath("src/files/einstein.jpg")
image = cv2.imread(absolute_path, cv2.IMREAD_GRAYSCALE)

if image is not None:
    resized_image = cv2.resize(image, (0,0), fx=0.3, fy=0.3)

    cv2.imshow('Einsten', resized_image)
    cv2.waitKey(0)


print(cv2.__version__)
