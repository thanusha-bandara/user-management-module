-- 1. Create Roles Table
CREATE TABLE roles (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL
);


INSERT INTO roles (role_name) VALUES 
('Admin'), 
('Sales_Agent'), 
('Manager'), 
('Customer');

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    full_name VARCHAR(150) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role_id INTEGER REFERENCES roles(role_id) ON DELETE RESTRICT,
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert a default Admin user (Password: admin123)
INSERT INTO users (full_name, username, email, password_hash, role_id, status)
VALUES ('System Admin', 'admin', 'admin@gmail.com', '$2a$10$kD5ZoBNDcsBnLFvFJAHG1eJxopAtoTRAGC5McTr9/gxHU65b110uq', 1, 'active');