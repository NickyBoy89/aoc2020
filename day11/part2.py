import copy

with open("input") as inputFile:
    input = inputFile.read().split("\n")[:-1]

    input = list(map(list, input))

changeIndexes = []

def numCanSee(rowIndex, colIndex):
    count = 0

    for check in range(rowIndex - 1, rowIndex + 2): # Loop through the 3x3 grid of surrounding seats
        for seat in range(colIndex - 1, colIndex + 2):
            curRow = check
            curCol = seat
            if (check == rowIndex and seat == colIndex): # Ignore the original seat
                continue
            if (check < 0 or seat < 0 or check >= len(input) or seat >= len(input[check])): # Handles OOB as no seat
                continue
            while (curRow >= 0 and curCol >= 0 and curRow < len(input) and curCol < len(input[curRow])): # Out-of-bounds
                if (input[curRow][curCol] != "."):
                    if (input[curRow][curCol] == "#"): # Handles empty seats as not incrementing seats counter
                        count += 1
                    break
                else: # Increment the searching counters
                    curRow -= rowIndex - check
                    curCol -= colIndex - seat
    return count

def prettyPrint(table):
    for i in list(map("".join, table)):
        print(i)

previousLayout = []

while (input != previousLayout):
# for _ in range(2):
    previousLayout = copy.deepcopy(input)
    for row in range(len(input)):
        for col in range(len(input[row])):
            if (input[row][col] == "L"): # Empty seat
                if (numCanSee(row, col) == 0):
                    changeIndexes.append(["#", row, col])
            if (input[row][col] == "#"): # Full seat
                if (numCanSee(row, col) >= 5):
                    changeIndexes.append(["L", row, col])

    for change in changeIndexes: # Make all the changes at once
        input[change[1]][change[2]] = change[0]
    changeIndexes.clear()

    # print("Shifted")

prettyPrint(input)

count = 0

for i in input:
    for j in i:
        if j == "#":
            count += 1

print(count)
