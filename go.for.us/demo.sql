-- demo.sql

-- Схема: отдельные списки для названий, лидеров, жанров + таблица bands
DROP TABLE IF EXISTS bands;
DROP TABLE IF EXISTS names;
DROP TABLE IF EXISTS leaders;
DROP TABLE IF EXISTS genres;

CREATE TABLE names (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE leaders (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE genres (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE bands (
  id SERIAL PRIMARY KEY,
  name_id INTEGER NOT NULL REFERENCES names(id),
  leader_id INTEGER NOT NULL REFERENCES leaders(id),
  genre_id INTEGER NOT NULL REFERENCES genres(id),
  album_count INTEGER NOT NULL
);

-- Наполним names (15 строк)
INSERT INTO names (name) VALUES
('Звёздный Пульс'),
('Ночные Фонари'),
('Сердце Прерий'),
('Бетон и Сон'),
('Эхо Улиц'),
('Голоса Ветра'),
('Кремниевые Сны'),
('Лунный Причал'),
('Ржавые Крылья'),
('Пески Времени'),
('Вишнёвый Ветер'),
('Северное Солнце'),
('Сумеречный Парад'),
('Золотая Дыра'),
('Тихие Гитары');

-- Наполним leaders (15 строк)
INSERT INTO leaders (name) VALUES
('Алексей Миронов'),
('Ольга Рождественская'),
('Иван Петренко'),
('Марина Соколова'),
('Дмитрий Ланской'),
('Наталья Ветрова'),
('Пётр Голубев'),
('Екатерина Смирнова'),
('Константин Лебедев'),
('Людмила Дыбова'),
('Роман Кузнецов'),
('Вера Орлова'),
('Сергей Дьяков'),
('Алина Степанова'),
('Юрий Басов');

-- Наполним genres (15 строк)
INSERT INTO genres (name) VALUES
('Рок'),
('Поп'),
('Инди'),
('Электроника'),
('Фолк'),
('Металл'),
('Джаз'),
('Блюз'),
('Панк'),
('Хип-хоп'),
('Альтернатива'),
('Регги'),
('Кантри'),
('Классика'),
('Экспериментал');

-- Наполним bands (15 строк), ссылки на выше таблицы
INSERT INTO bands (name_id, leader_id, genre_id, album_count) VALUES
(1, 1, 1, 5),
(2, 2, 2, 3),
(3, 3, 5, 7),
(4, 4, 1, 2),
(5, 5, 4, 6),
(6, 6, 3, 4),
(7, 7, 8, 10),
(8, 8, 4, 1),
(9, 9, 6, 8),
(10, 10, 11, 2),
(11, 11, 2, 9),
(12, 12, 13, 3),
(13, 13, 12, 4),
(14, 14, 14, 0),
(15, 15, 15, 1);
