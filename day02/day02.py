import pytest

def moves_in_direction(dir, line):
    move_dir = line.split(' ')
    if move_dir[0] == dir:
        return int(move_dir[1])
    else:
        return 0

def destination(input, dir):
    delta = 0
    for move in input:
        delta += moves_in_direction(dir, move)
    return delta

def horizontal_depth(input):
    depth = destination(input,'down') - destination(input,'up')
    return destination(input,'forward') * depth

example = """forward 5
down 5
forward 8
up 3
down 8
forward 2"""

assert horizontal_depth(example.split('\n')) == 150

f = open("input.txt", "r")
lines = f.readlines()
print('part 1 =', horizontal_depth(lines))
##########################################

# "instruction" more dscriptive than "line"

def new_aim(current_aim, instruction):
    current_aim -= moves_in_direction('up', instruction)
    current_aim += moves_in_direction('down', instruction)
    propelled = current_aim * moves_in_direction('forward', instruction)
    return current_aim, propelled

aim = 0
new_depth = 0
horizontal = 0
for instruction in lines:
    aim, depth = new_aim(aim, instruction)
    new_depth += depth
    horizontal += moves_in_direction('forward', instruction)

print('part 2 =', horizontal * new_depth)
#########################################