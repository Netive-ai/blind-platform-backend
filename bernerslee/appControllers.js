const entities = require('./appEntities');
const database = require('./appDatabase');
const services = require('./appServices');

async function getPatients(_, res) {
    const patients = await database.engine.collection('patients').find().toArray();
    return res.send(patients);
}

async function getPatient(_, res) {
    return res.send(res.locals.patient);
}

async function addExamination(req, res) {
    const name = req.body.name;
    const date = Date.now();
    const doctor = req.body.doctor;
    const attachment = req.file.filename;
    const details = "";
    const diagnosis = await services.doExamination(req.file.path, name);

    const examination = new entities.Examination(name, doctor, date, attachment, diagnosis, details);

    res.locals.patient.examinations.push(examination);
    await database.engine.collection('patients').updateOne({ _id: res.locals.patient._id }, {
        $set:
        {
            examinations: res.locals.patient.examinations
        }
    });

    const entity = await database.engine.collection('patients').findOne({ _id: res.locals.patient._id });
    return res.send(entity);
}

module.exports = {
    getPatient,
    getPatients,
    addExamination,
}