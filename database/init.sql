CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    resource VARCHAR(50) NOT NULL,
    action VARCHAR(50) NOT NULL,
    description TEXT,
    UNIQUE(resource, action)
);

CREATE TABLE IF NOT EXISTS user_roles (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

CREATE TABLE IF NOT EXISTS role_permissions (
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INTEGER REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

-- Seed data
INSERT INTO roles (name, description) VALUES 
('admin', 'Super User'),
('user', 'Regular User');

INSERT INTO permissions (resource, action, description) VALUES
('users', 'read', 'Read user profiles'),
('users', 'write', 'Modify users'),
('roles', 'read', 'Read roles'),
('roles', 'write', 'Modify roles');

-- Give admin all permissions
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r CROSS JOIN permissions p WHERE r.name = 'admin';

-- Create an admin user automatically (password: admin123)
-- Hash generated from golang bcrypt string for 'admin123'
INSERT INTO users (username, email, password_hash) VALUES
('admin', 'admin@example.com', '$2a$10$kkqs/FMzzrm2SY0sW8mTpOZU9YIdon5aAYXQSW0X5nBWbjcV0QO3a');

INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id FROM users u CROSS JOIN roles r WHERE u.username = 'admin' AND r.name = 'admin';

-- Create a regular user
INSERT INTO users (username, email, password_hash) VALUES
('johndoe', 'johndoe@example.com', '$2a$10$kkqs/FMzzrm2SY0sW8mTpOZU9YIdon5aAYXQSW0X5nBWbjcV0QO3a');
INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id FROM users u CROSS JOIN roles r WHERE u.username = 'johndoe' AND r.name = 'user';
