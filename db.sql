CREATE TABLE IF NOT EXISTS `users` (
   id INT AUTO_INCREMENT PRIMARY KEY,
   email VARCHAR(50) UNIQUE NOT NULL,
   password VARCHAR(255) NOT NULL,
   name VARCHAR(50) NOT NULL,
   role ENUM ("admin","user") DEFAULT "user",
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO `users` (`email`, `password`, `name`, `role`) VALUES ('admin@mail.com', '$2a$10$K7UWGDmSU8N6JfvzCoq3lu5816huk8xyOZrrsRiu6wKCCoV2RnBeC', 'Admin', 'admin');

-- table book

CREATE TABLE IF NOT EXISTS `books` (
   id INT AUTO_INCREMENT PRIMARY KEY,
   title VARCHAR(255) NOT NULL,
   author VARCHAR(255) NOT NULL,
   publication_year INT,
   genre VARCHAR(100),
   isbn VARCHAR(20) UNIQUE,
   stok INT NOT NULL DEFAULT 0,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `transactions` (
	id INT AUTO_INCREMENT PRIMARY KEY,
	user_id INT NOT NULL,
	book_id INT NOT NULL,
	duration INT NOT NULL COMMENT "borrow duration in day",
	start_date DATETIME NOT NULL,
	end_date DATETIME NOT NULL,
	`status` ENUM("BORROWING", "DONE", "LATE") DEFAULT "BORROWING",
	is_late BOOLEAN DEFAULT FALSE,
	return_date DATETIME NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Add foreign key constraint for user_id
ALTER TABLE transactions
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id) REFERENCES users(id);

-- Add foreign key constraint for book_id
ALTER TABLE transactions
ADD CONSTRAINT fk_book_id
FOREIGN KEY (book_id) REFERENCES books(id);