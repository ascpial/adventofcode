print(
    sum(i for i in __import__("itertools").chain(*(range(int(curRange[0]), int(curRange[1])+1) for curRange in (curRange.split("-") for curRange in open('input.txt').read().strip().split(",")))) if len(str(i)) %2 == 0 and str(i)[:len(str(i))//2] == str(i)[len(str(i))//2:] )
)
