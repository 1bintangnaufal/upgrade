-- Table: public.tb_project

-- DROP TABLE IF EXISTS public.tb_project;

CREATE TABLE IF NOT EXISTS public.tb_project
(
    id integer NOT NULL DEFAULT nextval('tb_project_id_seq'::regclass),
    author_id integer,
    project_title character varying(30) COLLATE pg_catalog."default" NOT NULL,
    start_date date NOT NULL,
    finish_date date NOT NULL,
    description character varying(1024) COLLATE pg_catalog."default",
    toggle_a boolean NOT NULL,
    toggle_b boolean NOT NULL,
    toggle_c boolean NOT NULL,
    toggle_d boolean NOT NULL,
    image character varying COLLATE pg_catalog."default",
    CONSTRAINT tb_project_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tb_project
    OWNER to postgres;