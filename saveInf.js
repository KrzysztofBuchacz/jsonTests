var fs = require('fs');

var maxNumber = Math.pow(10, 1000);
if (maxNumber === Infinity) {
  console.log("Infinity!");
}

var obj = { volume: 0.0 }
obj.volume = maxNumber;

var json = JSON.stringify(obj);
fs.writeFileSync('myjsonfile.json', json);