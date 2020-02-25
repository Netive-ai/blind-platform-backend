const fs = require('fs');
const axios = require('axios');
const FormData = require('form-data');

async function doExamination(attachment_path, type) {
    const form = new FormData();
    form.append('attachment', fs.createReadStream(attachment_path));

    const response = await axios({
        method: 'post',
        url: 'http://localhost:5001/' + type,
        data: form,
        headers: { 'content-type': `multipart/form-data; boundary=${form._boundary}` }
    });

    return response;
}


module.exports = {
    doExamination
}