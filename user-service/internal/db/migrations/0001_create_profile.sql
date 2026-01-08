CREATE TABLE local_profile (
    id uuid NOT NULL,
    name varchar NOT NULL,
    gender gender_enum NOT NULL,
    email varchar,
    status status_enum NOT NULL DEFAULT '0',
    role_id uuid,
    CONSTRAINT local_profile_pk PRIMARY KEY (id)
);