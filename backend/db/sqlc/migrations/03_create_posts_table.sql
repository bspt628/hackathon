CREATE TABLE posts (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    is_repost BOOLEAN DEFAULT FALSE,
    original_post_id VARCHAR(36) DEFAULT NULL,
    reply_to_id VARCHAR(36) DEFAULT NULL,
    root_post_id VARCHAR(36) DEFAULT NULL,
    is_reply BOOLEAN DEFAULT FALSE,
    media_urls JSON,
    likes_count INT DEFAULT 0,
    reposts_count INT DEFAULT 0,
    replies_count INT DEFAULT 0,
    views_count INT DEFAULT 0,
    visibility VARCHAR(20),
    is_pinned BOOLEAN DEFAULT FALSE,
    is_deleted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (original_post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (reply_to_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (root_post_id) REFERENCES posts(id) ON DELETE CASCADE
);