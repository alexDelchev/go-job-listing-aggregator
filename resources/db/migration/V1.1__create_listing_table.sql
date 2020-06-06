CREATE SEQUENCE listing_listing_id_seq
    INCREMENT BY 1
    NO MAXVALUE
    NO MINVALUE
    CACHE 1;

ALTER TABLE public.listing_listing_id_seq OWNER TO "go_job_listing_aggregator";

CREATE TABLE listing(
    listing_id INTEGER DEFAULT NEXTVAL('listing_listing_id_seq'::regclass) NOT NULL,
    external_id CHARACTER VARYING(500) NOT NULL,
    link CHARACTER VARYING(500) NOT NULL,
    name CHARACTER VARYING(500) NOT NULL,
    work_schedule CHARACTER VARYING(500) NOT NULL,
    company CHARACTER VARYING(500) NOT NULL,
    location CHARACTER VARYING(500) NOT NULL,
    posting_date CHARACTER VARYING(500) NOT NULL,
    description TEXT NOT NULL,
    keywords text[] NOT NULL,
    query_id INTEGER NOT NULL,
    source_name CHARACTER VARYING(500) NOT NULL
);

ALTER TABLE public.listing OWNER TO "go_job_listing_aggregator";

CREATE INDEX idx_listing_external_id ON listing USING btree (external_id);

CREATE INDEX idx_listing_query_id ON listing USING btree (external_id);

ALTER TABLE ONLY listing ADD CONSTRAINT listing_pkey PRIMARY KEY (listing_id);

ALTER TABLE ONLY listing ADD CONSTRAINT listing_query_fkey FOREIGN KEY (query_id) REFERENCES query(query_id);

