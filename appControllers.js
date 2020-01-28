const database = require('./appDatabase');

async function getPatients(_, res) {
    patients = await database.engine.collection('patients').find().toArray();
    return res.send(patients);
}
async function getPatient(_, res){ 
    return res.send(res.locals.patient);
}


module.exports = {
    getPatient,
    getPatients,
}