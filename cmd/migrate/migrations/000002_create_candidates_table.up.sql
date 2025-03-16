CREATE TABLE IF NOT EXISTS candidates(
    id bigserial PRIMARY KEY,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    email varchar(255) NOT NULL, --has to be unique, but for research purposes i accept duplicates
    created TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);