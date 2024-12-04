import requests
import datetime

with open("cookie.txt", 'r') as file:
    cookie = file.read().strip()
with open("leaderboard.txt", 'r') as file:
    leaderboard = file.read().strip()

data = requests.get(f"https://adventofcode.com/2024/leaderboard/private/view/{leaderboard}.json", cookies={"session": cookie}).json()

day = input("Quel jour afficher ? ")

members = sorted([member for id, member in data["members"].items()], key=lambda data: data["local_score"], reverse=True)

for member in members:
    score = member["local_score"]
    name = member["name"]
    stars = member["stars"]
    print(f"{score:<3} {name} ({stars}٭)")
    stars_data = member["completion_day_level"].get(day)
    if stars_data:
        for star, data in stars_data.items():
            print(f"  - {star} : {datetime.datetime.fromtimestamp(data['get_star_ts']).strftime("%m/%d/%Y, %H:%M:%S")}")
