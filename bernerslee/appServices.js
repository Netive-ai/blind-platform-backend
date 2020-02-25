const fs = require('fs');
const axios = require('axios');
const FormData = require('form-data');

async function doExamination(attachment_path, type) {
    const form = new FormData();
    form.append('attachment', fs.createReadStream(attachment_path));

    if (type == "1")
        type = "skin"
    else if (type == "2")
        type = "lungs"

    const response = await axios({
        method: 'post',
        url: 'http://mccarthy:5001/ai/' + type,
        data: form,
        headers: { 'content-type': `multipart/form-data; boundary=${form._boundary}` }
    });

    return response.data.result;
}


module.exports = {
    doExamination
}