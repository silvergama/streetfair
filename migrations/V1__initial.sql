-- Drop table

-- DROP TABLE public.free_fairs

CREATE TABLE free_fairs (
    id integer PRIMARY KEY,
    "long" character(12) NULL,
    lat character(12) NULL,
    setcens integer NULL,
    areap integer NULL,
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
