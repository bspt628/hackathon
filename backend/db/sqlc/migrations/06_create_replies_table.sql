CREATE TABLE replies (
    id VARCHAR(36) PRIMARY KEY,
    userId VARCHAR(36),
    postId VARCHAR(36),
    content TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES users(id),
    FOREIGN KEY (postId) REFERENCES posts(id)
);