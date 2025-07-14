# ❌ Anti-Pattern: API inside a loop

user_ids = [1, 2, 3, 4]

for user_id in user_ids:
    r = requests.get(f"https://api.example.com/user/{user_id}")
    data = r.json()
    print(data["name"])


# ✅ Better: Batch Endpoint or Parallel Calls

# If the API supports batching
r = requests.post("https://api.example.com/users/batch", json={"ids": user_ids})
users = r.json()
for user in users:
    print(user["name"])
