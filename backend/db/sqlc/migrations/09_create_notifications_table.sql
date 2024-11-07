CREATE TABLE notifications (
    id VARCHAR(36) PRIMARY KEY,
    userId VARCHAR(36),
    type VARCHAR(50),
    message TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    isRead BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (userId) REFERENCES users(id)
);