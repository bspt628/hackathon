CREATE TABLE blocks (
    id VARCHAR(36) PRIMARY KEY,
    blocked_by_id VARCHAR(36),
    blocked_user_id VARCHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (blocked_by_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (blocked_user_id) REFERENCES users(id) ON DELETE CASCADE
);