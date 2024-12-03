CREATE TABLE dms (
    id VARCHAR(36) PRIMARY KEY, 
    sender_id VARCHAR(36), 
    receiver_id VARCHAR(36), 
    content TEXT, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE
);