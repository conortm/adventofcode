import * as util from "./util.ts";

const lines = util.getLinesFromFile('data/day01.txt')
const len = lines.length;
let list1: number[] = [];
let list2: number[] = [];
lines.forEach(line => {
  const pair = line.split('   ');
  list1.push(+pair[0])
  list2.push(+pair[1])
})
list1.sort()
list2.sort()

// PART 1
let distance = 0;
for (let i = 0; i < len; i++) {
  distance += Math.abs(list1[i] - list2[i]);
}
console.log('distance:', distance);

// PART 2
let list2Frequency = new Map<number, number>();
list2.forEach((num) => {
  let freq = list2Frequency.get(num) || 0;
  list2Frequency.set(num, freq + 1);
});
let similarityScore = 0;
list1.forEach((num) => {
  let freq = list2Frequency.get(num) || 0;
  similarityScore += num * freq
});
console.log('similarityScore:', similarityScore);
