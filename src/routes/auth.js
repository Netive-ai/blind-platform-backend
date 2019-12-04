/*

Mocking for user authentificatioon

*/

const { valid } = require('../middleware');
const Joi = require('joi');
const router = require('express').Router();

module.exports = () => {

    const registerSchema = Joi.object().keys({
        mail: Joi.string().min(3).max(20).trim().required(),
        password: Joi.string().min(3).max(20).required()
    });

    router.post('/user/register', valid(registerSchema), async (req, res) => {
        return res.status(httpStatus.OK).send('Success');
    });

    const loginSchema = Joi.object().keys({
        mail: Joi.string().min(3).max(20).trim().required(),
        password: Joi.string().min(3).max(20).trim().required()
    });

    router.post('/user/login', valid(loginSchema), async (req, res) => {
        return res.status(httpStatus.OK).send({
            'token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c'     
        })
    });

    router.get('/user/me', async (req, res) => {
        return res.status(httpStatus.OK).send({
            jspquoi: "des infos sur le user quoi"
        });
    });

    return router;
};