import Helper from './Helpers';

const helper = new Helper;

test('Has winner on game state "000000000000000000000000000000000000000000" and player "1"', () => {
    expect(helper.hasWinner("000000000000000000000000000000000000000000", "1", 38)).toBe(false);
});
