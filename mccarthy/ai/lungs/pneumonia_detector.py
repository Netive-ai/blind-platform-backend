import tensorflow as tf
import keras
import numpy as np
import os
import cv2

IMG_SIZE = 128
MODEL_PATH = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'model.h5')

def analyse(path):
    inp = []
    img = cv2.imread(path, cv2.IMREAD_COLOR)
    img = cv2.resize(img, (IMG_SIZE, IMG_SIZE))
    img.shape

    inp.append(img)
    inp = np.array(inp) / 255
    inp.reshape(-1, IMG_SIZE, IMG_SIZE, 3)

    model = tf.keras.models.load_model(MODEL_PATH)
    prediction = model.predict(inp)[0][0]
    return prediction

