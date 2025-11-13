"""
On note t entier naturel le temps que dure la course.
Soit c un entier naturel correspondant au temps pour lequel on maintient le bouton.
Alors la distance d parcourue par le bateau sera de d = c * (l - c).
On note r le dernier record.
On résout : c * (t - c) > r.
        <=> -c ** 2 + t * c - r > 0
On pose d = t**2 - 4 * r

Les racines sont :
x1 = (t - sqrt(d)) / 2
x2 = (t + sqrt(d)) / 2

Les solutions sont comprises dans l'intervalle : [x1, x2] car x1 < x2

Il y a donc x1 - x2 solutions.

On implémente ça maintenant,
en se concentrant sur le traitement automatique du fichier :
"""

import math

with open("input.txt", 'r', encoding='utf-8') as file:
    input = file.readlines()

entries = []
for line in input:
    entries.append(int(''.join(line.split(":")[1].split(" "))))

print(entries)

time = entries[0]
record = entries[1]

delta = time ** 2 - 4 * record

x1 = math.ceil((time - math.sqrt(delta)) / 2)
x2 = math.floor((time + math.sqrt(delta)) / 2)

print(x1, x2)
print(x2 - x1 + 1)
