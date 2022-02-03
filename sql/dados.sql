-- senha 123456 ==> $2a$10$PVcz41lDbJ6fFlu1VnoQaevN/rgVn5gQj7nscRCdXWuI7cLoYR06m

insert into usuarios (nome, nick, email, senha)
values
("Usuario 1", "usuario_1", "usuario1@summdfdx.com", "$2a$10$PVcz41lDbJ6fFlu1VnoQaevN/rgVn5gQj7nscRCdXWuI7cLoYR06m"),
("Usuario 2", "usuario_2", "usuario2@summdfdx.com", "$2a$10$PVcz41lDbJ6fFlu1VnoQaevN/rgVn5gQj7nscRCdXWuI7cLoYR06m"),
("Usuario 3", "usuario_3", "usuario3@summdfdx.com", "$2a$10$PVcz41lDbJ6fFlu1VnoQaevN/rgVn5gQj7nscRCdXWuI7cLoYR06m");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);


