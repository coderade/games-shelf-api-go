--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: genres; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.genres
(
    id         integer NOT NULL,
    genre_name character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Name: genres_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.genres_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: genres_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.genres_id_seq OWNED BY public.genres.id;

--
-- Name: platforms; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.platforms
(
    id            integer NOT NULL,
    platform_name character varying,
    generation    character varying,
    created_at    timestamp without time zone,
    updated_at    timestamp without time zone
);


--
-- Name: platforms_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.platforms_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: platforms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.platforms_id_seq OWNED BY public.platforms.id;


--
-- Name: games; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.games
(
    id          integer NOT NULL,
    title       character varying,
    description text,
    year        integer,
    publisher   character varying,
    rawg_id      integer,
    rating integer,
    created_at  timestamp without time zone,
    updated_at  timestamp without time zone
);

--
-- Name: platforms_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.games_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: platforms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.games_id_seq OWNED BY public.games.id;

--
-- Name: games_genres; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.games_genres
(
    id         integer NOT NULL,
    game_id    integer,
    genre_id   integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: games_genres_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.games_genres_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: games_genres_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.games_genres_id_seq OWNED BY public.games_genres.id;


--
-- Name: games_platforms; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.games_platforms
(
    id          integer NOT NULL,
    game_id     integer,
    platform_id integer,
    created_at  timestamp without time zone,
    updated_at  timestamp without time zone
);


--
-- Name: games_platforms_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.games_platforms_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: games_platforms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.games_platforms_id_seq OWNED BY public.games_platforms.id;




--
-- Name: genres id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.genres
    ALTER COLUMN id SET DEFAULT nextval('public.genres_id_seq'::regclass);


--
-- Name: platforms id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.platforms
    ALTER COLUMN id SET DEFAULT nextval('public.platforms_id_seq'::regclass);

--
-- Name: games id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games
    ALTER COLUMN id SET DEFAULT nextval('public.games_id_seq'::regclass);


--
-- Name: games_genres id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games_genres
    ALTER COLUMN id SET DEFAULT nextval('public.games_genres_id_seq'::regclass);


--
-- Data for Name: genres; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.genres (id, genre_name, created_at, updated_at)
VALUES (1, 'Adventure', '2022-04-04 00:00:00', '2022-04-04 00:00:00'),
       (2, 'Sports', '2022-04-04 00:00:00', '2022-04-04 00:00:00'),
       (3, 'Action', '2022-04-04 00:00:00', '2022-04-04 00:00:00'),
       (4, 'FPS', '2022-04-04 00:00:00', '2022-04-04 00:00:00'),
       (5, 'RPG', '2022-04-04 00:00:00', '2022-04-04 00:00:00'),
       (6, 'Racing', '2022-04-04 00:00:00', '2022-04-04 00:00:00'),
       (7, 'Fighting', '2022-04-04 00:00:00', '2022-04-04 00:00:00'),
       (8, 'Platform', '2022-04-04 00:00:00', '2022-04-04 00:00:00');

--
-- Data for Name: genres; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.platforms (id, platform_name, generation,created_at , updated_at)
VALUES (1, 'Nintendo 64', '2022-04-04 00:00:00', 5, '2022-04-04 00:00:00'),
       (2, 'Super Nintendo Entertainment System',4, '2022-04-04 00:00:00', '2022-04-04 00:00:00'),
       (3, 'Playstation 1', '2022-04-04 00:00:00',5, '2022-04-04 00:00:00'),
       (4, 'Sega Genesis', '2022-04-04 00:00:00', 4, '2022-04-04 00:00:00');

--
-- Data for Name: games; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.games (id, title, description, year, publisher, rawg_id, rating, created_at, updated_at)
VALUES (1, 'The legend of Zelda: Ocarina of Time', 'The Legend of Zelda: Ocarina of Time is a fantasy action-adventure ' ||
                                                   'game set in an expansive environment. The player controls series ' ||
                                                   'protagonist Link from a third-person perspective in a three-dimensional world',
        1998, 'Nintendo', 25097, 99, '2022-04-04 00:00:00', '2022-04-04 00:00:00'),
       (2, 'Donkey Kong Country',
        'Donkey Kong Country is a side-scrolling platform game in which the player must complete 40 levels to recover ' ||
        'the Kongs banana hoard, which has been stolen by the crocodilian Kremlings',1994, 'Nintendo', 85, 90,
        '2022-04-04 00:00:00', '2022-04-04 00:00:00'),
       (3, 'Tony Hawks Pro Skater 2', 'Tony Hawks Pro Skater 2 is a skateboarding video game developed by Neversoft ' ||
                                      'and published by Activision.', 2000, 'Activision', 57944, 90 , '2022-04-04 00:00:00',
        '2022-04-04 00:00:00'),
       (4, 'Sonic the Hedgehog', 'Sonic the Hedgehog is a platform video game developed by Sonic Team and published ' ||
                                 'by Sega for the Sega Genesis home video game console. The first game in the Sonic ' ||
                                 'the Hedgehog franchise',  1991, 'Sega', 53551, 92, '2022-04-04 00:00:00', '2022-04-04 00:00:00');

--
-- Name: genres genres_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

--
-- Data for Name: games_genres; Type: TABLE DATA; Schema: public; Owner: -
--
INSERT INTO games_genres VALUES (1, 1, 1, NOW(), NOW() );
INSERT INTO games_genres VALUES (2, 1, 3, NOW(), NOW() );


--
-- Data for Name: games_genres; Type: TABLE DATA; Schema: public; Owner: -
--
INSERT INTO games_platforms VALUES (1, 1, 1, NOW(), NOW() );

ALTER TABLE ONLY public.genres
    ADD CONSTRAINT genres_pkey PRIMARY KEY (id);

--
-- Name: genres platforms_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.platforms
    ADD CONSTRAINT platforms_pkey PRIMARY KEY (id);


--
-- Name: games_genres games_genres_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games_genres
    ADD CONSTRAINT games_genres_pkey PRIMARY KEY (id);

--
-- Name: games_platforms games_platforms_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games_platforms
    ADD CONSTRAINT games_platforms_pkey PRIMARY KEY (id);

--
-- Name: games games_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games
    ADD CONSTRAINT games_pkey PRIMARY KEY (id);


--
-- Name: games_genres fk_game_genres_genre_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games_genres
    ADD CONSTRAINT fk_game_genres_genre_id FOREIGN KEY (genre_id) REFERENCES public.genres (id);

--
-- Name: games_genres fk_game_genres_game_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games_genres
    ADD CONSTRAINT fk_game_genres_game_id FOREIGN KEY (game_id) REFERENCES public.games (id);

--
-- Name: games_genres fk_game_platforms_platform_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games_platforms
    ADD CONSTRAINT fk_game_platforms_platform_id FOREIGN KEY (platform_id) REFERENCES public.platforms (id);

--
-- Name: games_genres fk_game_platforms_game_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games_platforms
    ADD CONSTRAINT fk_game_platforms_game_id FOREIGN KEY (game_id) REFERENCES public.games (id);


