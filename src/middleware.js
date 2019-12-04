const Joi = require('joi');

const valid = (schema, location = 'body') => (req, res, next) => {
    const { error, value } = Joi.validate(req[location], schema.required());
    if (error) {
      return res.status(httpStatus.BAD_REQUEST).send(error);
    }
    req.value = value;
    return next();
};

module.exports = {
    valid
};