DROP TABLE IF EXISTS public.user_fav;
DROP TABLE IF EXISTS public.music;
DROP TABLE IF EXISTS public.user;

CREATE TABLE public.user (
	id bigserial NOT NULL,
	user_name varchar(50) NOT NULL UNIQUE,
	full_name varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    gender varchar(2) NULL,
    hobby varchar(255) NULL,
    address text NULL,
	created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamp NULL,
	CONSTRAINT user_pkey PRIMARY KEY (id)
);

CREATE TABLE public.music (
    id bigserial NOT NULL,
	title varchar(255) NOT NULL,
	artist varchar(255) NOT NULL,
    position int2 NOT NULL,
	created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamp NULL,
	CONSTRAINT music_pkey PRIMARY KEY (id)
);

CREATE TABLE public.user_fav (
    id bigserial NOT NULL,
    user_id int8 NOT NULL,
    music_id int8 NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamp NULL,
	CONSTRAINT user_fav_pkey PRIMARY KEY (id)
);

ALTER TABLE public.user_fav ADD CONSTRAINT user_fav_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.user(id);
ALTER TABLE public.user_fav ADD CONSTRAINT user_fav_music_id_foreign FOREIGN KEY (music_id) REFERENCES public.music(id);