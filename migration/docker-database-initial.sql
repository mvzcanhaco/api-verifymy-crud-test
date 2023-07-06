create table users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    address VARCHAR(255)
);

INSERT INTO users (name, age, email, password, address)
VALUES ('Ayrton Senna', 34, 'ayrton@senna.com', 'senna123', 'São Paulo, Brasil');

INSERT INTO users (name, age, email, password, address)
VALUES ('Gisele Bündchen', 41, 'gisele@bundchen.com', 'bundchen123', 'Rio Grande do Sul, Brasil');

INSERT INTO users (name, age, email, password, address)
VALUES ('Pelé', 80, 'pele@pele.com', 'pele123', 'Três Corações, Brasil');

INSERT INTO users (name, age, email, password, address)
VALUES ('Carmen Miranda', 46, 'carmen@miranda.com', 'miranda123', 'Portugal');
