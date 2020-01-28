const database = require('./appDatabase');

async function patientMiddleware(req, res, next) {
    const id = req.params.id
    
    res.locals.patient = database.engine.collection('patients').find({_id: id});

    next();
}

module.exports = {
    patientMiddleware
}