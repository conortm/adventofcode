import { readFileSync } from 'fs';

export function getDataFromFile(path: string) {
    return readFileSync(path, 'utf8');
}

export function getLinesFromFile(path: string) {
    const data = getDataFromFile(path);
    return data.split('\n');
}

export function getNumberArrayFromString(s: string, separator: string = ' ') {
    const numberArray: number[] = [];
    s.split(separator).forEach((val) => {
        numberArray.push(+val)
    });
    return numberArray;
}
