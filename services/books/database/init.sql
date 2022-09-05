CREATE TABLE books (
	id BIGINT PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY, 
	title VARCHAR(40) NOT NULL CHECK (title <> ''), 
	subtitle VARCHAR(30) NOT NULL CHECK (subtitle <> ''), 
	author VARCHAR(20) NOT NULL CHECK (author <> ''), 
	isbn VARCHAR(14) NOT NULL CONSTRAINT isbn_check CHECK (isbn ~* '\d{3}-\d{10}') UNIQUE, 
	edition INT NOT NULL, 
	year INT NOT NULL,
    shelf_name VARCHAR(50) NOT NULL CHECK (shelf_name <> ''),
    shelf_level INTEGER NOT NULL
);