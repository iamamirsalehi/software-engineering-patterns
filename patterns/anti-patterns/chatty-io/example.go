package main

// ❌ Anti-Pattern: One Query Per Item

for _, id := range ids {
    row := db.QueryRow("SELECT name FROM users WHERE id = ?", id)
    var name string
    err := row.Scan(&name)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name)
}

// ✅ Better: Use IN (...)

placeholders := strings.Repeat("?,", len(ids))
placeholders = placeholders[:len(placeholders)-1] // remove last comma

query := fmt.Sprintf("SELECT id, name FROM users WHERE id IN (%s)", placeholders)
args := make([]any, len(ids))
for i, v := range ids {
    args[i] = v
}

rows, err := db.Query(query, args...)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
    fmt.Printf("%d => %s\n", id, name)
}
