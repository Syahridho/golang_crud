CREATE DATABASE IF NOT EXISTS golang_crud;

USE golang_crud;

CREATE TABLE IF NOT EXISTS employee (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    npwp VARCHAR(50) NOT NULL,
    address TEXT NOT NULL
);

-- Insert some sample data for testing
INSERT INTO employee (name, npwp, address) VALUES 
('John Doe', '123456789012345', '123 Main St, City'),
('Jane Smith', '987654321098765', '456 Oak Ave, Town'),
('Bob Johnson', '456789012345678', '789 Pine Rd, Village');