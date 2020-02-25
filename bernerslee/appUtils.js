const faker = require('faker');

const Entities = require('./appEntities');
const Patient = Entities.Patient;
const Disease = Entities.Disease;
const Vaccine = Entities.Vaccine;

const Constants = require('./appConstants');
const doctors = Constants.doctors;
const diseases = Constants.diseases;
const vaccines = Constants.vaccines;

function rand(low, high) {
    return Math.round(Math.random() * (high - low) + low);
}

function randFromList(list) {
    const index = rand(0, list.length);
    return list[index];
}

function generateBackground() {
    const count = rand(0, 1);

    let background = [];
    for (let i = 0; i < count; i++) {
        const date = faker.date.past();
        const name = randFromList(diseases);
        const doctor = randFromList(doctors);
        const details = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris in luctus ex, a tincidunt urna. ";
        const disease = new Disease(name, date, doctor, details);

        background.push(disease);
    }

    return background;
}

function generateVaccines() {
    const count = rand(0, 3);

    let result = [];
    for (let i = 0; i < count; i++) {
        const date = faker.date.recent();
        const name = randFromList(vaccines);

        const vaccine = new Vaccine(name, date)
        result.push(vaccine);
    }

    return result;
}

function generatePatient() {
    const firstname = faker.name.firstName();
    const lastname = faker.name.lastName();
    const birthday = faker.date.past();
    const birthplace = faker.address.city();

    const background = generateBackground();
    const vaccines = generateVaccines();

    //We currently do not generate any examination.
    const patient = new Patient(firstname, lastname, birthday, birthplace, background, vaccines, []);
    return patient;
}

module.exports = {
    generatePatient
};