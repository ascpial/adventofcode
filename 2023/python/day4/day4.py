with open("input.txt", 'r', encoding='utf-8') as file:
    text = file.readlines()

# Parses the content of a card in the following format
# (id, wining_numbers, numbers_you_have)

cards = []

for line in text:
    line = line.strip()
    raw_id, raw_numbers = line.split(": ")
    id = int(raw_id[4:].strip())

    wining_numbers, your_numbers = raw_numbers.split("|")
    wining_numbers = [int(number) for number in wining_numbers.strip().split(" ") if number != '']
    your_numbers = [int(number) for number in your_numbers.strip().split(" ") if number != '']

    cards.append((id, wining_numbers, your_numbers))

if False:
    scores = []

    for id, wining_numbers, your_numbers in cards:
        score = 0
        for number in your_numbers:
            if number in wining_numbers:
                if score == 0:
                    score = 1
                else:
                    score *= 2
        scores.append(score)

    print(sum(scores))

if True:
    # On génère une liste avec l'identifiant des scratchcards et le nombre de cartes qu'elle rapporte
    scores = {}
    for id, wining_numbers, your_numbers in cards:
        score = 0
        for number in your_numbers:
            if number in wining_numbers:
                score += 1
        scores[id] = score

    # On génère un dictionnaire contenant le nombre de carte de base
    cards_amount = {id: 1 for id,_,_ in cards}

    MAX = max(cards_amount.keys())

    def process_cards(card, scores, cards_amount):
        for i in range(scores[card]):
            card_id = card + i + 1
            if card_id > MAX:
                continue
            cards_amount[card_id] += 1
            process_cards(card_id, scores, cards_amount)

    for card in cards_amount.keys():
        print(card)
        process_cards(card, scores, cards_amount)

    print(cards_amount)
    print(sum([amount for _, amount in cards_amount.items()]))
