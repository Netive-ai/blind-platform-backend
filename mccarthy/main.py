import os
import uuid

from flask import Flask
from flask import request, jsonify, make_response

from ai.skin import skin_lesion_detector
from ai.lungs import pneumonia_detector
from ai.fall import fall_detector

app = Flask(__name__)

artificial_intelligences = ['skin', 'falls', 'lungs']

def save_attachment(attachment):
    _, ext = os.path.splitext(attachment.filename)

    uploads_path = os.path.join(os.getcwd(), 'uploads')
    filename = "%s%s" % (str(uuid.uuid1()), ext)
    path = os.path.join(uploads_path, filename)

    attachment.save(path)
    return path


def analyse(ai, attachment):
    result = None

    if ai == 'skin':
        result = skin_lesion_detector.analyse(attachment)
    elif ai == 'lungs':
        result = str(pneumonia_detector.analyse(attachment))
    elif ai == 'falls':
        result = fall_detector.analyse(attachment)
    
    return result


@app.route('/ai/<ai>', methods=['POST'])
def func(ai):
    if ai not in artificial_intelligences:
        return make_response('unknown artificial intelligence.', 403)
    if 'attachment' not in request.files:
        return make_response('malformated request.', 403)
    
    f = request.files['attachment']
    attachment = save_attachment(f)
    result = analyse(ai, attachment)

    response = jsonify({'result': result})
    response.headers.add('Access-Control-Allow-Origin', '*')
    return response


if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5001)

