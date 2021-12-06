import sys


def main():
    raw_input_data = sys.stdin.read().split('\n\n')
    # print(raw_input_data)
    raw_numbers, raw_boards = raw_input_data[0], raw_input_data[1:]
    drawn_numbers = [int(x) for x in raw_numbers.split(',')]

    boards = []
    for raw_board in raw_boards:
        board = {}
        lines = raw_board.split('\n')
        for r, line in enumerate(lines):
            for c, val in enumerate(line.split()):
                board[int(val)] = (r, c)
        boards.append(board)

    marked_boards = [set() for _ in boards]

    for num in drawn_numbers:
        for i, board in enumerate(boards):
            if num in board:
                marked_boards[i].add(board[num])
                if is_win(marked_boards[i]):
                    sum_of_unmarked = sum(val for val in board if board[val] not in marked_boards[i])
                    print(sum_of_unmarked * num)
                    return


def is_win(b):
    # Check rows
    for r in range(5):
        bingo = True
        for c in range(5):
            if (r, c) not in b:
                bingo = False
                break
        if bingo:
            return True

    # Check cols
    for c in range(5):
        bingo = True
        for r in range(5):
            if (r, c) not in b:
                bingo = False
                break
        if bingo:
            return True

    # # check down-left diagonal
    # bingo = True
    # for r, c in zip(list(range(5)), list(range(5))):
    #     if (r, c) not in b:
    #         bingo = False
    #         break
    # if bingo:
    #     return True

    # # check down-right diagonal
    # bingo = True
    # for r, c in zip(list(range(5)), list(range(4, -1, -1))):
    #     if (r, c) not in b:
    #         bingo = False
    #         break
    # if bingo:
    #     return True

    return False



main()
