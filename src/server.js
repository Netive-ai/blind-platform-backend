/*

Netive platform backend

*/

const express = require('express');
const auth = require('./routes/auth')();
const health = require('./routes/health')();

const app = express();

app.get('/ping', (req, res) => {
    return res.send('pong');
});

app.use(auth);
app.use(health);

app.listen(8080);

console.log(`Listening on ${8080}`);