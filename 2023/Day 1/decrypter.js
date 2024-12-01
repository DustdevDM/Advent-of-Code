import fs from "fs";

var file = fs.readFileSync("./input").toString();

var replaceValues = [
  { search: /one/g, replace: "o1e" },
  { search: /two/g, replace: "t2o" },
  { search: /three/g, replace: "t3e" },
  { search: /four/g, replace: "f4r" },
  { search: /five/g, replace: "f5e" },
  { search: /six/g, replace: "s6x" },
  { search: /seven/g, replace: "s7n" },
  { search: /eight/g, replace: "e8t" },
  { search: /nine/g, replace: "n9e" },
];

var lineArray = file.split("\n");
var extracetValues = [];

lineArray.forEach((line, index) => {
  var lineReplacement = line;
  replaceValues.forEach(rv => {
    lineReplacement = lineReplacement.replace(rv.search, rv.replace);
  });
  lineArray[index] = lineReplacement;
});

lineArray.forEach(line => {
  var numbers = line.split("").filter(x => {
    return x.match(/\d+/);
  });

  var number = Number.parseInt(`${numbers[0]}${numbers[numbers.length - 1]}`);
  if (!isNaN(number)) extracetValues.push(number);
});

console.log(extracetValues);

var result = 0;

extracetValues.forEach(y => {
  result += y;
});

console.log("The result is: " + result);
