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

INSERT INTO publications (title, content, authorId) 
VALUES 
("Publicação Higor K.", "Conteudo legal.", 1),
("Publicação Pedro R.", "Conteudo interessante.", 2),
("Publicação Luis Q.", "Conteudo engraçado.", 3);
