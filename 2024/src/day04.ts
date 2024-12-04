import * as util from "./util.ts";

// const data = `MMMSXXMASM
// MSAMXMSMSA
// AMXSXMAAMM
// MSAMASMSMX
// XMASAMXAMM
// XXAMMXXAMA
// SMSMSASXSS
// SAXAMASAAA
// MAMMMXMMMM
// MXMXAXMASX`;
const data = util.getDataFromFile('data/day04.txt');

function getMatrixFromData(data: string): string[][] {
    let matrix: string[][] = [];
    data.split('\n').forEach(line => {
        let row: string[] = [];
        line.split('').forEach((val) => {
            row.push(val);
        });
        matrix.push(row);
    });
    return matrix;
}

const m = getMatrixFromData(data);
const numRows = m.length;
const numCols = m[0].length;

// PART 1
let count = 0;
for (let r = 0; r < numRows; r++) {
    for (let c = 0; c < numCols; c++) {
        const letter = m[r][c];
        if (letter == 'X') {
            // r
            if (c + 3 < numCols) {
                const word = letter + m[r][c+1] + m[r][c+2] + m[r][c+3];
                if (word == 'XMAS') count++;
            }
            // l
            if (c - 3 >= 0) {
                const word = letter + m[r][c-1] + m[r][c-2] + m[r][c-3];
                if (word == 'XMAS') count++;
            }
            // d
            if (r + 3 < numRows) {
                const word = letter + m[r+1][c] + m[r+2][c] + m[r+3][c];
                if (word == 'XMAS') count++;
            }
            // u
            if (r - 3 >= 0) {
                const word = letter + m[r-1][c] + m[r-2][c] + m[r-3][c];
                if (word == 'XMAS') count++;
            }
            // rd
            if (c + 3 < numCols && r + 3 < numRows) {
                const word = letter + m[r+1][c+1] + m[r+2][c+2] + m[r+3][c+3];
                if (word == 'XMAS') count++;
            }
            // ru
            if (c + 3 < numCols && r - 3 >= 0) {
                const word = letter + m[r-1][c+1] + m[r-2][c+2] + m[r-3][c+3];
                if (word == 'XMAS') count++;
            }
            // ld
            if (c - 3 >= 0 && r + 3 < numRows) {
                const word = letter + m[r+1][c-1] + m[r+2][c-2] + m[r+3][c-3];
                if (word == 'XMAS') count++;
            }
            // lu
            if (c - 3 >= 0 && r - 3 >= 0) {
                const word = letter + m[r-1][c-1] + m[r-2][c-2] + m[r-3][c-3];
                if (word == 'XMAS') count++;
            }
        }
    }
}
console.log('XMAS count:', count);

// PART 2
count = 0;
for (let r = 0; r < numRows; r++) {
    for (let c = 0; c < numCols; c++) {
        const letter = m[r][c];
        if (letter == 'A') {
            if (r-1 >= 0 && r+1 < numRows && c-1 >=0 && c+1 < numCols) {
                // ld
                let left = m[r-1][c-1] + letter + m[r+1][c+1];
                let right = m[r+1][c-1] + letter + m[r-1][c+1];
                if ((left == 'MAS' || left == 'SAM') && (right == 'MAS' || right == 'SAM')) {
                    count++;
                }
            }
        }
    }
}
console.log('X-MAS count:', count);
