DROP DATABASE IF EXISTS `LIBRARY`

CREATE DATABASE `LIBRARY`;

USE `LIBRARY`;

CREATE Table `BOOKS` (
    book_id INT NOT NULL AUTO_INCREMENT,
    book_name VARCHAR(50) NOT NULL,
    price DECIMAL(5,2) NOT NULL DEFAULT '0',
    publish_date DATE,
    author_id INT,
    PRIMARY KEY (book_id),
    FOREIGN KEY(author_id) REFERENCES Authors(author_id)
);
INSERT INTO `BOOKS` (book_name, price, publish_date, author_id) VALUES
('The Great Gatsby', 9.99, '1925-04-10', 1),
('To Kill a Mockingbird', 12.99, '1960-07-11', 2),
('1984', 10.99, '1949-06-08', 3),
('Pride and Prejudice', 8.99, '1813-01-28', 4),
('The Catcher in the Rye', 11.99, '1951-07-16', 5),
('Lord of the Rings', 15.99, '1954-07-29', 6),
('Harry Potter and the Sorcerer''s Stone', 14.99, '1997-06-26', 7),
('The Hobbit', 13.99, '1937-09-21', 6),
('Brave New World', 10.99, '1932-01-01', 8),
('Animal Farm', 9.99, '1945-08-17', 3),
('The Da Vinci Code', 12.99, '2003-03-18', 9),
('The Hunger Games', 13.99, '2008-09-14', 10),
('The Chronicles of Narnia', 14.99, '1950-10-16', 11),
('The Alchemist', 11.99, '1988-01-01', 12),
('The Little Prince', 8.99, '1943-04-06', 13),
('Don Quixote', 12.99, '1605-01-01', 14),
('The Book Thief', 13.99, '2005-01-01', 15),
('Fahrenheit 451', 10.99, '1953-10-19', 16),
('Jane Eyre', 9.99, '1847-10-16', 17),
('The Odyssey', 8.99, '1890-01-01', 18),
('The Road', 12.99, '2006-09-26', 19),
('One Hundred Years of Solitude', 13.99, '1967-05-30', 20),
('The Grapes of Wrath', 11.99, '1939-04-14', 21),
('Moby Dick', 10.99, '1851-10-18', 22),
('War and Peace', 16.99, '1869-01-01', 23),
('The Divine Comedy', 12.99, '1320-01-01', 24),
('Crime and Punishment', 13.99, '1866-01-01', 25),
('The Adventures of Sherlock Holmes', 11.99, '1892-10-14', 26),
('Les Misérables', 14.99, '1862-01-01', 27),
('The Count of Monte Cristo', 13.99, '1844-01-01', 28),
('The Picture of Dorian Gray', 10.99, '1890-07-01', 29),
('Wuthering Heights', 9.99, '1847-12-01', 30),
('The Tale of Two Cities', 11.99, '1859-04-30', 31),
('The Old Man and the Sea', 9.99, '1952-09-01', 32),
('The Art of War', 8.99, '1910-01-01', 33),
('The Prince', 9.99, '1532-01-01', 34),
('The Republic', 12.99, '1900-01-01', 35),
('Dracula', 11.99, '1897-05-26', 36),
('Frankenstein', 10.99, '1818-01-01', 37),
('The Metamorphosis', 9.99, '1915-01-01', 38),
('The Trial', 11.99, '1925-01-01', 38),
('The Stranger', 10.99, '1942-01-01', 39),
('The Name of the Rose', 13.99, '1980-01-01', 40),
('The Master and Margarita', 12.99, '1967-01-01', 41),
('Doctor Zhivago', 13.99, '1957-01-01', 42),
('The Brothers Karamazov', 14.99, '1880-01-01', 25),
('Anna Karenina', 13.99, '1877-01-01', 23),
('Madame Bovary', 11.99, '1856-12-15', 43),
('The Canterbury Tales', 12.99, '1400-01-01', 44),
('Paradise Lost', 11.99, '1667-01-01', 45),
('The Aeneid', 10.99, '1900-01-01', 46);

