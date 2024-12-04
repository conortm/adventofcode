import * as util from "./util.ts";

const lines = util.getLinesFromFile('data/day03.txt');
// const lines = ["xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"]

// PART 1
let result = 0;

lines.forEach((line) => {
    const regex = /mul\(\d{1,3}\,\d{1,3}\)/g;
    const found = line.match(regex);
    found?.forEach((mul) => {
        const digits = mul.slice(4, -1).split(',');
        const product = +digits[0] * +digits[1];
        result += product;
    })
});

console.log('result:', result);

// PART 2
let resultWithConditionals = 0;

let enabled = true;
lines.forEach((line) => {
    const regex = /(mul\(\d{1,3}\,\d{1,3}\)|do\(\)|don\'t\(\))/g;
    const found = line.match(regex);
    found?.forEach((mul) => {
        switch (mul) {
            case 'do()':
                enabled = true;
                break;
            case 'don\'t()':
                enabled = false;
                break;
            default:
                if (enabled) {
                    const digits = mul.slice(4, -1).split(',');
                    const product = +digits[0] * +digits[1];
                    resultWithConditionals += product;
                }
        }
    })
});

console.log('resultWithConditionals:', resultWithConditionals);
