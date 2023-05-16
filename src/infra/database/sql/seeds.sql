INSERT INTO users (name, nickName, email, password) 
VALUES 
("Higor k.", "koakovski", "hkoakovskip@gmail.com", "$2a$10$XLt1xr7rm3OQavA7hC3MxOGbIj4j6kSiYuYXLtsQv.P02pigEvYTW"),
("Pedro R.", "rosa", "pedrorosa@gmail.com", "$2a$10$XLt1xr7rm3OQavA7hC3MxOGbIj4j6kSiYuYXLtsQv.P02pigEvYTW"),
("Luis Q.", "queiroz", "luisqueiroz@gmail.com", "$2a$10$XLt1xr7rm3OQavA7hC3MxOGbIj4j6kSiYuYXLtsQv.P02pigEvYTW");

-- passwords: Lace-4242

INSERT INTO followers (userId, followerId) 
VALUES 
(1, 2),
(3, 1),
(1, 3);
