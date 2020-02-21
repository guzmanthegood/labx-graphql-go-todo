CREATE TABLE public.user (
    id SERIAL PRIMARY KEY,
    name character varying(100)
);

ALTER TABLE public.user
    ADD CONSTRAINT client_name_key UNIQUE (name);

CREATE TABLE public.todo (
     id SERIAL PRIMARY KEY,
     description character varying(255),
     status character varying(20),
     user_id integer REFERENCES public.user(id) ON DELETE CASCADE
);
