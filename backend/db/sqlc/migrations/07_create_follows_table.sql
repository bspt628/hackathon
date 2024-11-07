CREATE TABLE follows (
    id VARCHAR(36) PRIMARY KEY,
    followerId VARCHAR(36),
    followingId VARCHAR(36),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (followerId) REFERENCES users(id),
    FOREIGN KEY (followingId) REFERENCES users(id)
);