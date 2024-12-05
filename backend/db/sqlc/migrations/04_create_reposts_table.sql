CREATE TABLE reposts (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36),
    original_post_id VARCHAR(36),
    reposted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_quote_repost BOOLEAN DEFAULT FALSE,
    additional_comment TEXT,
    UNIQUE (user_id, original_post_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (original_post_id) REFERENCES posts(id) ON DELETE CASCADE
);