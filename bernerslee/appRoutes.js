const express = require('express');
const controllers = require('./appControllers');
const middlewares = require('./appMiddlewares');

const router = express.Router();
router.get('/patient', controllers.getPatients);

router.get('/patient/:id', middlewares.patientMiddleware, controllers.getPatient)
router.post('/patient/:id/examination', middlewares.patientMiddleware,
                                        (req, res, next) => middlewares.uploadMiddleware(req.params.id)(req, res, next),
                                        controllers.addExamination);

module.exports = router