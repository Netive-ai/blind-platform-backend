const database = require('./appDatabase');
const multer = require('multer');
const uuid = require('uuid/v4');
const path = require('path');
const fs = require('fs');

const ObjectId = require('mongodb').ObjectId;

async function patientMiddleware(req, res, next) {
    const id = new ObjectId(req.params.id);
    const patient = await database.engine.collection('patients').findOne({ _id: id })

    res.locals.patient = patient;

    next();
}

const uploadMiddleware = (patientId) => {
    const storePath = 'public/' + patientId;
    const storageOptions = multer.diskStorage({
        destination: (_, __, next) => {
            if (fs.existsSync(storePath) == false)
                fs.mkdirSync(storePath);
            next(null, storePath);
        },
        filename: (_, file, next) => next(null, uuid() + path.extname(file.originalname))
    });
    
    return multer({ storage: storageOptions }).single('attachment');
}

module.exports = {
    uploadMiddleware,
    patientMiddleware
};