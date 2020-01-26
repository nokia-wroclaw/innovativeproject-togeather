CREATE TABLE IF NOT EXISTS restaurants
(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    delivery REAL NOT NULL
);

CREATE TABLE IF NOT EXISTS meals
(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price REAL NOT NULL,
    description TEXT,
    owning_restaurant INT NOT NULL,

    FOREIGN KEY (owning_restaurant) REFERENCES restaurants(id)
);

CREATE TABLE IF NOT EXISTS lobbies
(
    id SERIAL PRIMARY KEY,
    restaurant INT NOT NULL,
    expires timestamp NOT NULL,
    geolat double precision NOT NULL,
    geolon double precision NOT NULL,
    address TEXT NOT NULL,

    FOREIGN KEY (restaurant) REFERENCES restaurants(id)
);

CREATE TABLE IF NOT EXISTS clients
(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS orders
(
    id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    lobby_id INT NOT NULL,
    meal_id INT NOT NULL,

    FOREIGN KEY (lobby_id) REFERENCES lobbies(id),
    FOREIGN KEY (meal_id) REFERENCES meals(id),
    FOREIGN KEY (client_id) REFERENCES clients(id)
);
