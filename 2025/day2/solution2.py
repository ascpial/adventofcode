with open('input.txt', 'r') as input:
    result = 0
    for curRange in input.read().strip().split(","):
        ends = curRange.split("-")
        start = int(ends[0])
        end = int(ends[1])
        for cur in range(start, end+1):
            base10 = str(cur)
            for k in range(1, len(base10)//2+1):
                if len(base10)%k==0 and len(base10)//k>=2:
                    repeat = True
                    for l in range(len(base10)//k-1):
                        repeat = repeat and base10[l*k:(l+1)*k] == base10[(l+1)*k:(l+2)*k]
                    if repeat:
                        result = result + cur
                        break
    print(result)

