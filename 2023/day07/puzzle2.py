from functools import cmp_to_key

with open("input.txt", 'r', encoding='utf-8') as file:
    data = file.readlines()

# Format: [(hand, bid, number_of_different_cards, max_occurences)]

JOKER = "J"
MAP = {
    "J":1,
    "2":2,
    "3":3,
    "4":4,
    "5":5,
    "6":6,
    "7":7,
    "8":8,
    "9":9,
    "T":10,
    "Q":11,
    "K":12,
    "A":13,
}

hands = []

for line in data:
    raw_hand_data, bid = line.strip().split(" ")
    bid = int(bid)
    cards = {}
    jokers = 0
    for card in raw_hand_data:
        if card == JOKER:
            jokers += 1
        else:
            if card in cards:
                cards[card] += 1
            else:
                cards[card] = 1

    hand_data = [MAP[i] for i in raw_hand_data]

    hands.append((
        hand_data,
        bid,
        1 if len(cards) == 0 else len(cards),
        jokers if len(cards) == 0 else max(cards.values()) + jokers,
        raw_hand_data,
        ))

def compare(hand1, hand2):
    """Return True if hand1 is better than hand2"""
    if hand1[3] != hand2[3]:
        return hand1[3] > hand2[3]
    if hand1[2] != hand2[2]:
        return hand1[2] < hand2[2]
    
    for i in range(len(hand1)):
        if hand1[0][i] != hand2[0][i]:
            return hand1[0][i] > hand2[0][i]

    raise ValueError("Tried comparing identical hands")


def triBulle(L):
    n_c = 0
    for k in range(len(L)):
        for i in range(len(L) - k - 1):
            j = i + 1
            n_c += 1
            if compare(L[i], L[j]):
                L[i], L[j] = L[j], L[i]
    return n_c

triBulle(hands)
print(hands)

total_winnings = 0
for i, hand in enumerate(hands):
    print(hand)
    total_winnings += hand[1] * (i + 1)

print(total_winnings)
