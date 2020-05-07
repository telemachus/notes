# Programming in Lua 3rd edition

## Chapter 10 Complete Examples

### The Eight-Queen Puzzle

The puzzle is this: how can you place eight queens on a chessboard so that
no queen can attack another one.

He immediately notes that any valid solution will have exactly one queen in
each row. As such, he chooses to represent solutions as a simple table with
eight numbers. Each index position of the table is a row, and the number
value at that index is the column. For example:

    {3, 7, 2, 1, 8, 6, 5, 4}

The queen in row 1 is at column 3, the queen in row 2 is at column 7, and
so on. (This solution isn't valid, by the way. It's just an example to
explain the datatype he chooses to solve the problem)
