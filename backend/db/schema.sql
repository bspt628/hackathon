-- schema.sql

CREATE DATABASE IF NOT EXISTS social_media_platform;
USE social_media_platform;

CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    display_name VARCHAR(100),
    profile_image_url VARCHAR(255),
    bio TEXT,
    location VARCHAR(100),
    website VARCHAR(255),
    birth_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    last_login_at TIMESTAMP,
    is_verified BOOLEAN DEFAULT FALSE,
    is_private BOOLEAN DEFAULT FALSE,
    is_banned BOOLEAN DEFAULT FALSE,
    followers_count INT DEFAULT 0,
    following_count INT DEFAULT 0,
    posts_count INT DEFAULT 0,
    likes_count INT DEFAULT 0,
    language VARCHAR(10),
    theme VARCHAR(20),
    notification_settings JSON,
    two_factor_enabled BOOLEAN DEFAULT FALSE,
    last_password_change TIMESTAMP
);

CREATE TABLE posts (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    is_repost BOOLEAN DEFAULT FALSE,
    original_post_id VARCHAR(36),
    reply_to_id VARCHAR(36),
    root_post_id VARCHAR(36),
    is_reply BOOLEAN DEFAULT FALSE,
    media_urls JSON,
    likes_count INT DEFAULT 0,
    reposts_count INT DEFAULT 0,
    replies_count INT DEFAULT 0,
    views_count INT DEFAULT 0,
    visibility VARCHAR(20),
    language VARCHAR(10),
    location VARCHAR(100),
    device VARCHAR(50),
    is_pinned BOOLEAN DEFAULT FALSE,
    is_deleted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (original_post_id) REFERENCES posts(id),
    FOREIGN KEY (reply_to_id) REFERENCES posts(id),
    FOREIGN KEY (root_post_id) REFERENCES posts(id)
);

CREATE TABLE reposts (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36),
    original_post_id VARCHAR(36),
    reposted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_quote_repost BOOLEAN DEFAULT FALSE,
    additional_comment TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (original_post_id) REFERENCES posts(id)
);

CREATE TABLE likes (
    id VARCHAR(36) PRIMARY KEY,
    userId VARCHAR(36),
    postId VARCHAR(36),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES users(id),
    FOREIGN KEY (postId) REFERENCES posts(id)
);

CREATE TABLE replies (
    id VARCHAR(36) PRIMARY KEY,
    userId VARCHAR(36),
    postId VARCHAR(36),
    content TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES users(id),
    FOREIGN KEY (postId) REFERENCES posts(id)
);

CREATE TABLE follows (
    id VARCHAR(36) PRIMARY KEY,
    followerId VARCHAR(36),
    followingId VARCHAR(36),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (followerId) REFERENCES users(id),
    FOREIGN KEY (followingId) REFERENCES users(id)
);

CREATE TABLE blocks (
    id VARCHAR(36) PRIMARY KEY,
    blockedById VARCHAR(36),
    blockedUserId VARCHAR(36),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (blockedById) REFERENCES users(id),
    FOREIGN KEY (blockedUserId) REFERENCES users(id)
);

CREATE TABLE notifications (
    id VARCHAR(36) PRIMARY KEY,
    userId VARCHAR(36),
    type VARCHAR(50),
    message TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    isRead BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (userId) REFERENCES users(id)
);

CREATE TABLE dms (
    id VARCHAR(36) PRIMARY KEY,
    senderId VARCHAR(36),
    receiverId VARCHAR(36),
    content TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (senderId) REFERENCES users(id),
    FOREIGN KEY (receiverId) REFERENCES users(id)
);