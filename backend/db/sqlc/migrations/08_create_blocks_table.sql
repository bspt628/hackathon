CREATE TABLE blocks (
    id VARCHAR(36) PRIMARY KEY,
    blockedById VARCHAR(36),
    blockedUserId VARCHAR(36),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (blockedById) REFERENCES users(id),
    FOREIGN KEY (blockedUserId) REFERENCES users(id)
);