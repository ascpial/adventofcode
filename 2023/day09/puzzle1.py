with open("input.txt", 'r', encoding='utf-8') as file:
    data = file.readlines()

histories = [[int(i) for i in line.split(" ")] for line in data]

output = []

def predict_value(history: list[int]) -> int:
    differences = [history[i+1] - history[i] for i in range(len(history)-1)]
    
    if differences != [0] * len(differences):
        next_predicted_value = predict_value(differences)
        return next_predicted_value + history[-1]
    else:
        return history[-1]

for history in histories:
    predicted_sub_value = predict_value(history)
    output.append(predicted_sub_value)

print(output)
print(sum(output))
