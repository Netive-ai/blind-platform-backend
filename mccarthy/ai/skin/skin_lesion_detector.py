import os

import keras
import numpy as np
from PIL import Image

lesion_type = [
    "Actinic keratoses",
    "Basal cell carcinoma",
    "Benign keratosis-like lesions",
    "Dermatofibroma",
    "Melanocytic nevi",
    "Melanoma",
    "Vascular lesions",
]


model_path = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'model.h5')
model = keras.models.load_model(model_path)
model._make_predict_function()

def analyse(path):
    img = Image.open(path).resize((100, 75))
    img = np.array(img)
    img = (img - img.mean()) / img.std()
    img = img.reshape(1, 75, 100, 3)

    prediction = model.predict(img).argmax()
    print('prediction:' + lesion_type[prediction])

    return lesion_type[prediction]
