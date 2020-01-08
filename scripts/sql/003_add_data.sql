INSERT INTO restaurants(name, address, delivery)
VALUES	('Pizza Station', 'Grabiszyńska 66E', 5.00),
	('Subway', 'Legnicka 58', 3.50),
	('Pasibus', 'Świdnicka 11', 6.00),
	('Djerba Kebab', 'Grabiszyńska 66B', 2.50),
	('McDonalds', 'Powstańców Śląskich 95', 8.00),
	('KFC','Legnicka 58', 2.00),
	('Hoshi Sushi', 'Strzegomska 3b/3c', 4.50),
	('Bar Mleczny Miś', 'Kuźnicza 48', 7.00);

INSERT INTO meals(name, price, description, owning_restaurant)
VALUES	('Pizza Hawaii', 19.99, 'Pizza z szynką, ananasem i mandarynkami', 1),
	('Pizza California', 25.49, 'Pizza z pieczarkami, salami pepperoni, kurczakiem i kukurydzą', 1),
	('Pizza Alabama', 22.45, 'Pizza z pieczarkami, szynką, kukurydzą, brokułami i pomidorem', 1),
	('Sandwich Spicy Intallian', 15.99, 'Kanapka z salami i peperoni', 2),
	('Sandwich Kurczak Teriyaki', 12.99, 'Kanapka z piersią z kurczaka marynowaną w zalewie teriyaki', 2),
	('Sandwich Triple Becon', 10.99, 'Kanapka z potrójną porcją bekonu', 2),
	('Burger Standard', 20.99, 'Burger z wołowiną, sałatą, serem cheddar, ogórkiem, pomidorem i cebulą', 3),
	('Burger Bebek Junior', 17.99, 'Burger z wołowiną, sałatą, sosem BBQ, serem i boczkiem', 3),
	('Burger Gonzales', 19.99, 'Burger z wołowiną, rukolą, pastą curry i jalapeno', 3),
	('Tortilla z kurczakiem', 10.99, 'Tortilla z kurczakiem, sosem i warzywami', 4),
	('Pita z baraniną', 15.99, 'Pita z baraniną, sosem i warzywami', 4),
	('Zestaw Gyros', 17.00, 'Zestaw zawierający mięso z kurczaka, frytki, sałatkę i sos', 4),
	('Kurczakburger', 3.99, 'Burger z kurczakiem, sałatą i ostrym sosem', 5),
	('Frytki', 4.49, 'Frytki z ziemniaków', 5),
	('McZestaw', 19.99, 'Zestaw zawietający burgera McRoyal, frytki i napój', 5),
	('Longer', 3.99, 'Kanapka z kurczakiem, sałatą i ostrym sosem', 6),
	('Mega Pocket', 14.49, 'Tortilla z kurczakiem, warzywami, serem i ostrym sosem', 6),
	('Frytki', 2.99, 'Frytki z ziemniaków', 6),
	('Tai Roll', 15.99, 'Roll zawierający dorada sous vide / tymianek / ogórek / yuzu', 7),
	('Spicy Shrimps', 24.99, 'Roll zawierający krewetki w pikantnej tempurze / łosoś / kanpyo / oshinko', 7),
	('Special Dragon', 32.00, 'Roll zawierający krewetki w tempurze / grillowany węgorz / awokado / ogórek', 7),
	('Pierogi Ruskie', 5.99, 'Pierogi z mięsem i cebulką', 8),
	('Naleśniki', 3.99, 'Naleśniki z serem i śmietaną', 8),
	('Kwaśne Mleko', 1.99, 'Szklanka zsiadłego mleka', 8);

INSERT INTO clients(name)
VALUES 	('Robert Lewandowski'),
	('Arkadiusz Milik'),
	('Grzegorz Krychowiak'),
	('Paweł Talar'),
	('Piotr Zieliński'),
	('Tomasz Siemieniuk');
