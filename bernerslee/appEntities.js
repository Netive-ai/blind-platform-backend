//firstname: string
//lastname: string
//birthdate: string
//birthplace: string
//background: Disease array
//vaccines: vaccine array
//examinations: examination array
class Patient {
    constructor(firstname, lastname, birthdate, birthplace, background, vaccines, examinations)
    {
        this.firstname = firstname;
        this.lastname = lastname;
        this.birthdate = birthdate;
        this.birthplace = birthplace;
        this.background = background;
        this.vaccines = vaccines;
        this.examinations = examinations;
    }
};

//name: string
//date: string
//doctor: string
//details: string
class Disease {
    constructor (name, date, doctor, details)
    {
        this.name = name;
        this.date = date;
        this.doctor = doctor;
        this.details = details;
    }
}

//name: string
//date: string
class Vaccine {
    constructor (name, date)
    {
        this.name = name;
        this.date = date;
    }
}

//name: string
//doctor: string
//date: string
//attachment: string (uri)
//diagnostic: string
//details: string
class Examination {
    constructor (type, doctor, date, attachment, diagnostic, details)
    {
        this.type = type;
        this.doctor = doctor;
        this.date = date;
        this.attachment = attachment;
        this.diagnostic = diagnostic;
        this.details = details;
    }
}

module.exports = {
    Patient,
    Disease,
    Vaccine,
    Examination
};