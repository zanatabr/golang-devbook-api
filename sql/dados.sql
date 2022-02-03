-- senha 432322 ==> $2a$10$QUdgKrQUAwGKv4l3Dyv4G.WeIFPmD8SetKBmqK92L2bqJbiC.30b2

insert into usuarios (nome, nick, email, senha)
values
("Usuario 1", "usuario_1", "usuario1@summdfdx.com", "$2a$10$QUdgKrQUAwGKv4l3Dyv4G.WeIFPmD8SetKBmqK92L2bqJbiC.30b2"),
("Usuario 2", "usuario_2", "usuario2@summdfdx.com", "$2a$10$QUdgKrQUAwGKv4l3Dyv4G.WeIFPmD8SetKBmqK92L2bqJbiC.30b2"),
("Usuario 3", "usuario_3", "usuario3@summdfdx.com", "$2a$10$QUdgKrQUAwGKv4l3Dyv4G.WeIFPmD8SetKBmqK92L2bqJbiC.30b2");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);


