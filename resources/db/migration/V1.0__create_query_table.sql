CREATE SEQUENCE query_query_id_seq
    INCREMENT BY 1
    NO MAXVALUE
    NO MINVALUE
    CACHE 1;

ALTER TABLE public.query_query_id_seq OWNER TO "go_job_listing_aggregator";

CREATE TABLE query(
    query_id INTEGER DEFAULT NEXTVAL('query_query_id_seq'::regclass) NOT NULL,
    keywords text[] NOT NULL,
    location CHARACTER VARYING(500) NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    creation_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL
);

ALTER TABLE public.query OWNER TO "go_job_listing_aggregator";

ALTER TABLE ONLY query ADD CONSTRAINT query_pkey PRIMARY KEY (query_id);