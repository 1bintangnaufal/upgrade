-- Table: public.tb_users

-- DROP TABLE IF EXISTS public.tb_users;

CREATE TABLE IF NOT EXISTS public.tb_users
(
    id integer NOT NULL DEFAULT nextval('tb_users_id_seq'::regclass),
    name character varying(128) COLLATE pg_catalog."default" NOT NULL,
    email character varying(128) COLLATE pg_catalog."default" NOT NULL,
    password character varying(128) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT tb_users_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tb_users
    OWNER to postgres;