drop table replies;
drop table instances;
drop table answers;
drop table questions;
drop table polls;

CREATE TABLE IF NOT EXISTS polls (
    id SERIAL PRIMARY KEY,
    name VARCHAR(512) NOT NULL,
    owner_cn VARCHAR(64) NOT NULL,
    -- default
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS questions (
    id SERIAL PRIMARY KEY,
    main_text VARCHAR(2048),
    position INTEGER NOT NULL,
    poll_id INTEGER REFERENCES polls(id),
    -- default
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS answers (
    id SERIAL PRIMARY KEY,
    question_id INTEGER REFERENCES questions(id),
    main_text VARCHAR(2048),
    -- default
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS instances (
    id SERIAL PRIMARY KEY,
    poll_id INTEGER REFERENCES polls(id),
    owner_cn VARCHAR(64) NOT NULL,
    -- default
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS replies (
    instance_id INTEGER REFERENCES instances(id),
    question_id INTEGER REFERENCES questions(id),
    answer_id INTEGER REFERENCES answers(id),
    main_text VARCHAR(2048),
    PRIMARY KEY (instance_id, question_id, answer_id),
    -- default
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

select p.id, p.name, pi.id, pi.poll_id
from polls p
join instances pi
on p.id = pi.poll_id;