const _module = {
	engine: null,
	seed: seed,
	connect: connect
}

const mongo = require('mongodb').MongoClient;
const url = 'mongodb://root:1234@db:27017';

const utils = require('./appUtils');

const sleep = (milliseconds) => {
	return new Promise(resolve => setTimeout(resolve, milliseconds))
}

async function seed(amount=15) {
	const exists = _module.engine.collection('patients').countDocuments() > 0;
	if (exists)
		return;

	const patients = await _module.engine.createCollection('patients');
	while (amount > 0)
	{
		patient = utils.generatePatient();
		await patients.insertOne(patient);

		amount -= 1;
	}
}

async function connect(amount = 5) {
	if (amount == 0)
		throw new Error("Cannot reach the database.");

	console.log("Connecting to database...");

	try {
		const options = { useNewUrlParser: true }
		const connection = await mongo.connect(url, options);
		_module.engine = connection.db('medical-platform');
	}
	catch {
		await sleep(1000);
		return connect(amount - 1)
	}
	console.log("Connected to database.");

}

module.exports = _module