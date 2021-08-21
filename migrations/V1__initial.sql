CREATE EXTENSION unaccent;
CREATE EXTENSION pg_trgm;

-- Drop table

-- DROP TABLE public.free_fair

CREATE TABLE free_fair (
    id integer PRIMARY KEY,
    long numeric NULL,
    lat numeric NULL,
    setcens bigint NULL,
    areap bigint NULL,
    coddist integer NULL,
    distrito varchar(55) NULL,
    codsubpref integer NULL,
    subprefe varchar(150),
    regiao5 character(15),
    regiao8 character(15),
    nome_feira varchar(175) NULL,
    registro varchar(8),
    logradouro varchar(175) NULL,
    numero character(20) NULL,
    bairro varchar(72) NULL,
    referencia varchar(150)
);
