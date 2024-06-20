CREATE TABLE phones
(
	id     SERIAL PRIMARY KEY,
	number varchar(255) UNIQUE NOT NULL
);
COMMENT ON TABLE phones IS 'Таблица телефонов';

INSERT INTO phones
	(number)
values ('1234567890'),
			 ('123 456 7891'),
			 ('(123) 456 7892'),
			 ('(123) 456-7893'),
			 ('123-456-7894'),
			 ('123-456-7890'),
			 ('1234567892'),
			 ('(123)456-7892');
