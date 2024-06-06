-- SEQUENCE: public.city_id_seq

-- DROP SEQUENCE IF EXISTS public.city_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.city_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.city_id_seq
    OWNED BY public.city.id;

ALTER SEQUENCE public.city_id_seq
    OWNER TO postgres;


-- SEQUENCE: public.continent_id_seq

-- DROP SEQUENCE IF EXISTS public.continent_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.continent_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.continent_id_seq
    OWNED BY public.continent.id;

ALTER SEQUENCE public.continent_id_seq
    OWNER TO postgres;


-- SEQUENCE: public.country_id_seq

-- DROP SEQUENCE IF EXISTS public.country_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.country_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.country_id_seq
    OWNER TO postgres;


-- Table: public.continent

-- DROP TABLE IF EXISTS public.continent;

CREATE TABLE IF NOT EXISTS public.continent
(
    id integer NOT NULL DEFAULT nextval('continent_id_seq'::regclass),
    name character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT continent_pkey PRIMARY KEY (id),
    CONSTRAINT continent_name_index UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.continent
    OWNER to postgres;


-- Table: public.country

-- DROP TABLE IF EXISTS public.country;

CREATE TABLE IF NOT EXISTS public.country
(
    id integer NOT NULL DEFAULT nextval('country_id_seq'::regclass),
    continent_id integer,
    name character varying(255) COLLATE pg_catalog."default",
    population numeric,
    area numeric,
    CONSTRAINT country_pkey PRIMARY KEY (id),
    CONSTRAINT country_name_index UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.country
    OWNER to postgres;


-- Table: public.city

-- DROP TABLE IF EXISTS public.city;

CREATE TABLE IF NOT EXISTS public.city
(
    id integer NOT NULL DEFAULT nextval('city_id_seq'::regclass),
    country_id integer,
    name character varying(255) COLLATE pg_catalog."default",
    population numeric(10,0),
    area numeric(10,0),
    is_capital boolean,
    CONSTRAINT city_pkey PRIMARY KEY (id),
    CONSTRAINT city_name_index UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.city
    OWNER to postgres;

