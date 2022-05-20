CREATE EXTENSION unaccent;
CREATE EXTENSION pg_trgm;

-- Drop table

-- DROP TABLE public.streetfair

CREATE TABLE streetfair (
    id integer PRIMARY KEY,
    long numeric NULL,
    lat numeric NULL,
    setcens bigint NULL,
    areap bigint NULL,
    coddist integer NULL,
    distrito varchar2(55) NULL,
    codsubpref integer NULL,
    subprefe varchar2(150),
    regiao5 varchar2(12),
    regiao8 varchar2(12),
    nome_feira varchar2(175) NULL,
    registro varchar2(10),
    logradouro varchar2(175) NULL,
    numero varchar2(12) NULL,
    bairro varchar2(72) NULL,
    referencia varchar2(150)
);
