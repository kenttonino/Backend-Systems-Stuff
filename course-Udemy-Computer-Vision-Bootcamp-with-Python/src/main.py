import cv2
import os

absolute_path = os.path.abspath("src/files/einstein.jpg")
image = cv2.imread(absolute_path)

if image is not None:
    # Resize the image.
    resized_image = cv2.resize(image, (0,0), fx=0.5, fy=0.5)

    # Show the resized image.
    cv2.imshow('Einsten', resized_image)

    # Exit when a key is pressed.
    cv2.waitKey(0)

    # Destroy all windows upon exist.
    cv2.destroyAllWindows()


print(cv2.__version__)
