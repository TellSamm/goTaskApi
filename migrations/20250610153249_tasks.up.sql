CREATE TABLE tasks (
        id VARCHAR(36) PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        status VARCHAR(50) NOT NULL,
        deleted_at TIMESTAMP DEFAULT NULL
);