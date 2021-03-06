-- Table User --
INSERT INTO public."user"
(id, user_name, full_name, "password", gender, hobby, address, created_at, updated_at, deleted_at)
VALUES(1, 'yusuf', 'Ade Yusuf', '$2a$08$rPq9Pw1iK42J09jY4TyRhuD0KadEzD04olXS46KcXfQJnQyZ/hzue', NULL, NULL, NULL, '2021-02-03 03:49:33.436', '2021-02-03 03:49:33.436', NULL);

-- Table Music --
INSERT INTO public.music(title, artist, position, created_at, updated_at)
VALUES
('Independent Women Part I', E'Destiny\'s Child', 1, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Maria, Maria', 'Santana', 2, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('I Knew I Loved You', 'Savage Garden', 3, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Music', 'Madonna', 4, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Come On Over Baby (All I Want Is You)', 'Aguilera, Christina', 5, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
(E'Doesn\'t Really Matter', 'Janet', 6, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Say My Name', E'Destiny\'s Child', 7, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Be With You', 'Iglesias, Enrique', 8, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Incomplete', 'Sisqo', 9, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Amazed', 'Lonestar', 10, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
(E'It\'s Gonna Be Me', E'N\'Sync', 11, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('What A Girl Wants', 'Aguilera, Christina', 12, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Everything You Want', 'Vertical Horizon', 13, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('With Arms Wide Open', 'Creed', 14, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Try Again', 'Aaliyah', 15, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Bent', 'matchbox twenty', 16, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Thank God I Found You', 'Carey, Mariah', 17, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Breathe', 'Hill, Faith', 18, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Case Of The Ex (Whatcha Gonna Do)', 'Mya', 19, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('He Loves U Not', 'Dream', 20, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
(E'He Wasn\'t Man Enough', 'Braxton, Toni', 21, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('You Sang To Me', 'Anthony, Marc', 22, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Give Me Just One Night (Una Noche)', '98¡', 23, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
(E'Jumpin\' Jumpin\'', E'Destiny\'s Child', 24, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Kryptonite', '3 Doors Down', 25, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Thong Song', 'Sisqo', 26, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('I Turn To You', 'Aguilera, Christina', 27, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('My Love Is Your Love', 'Houston, Whitney', 28, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('I Wanna Know', 'Joe', 29, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Get It On.. Tonite', 'Jordan, Montell', 30, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Bye Bye Bye', E'N\'Sync', 31, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('The Real Slim Shady', 'Eminem', 32, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Most Girls', 'Pink', 33, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Gotta Tell You', 'Mumba, Samantha', 34, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Bring It All To Me', 'Blaque', 35, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Hot Boyz', 'Elliott, Missy "Misdemeanor"', 36, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('I Try', 'Gray, Macy', 37, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('This I Promise You', E'N\'Sync', 38, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('No More', 'Ruff Endz', 39, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Show Me The Meaning Of Being Lonely', 'Backstreet Boys, The', 40, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Absolutely (Story Of A Girl)', 'Nine Days', 41, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
(E'That\'s The Way It Is', 'Dion, Celine', 42, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('All The Small Things', 'Blink-182', 43, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Blue (Da Ba Dee)', 'Eiffel 65', 44, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Bag Lady', 'Badu, Erkyah', 45, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('There U Go', 'Pink', 46, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Higher', 'Creed', 47, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('(Hot S**t) Country Grammar', 'Nelly', 48, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Wifey', 'Next', 49, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z'),
('Auld Lang Syne (The Millenium Mix)', 'Kenny G', 50, '2021-02-03T06:00:00Z', '2021-02-03T06:00:00Z');

-- Table User Fav --
INSERT INTO public.user_fav
(user_id, music_id, created_at, updated_at, deleted_at)
VALUES
(1, 3, '2021-02-03 00:00:00.000', '2021-02-03 00:00:00.000', NULL),
(1, 20, '2021-02-03 00:00:00.000', '2021-02-03 00:00:00.000', NULL);