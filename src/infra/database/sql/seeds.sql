INSERT INTO users (name, nickName, email, password) 
VALUES 
("Higor k.", "koakovski", "hkoakovskip@gmail.com", "$2a$10$KJUSbHT6IXCNHBH2SrWtKOzx3XOAj2JZN2oArwqBD7ftwMWpbJxmi"),
("Pedro R.", "rosa", "pedrorosa@gmail.com", "$2a$10$KJUSbHT6IXCNHBH2SrWtKOzx3XOAj2JZN2oArwqBD7ftwMWpbJxmi"),
("Luis Q.", "queiroz", "luisqueiroz@gmail.com", "$2a$10$KJUSbHT6IXCNHBH2SrWtKOzx3XOAj2JZN2oArwqBD7ftwMWpbJxmi");

INSERT INTO followers (userId, followerId) 
VALUES 
(1, 2),
(3, 1),
(1, 3);
