import sys
sys.setrecursionlimit(10000)

# théoriquement, si la stack était plus grande par défaut et que je rendait le code très moche,
# ça tiendrait en une ligne
print(
    (
    lambda a:
        (lambda f: f(f, 50, 0))(
            lambda f, v, i: int(v == 0) if i == len(a) else f(f, (v+a[i])%100, i+1) + int(v == 0)
        )
    )([int(i) for i in open('input.txt').read().strip().replace("L", "").replace("R", "-").split("\n")])
)

