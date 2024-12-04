import * as util from "../src/util.ts";

describe('testing getDataFromFile', () => {
  test('load text from test.txt', () => {
    expect(util.getDataFromFile('test/data/test.txt')).toBe("1 2 3\n4 5 6");
  });
});

describe('testing getLinesFromFile', () => {
  test('load lines from test.txt', () => {
    expect(util.getLinesFromFile('test/data/test.txt')).toStrictEqual(['1 2 3', '4 5 6']);
  });
});
