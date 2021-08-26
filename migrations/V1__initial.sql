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
    regiao5 varchar(12),
    regiao8 varchar(12),
    nome_feira varchar(175) NULL,
    registro varchar(10),
    logradouro varchar(175) NULL,
    numero varchar(12) NULL,
    bairro varchar(72) NULL,
    referencia varchar(150)
);
