CREATE TABLE replies (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36),
    post_id VARCHAR(36),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);