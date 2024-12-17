import sys
import re

a, b, c, *program = map(int, re.findall("\\d+", open(sys.argv[1]).read()))


def combo(operand):
    match operand:
        case 0 | 1 | 2 | 3:
            return operand
        case 4:
            return a
        case 5:
            return b
        case 6:
            return c
        case _:
            raise RuntimeError(f"unknown operand {operand}")


pointer = 0
output = []
while pointer < len(program):
    code, operand = program[pointer], program[pointer + 1]
    pointer += 2

    match code:
        case 0:  # adv
            a = a >> combo(operand)
        case 1:  # bxl
            b = b ^ operand
        case 2:  # bst
            b = combo(operand) % 8
        case 3:  # jnz
            if a != 0:
                pointer = operand
                continue
        case 4:  # bxc
            b = b ^ c
        case 5:  # out
            output.append(combo(operand) % 8)
        case 6:  # bdv
            b = a >> combo(operand)
        case 7:  # cdv
            c = a >> combo(operand)

print(*output, sep=",")
