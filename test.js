var fs = require('fs');
var json = JSON.parse(fs.readFileSync('badInf.json', 'utf8'));
console.log("Data:")
console.log(json)