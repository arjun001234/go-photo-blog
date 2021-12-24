
DROP DATABASE GO_PHOTO_BLOG;

CREATE DATABASE GO_PHOTO_BLOG;

USE GO_PHOTO_BLOG;

CREATE TABLE IF NOT EXISTS PB_USERS(
    id INT AUTO_INCREMENT PRIMARY KEY,
    pb_name VARCHAR(25) NOT NULL,
    pb_email VARCHAR(50) NOT NULL UNIQUE,
    pb_password VARCHAR(255) NOT NUll,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS PB_SESSIONS(
    pb_session VARCHAR(50) PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    pb_user_id INT NOT NULL,
    FOREIGN KEY (pb_user_id)  REFERENCES PB_USERS(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS PB_PHOTOS(
    id INT AUTO_INCREMENT PRIMARY KEY,
    pb_url VARCHAR(100) NOT NULL UNIQUE,
    pb_user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (pb_user_id)  REFERENCES PB_USERS(id) ON DELETE CASCADE
);