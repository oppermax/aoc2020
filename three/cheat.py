def getLines(filename):
    infile = open(filename, "r")
    return infile.read().split("\n")

def main():
    filename = "input"
    lines = getLines(filename)
    treeCnt = 0
    x = 0

    for y in range(0, len(lines)):
        x = x % 31
        treeCnt += 1 if lines[y][x] == "#" else 0
        x += 3

    print("Tree Count: ", treeCnt)

main()