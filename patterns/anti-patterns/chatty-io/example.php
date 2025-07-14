<?php

// ❌ Anti-Pattern: Querying inside a loop

$posts = Post::all();

foreach ($posts as $post) {
    // N+1 Problem
    $author = $post->author; // Triggers a separate query for each post
}


// ✅ Better: Eager Loading

$posts = Post::with('author')->get();

foreach ($posts as $post) {
    $author = $post->author; // Already loaded
}
