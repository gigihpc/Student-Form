CREATE TABLE public.users
(
  id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  name text,
  email text,
  password text,
  country text,
  CONSTRAINT users_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.users
  OWNER TO postgres;


CREATE TABLE public.mhsw
(
  id integer NOT NULL DEFAULT nextval('mhsw_id_seq'::regclass),
  name text,
  address text,
  old text,
  CONSTRAINT mhsw_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.mhsw
  OWNER TO postgres;
