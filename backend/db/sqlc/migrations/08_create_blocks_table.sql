CREATE TABLE blocks (
    id VARCHAR(36) PRIMARY KEY,
    blocked_by_id VARCHAR(36),
    blocked_user_id VARCHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (blocked_byId) REFERENCES users(id),
    FOREIGN KEY (blocked_userId) REFERENCES users(id)
);