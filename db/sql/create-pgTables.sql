CREATE TABLE userinfo
(
    uid serial NOT NULL,
    username character varying(100) NOT NULL,
    department character varying(500) NOT NULL,
    Created date,
    CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);
CREATE TABLE userdetail
(
    uid integer,
    intro character varying(100),
    profile character varying(100)
)
WITH(OIDS=FALSE);

INSERT INTO userinfo
(username, department, created)
VALUES
('Lily Lonika', '研发部门', '2020-12-09'),
('Gerry Mulligan', '测试部门', '2020-07-13');