CREATE Table `Authors` (
    author_id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50),
    books_publish INT,
    birth_date DATE,
    PRIMARY KEY(author_id)
);

INSERT INTO `Authors` (author_id, name, books_publish, birth_date) VALUES
(1, 'F. Scott Fitzgerald', 1, '1896-09-24'),
(2, 'Harper Lee', 1, '1926-04-28'),
(3, 'George Orwell', 2, '1903-06-25'),
(4, 'Jane Austen', 1, '1775-12-16'),
(5, 'J.D. Salinger', 1, '1919-01-01'),
(6, 'J.R.R. Tolkien', 2, '1892-01-03'),
(7, 'J.K. Rowling', 1, '1965-07-31'),
(8, 'Aldous Huxley', 1, '1894-07-26'),
(9, 'Dan Brown', 1, '1964-06-22'),
(10, 'Suzanne Collins', 1, '1962-08-10'),
(11, 'C.S. Lewis', 1, '1898-11-29'),
(12, 'Paulo Coelho', 1, '1947-08-24'),
(13, 'Antoine de Saint-Exupéry', 1, '1900-06-29'),
(14, 'Miguel de Cervantes', 1, '1547-09-29'),
(15, 'Markus Zusak', 1, '1975-06-23'),
(16, 'Ray Bradbury', 1, '1920-08-22'),
(17, 'Charlotte Brontë', 1, '1816-04-21'),
(18, 'Homer', 1, NULL),
(19, 'Cormac McCarthy', 1, '1933-07-20'),
(20, 'Gabriel García Márquez', 1, '1927-03-06'),
(21, 'John Steinbeck', 1, '1902-02-27'),
(22, 'Herman Melville', 1, '1819-08-01'),
(23, 'Leo Tolstoy', 2, '1828-09-09'),
(24, 'Dante Alighieri', 1, '1265-05-21'),
(25, 'Fyodor Dostoevsky', 2, '1821-11-11'),
(26, 'Arthur Conan Doyle', 1, '1859-05-22'),
(27, 'Victor Hugo', 1, '1802-02-26'),
(28, 'Alexandre Dumas', 1, '1802-07-24'),
(29, 'Oscar Wilde', 1, '1854-10-16'),
(30, 'Emily Brontë', 1, '1818-07-30'),
(31, 'Charles Dickens', 1, '1812-02-07'),
(32, 'Ernest Hemingway', 1, '1899-07-21'),
(33, 'Sun Tzu', 1, NULL),
(34, 'Niccolò Machiavelli', 1, '1469-05-03'),
(35, 'Plato', 1, NULL),
(36, 'Bram Stoker', 1, '1847-11-08'),
(37, 'Mary Shelley', 1, '1797-08-30'),
(38, 'Franz Kafka', 2, '1883-07-03'),
(39, 'Albert Camus', 1, '1913-11-07'),
(40, 'Umberto Eco', 1, '1932-01-05'),
(41, 'Mikhail Bulgakov', 1, '1891-05-15'),
(42, 'Boris Pasternak', 1, '1890-02-10'),
(43, 'Gustave Flaubert', 1, '1821-12-12'),
(44, 'Geoffrey Chaucer', 1, '1343-01-01'),
(45, 'John Milton', 1, '1608-12-09'),
(46, 'Virgil', 1, '70-10-15');


SELECT book_name FROM books WHERE book_id = 2
UNION
SELECT name FROM authors WHERE author_id = (SELECT author_id FROM books WHERE book_id = 2);

SELECT book_name FROM books WHERE book_id = 2
NATURAL JOIN SELECT name FROM authors WHERE author_id = book_id

SELECT b.book_id,b.book_name, b.price, b.publish_date, a.name as author_name
FROM books b
JOIN authors a ON b.author_id = a.author_id
WHERE b.book_id = 2;

