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
    id SERIAL PRIMARY KEY,
    instance_id INTEGER REFERENCES instances(id),
    question_id INTEGER REFERENCES questions(id),
    answer_id INTEGER REFERENCES answers(id),
    main_text VARCHAR(2048),
    -- default
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

select p.id, p.name, pi.id, pi.poll_id
from polls p
join instances pi
on p.id = pi.poll_id;

select 
    i.owner_cn as instance_owner_cn,
    i.id as instance_id,
    p.name as poll_name,
    json_agg(json_build_object(
        'question_main_text', q.main_text,
        'question_position', q.position,
        'answers', (SELECT json_agg(json_build_object('id', id, 'main_text', main_text)) FROM answers WHERE question_id = q.id)
    ))
from
    instances i
join
    polls p
on 
    i.poll_id = p.id
left outer join
    questions q
on 
    q.poll_id = i.poll_id
GROUP BY i.id, p.name, i.owner_cn
;