const express = require('express');
const controllers = require('./appControllers');
const middlewares = require('./appMiddlewares');

const router = express.Router();
router.get('/patient', controllers.getPatients);
router.use('/patient/:id', middlewares.patientMiddleware);
router.get('/patient/:id', controllers.getPatient)

module.exports = router