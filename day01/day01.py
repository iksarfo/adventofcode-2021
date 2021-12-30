from queue import Queue

f = open("input.txt", "r")
lines = f.readlines()
stripped = map(str.strip, lines)
depths = list(map(int, stripped))

increases = 0
for index, item in enumerate(depths):
    if index > 0 and depths[index] > depths[index-1]:
        increases += 1

print('part 1 =', increases)
############################

increases = 0
window = 3
q = Queue(maxsize=window+1)

for index, item in enumerate(depths):
    q.put(item)

    if(q.full()):
        if sum(list(q.queue)[1:]) > sum(list(q.queue)[:-1]):
            increases += 1
        q.get()

print('part 2 =', increases)
############################
