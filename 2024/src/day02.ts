import * as util from "./util.ts";

const lines = util.getLinesFromFile('data/day02.txt');

function isSafe(report:number[]) {
    let lastVal:number;
    let increasing:boolean = false;
    for (let i = 0; i < report.length; i++) {
        let val = report[i];
        if (i > 0) {
            lastVal = report[i-1];
            let diff = val - lastVal;
            if (diff == 0) return false;
            if (i == 1) {
                increasing = (diff > 0);
            }
            if (increasing) {
                if (diff < 1 || diff > 3) {
                    return false;
                }
            } else {
                if (diff > -1 || diff < -3) {
                    return false;
                }
            }
        }
    }
    return true;
}

function removeLevel(report:number[], level:number) {
    let newReport:number[] = [];
    for (let i = 0; i < report.length; i++) {
        if (i != level) {
            newReport.push(report[i]);
        }
    }
    return newReport;
}

// PART 1
let numSafeReports = 0
lines.forEach((line) => {
    const levels = util.getNumberArrayFromString(line);
    if (isSafe(levels)) {
        numSafeReports++;
    }
});
console.log('numSafeReports:', numSafeReports);

// PART 2
let numSafeReportsWithProblemDampener = 0
lines.forEach((line) => {
    const report = util.getNumberArrayFromString(line);
    for (let i = -1; i < report.length; i++) {
        const newReport = removeLevel(report, i);
        if (isSafe(newReport)) {
            numSafeReportsWithProblemDampener++;
            return;
        }
    }
});
console.log('numSafeReportsWithProblemDampener:', numSafeReportsWithProblemDampener);
