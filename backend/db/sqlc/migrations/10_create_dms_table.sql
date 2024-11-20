CREATE TABLE dms (
    id VARCHAR(36) PRIMARY KEY, 
    senderId VARCHAR(36), 
    receiverId VARCHAR(36), 
    content TEXT, 
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    FOREIGN KEY (senderId) REFERENCES users(id), 
    FOREIGN KEY (receiverId) REFERENCES users(id)
);