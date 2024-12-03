import { readFileSync } from 'fs';

export function getDataFromFile(path: string) {
    return readFileSync(path, 'utf8');
}

export function getLinesFromFile(path: string) {
    const data = getDataFromFile(path)
    return data.split('\n')
}
