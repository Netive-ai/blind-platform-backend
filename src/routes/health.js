/*

    Routes to check API status

*/

const { valid } = require('../middleware');
const httpStatus = require('http-status-codes');
const Joi = require('joi');
const router = require('express').Router();

module.exports = () => {

    router.get('/ping', async (req, res) => {
        return res.status(httpStatus.OK).send('pong');
    });

    return router;
};