import * as fs from 'fs';

// PART 1

fs.readFile('data/day01.txt', 'utf8', (err, data) => {
  if (err) {
    console.error('Error reading the file: ' + err);
    return;
  }
  let list1: number[] = [];
  let list2: number[] = [];
  data.split('\n').forEach( line => {
    const pair = line.split('   ');
    list1.push(+pair[0])
    list2.push(+pair[1])
  })
  list1.sort()
  list2.sort()
  let distance = 0;
  const len = list1.length;
  for (let i = 0; i < len; i++) {
    distance += Math.abs(list1[i] - list2[i]);
  }
  console.log('distance:', distance);
});

// PART 2

fs.readFile('data/day01.txt', 'utf8', (err, data) => {
    if (err) {
      console.error('Error reading the file: ' + err);
      return;
    }
    let list1: number[] = [];
    let list2: number[] = [];
    data.split('\n').forEach( line => {
      const pair = line.split('   ');
      list1.push(+pair[0])
      list2.push(+pair[1])
    })
    list1.sort()
    const len = list1.length;
    list2.sort()
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
  });
  