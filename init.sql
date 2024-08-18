--- TABLES
CREATE TABLE candidate (
	id VARCHAR PRIMARY KEY,
	name VARCHAR NOT NULL,
);

CREATE TABLE vote (
	id VARCHAR PRIMARY KEY,
	candidate_id VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_candidate_vote_id
		FOREIGN KEY (candidate_id) REFERENCES candidate(id)
);

--- INDEX
CREATE UNIQUE INDEX idx_candidate_id ON candidate (id);
CREATE UNIQUE INDEX idx_vote_id ON vote (id);

--- SEED
INSERT INTO candidate (id, name)
VALUES 
	('clzc1pqd0000008mnfmkq9r50', 'Bruno Uemura'),
    ('cjd7n3v6g0001gq9g7j2m3pbk', 'John Doe'),
    ('cjd7n3v6g0002gq9g7j2m3pbk', 'Jane Smith'),
    ('cjd7n3v6g0003gq9g7j2m3pbk', 'Alice Johnson'),
    ('cjd7n3v6g0004gq9g7j2m3pbk', 'Bob Brown'),
    ('cjd7n3v6g0005gq9g7j2m3pbk', 'Charlie Davis');