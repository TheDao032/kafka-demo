CREATE TABLE IF NOT EXISTS public.factinternetsales_streaming
(
    productkey bigint,
    customerkey bigint,
    salesterritorykey bigint,
    salesordernumber text COLLATE pg_catalog."default",
    totalproductcost double precision,
    salesamount double precision
) TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.factinternetsales_streaming
    OWNER to postgres;

-- INSERT INTO public.factinternetsales_streaming ("productkey", "customerkey", "salesterritorykey", "salesordernumber", "totalproductcost", "salesamount") VALUES (1, 1, 2, 'test', 0.1, 0.2);