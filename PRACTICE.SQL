sudo -i -u postgres     --seting-up postgres
psql                    --start postgresql

CREATE USER name_of_user WITH PASSWORD 'the_password'; --create new user with password

ALTER ROLE name_of_user WITH SUPERUSER;                --set root permissions to new user

SELECT CURRENT_ROLE;                                   --show current user

SET ROLE name_of_user;                                 --swith to another user

CREATE DATABASE database_name WITH OWNER name_of_user  --create new database with selected user

CREATE TABLE table_name(
    ID SERIAL PRIMARY KEY,                              --serial gives number by default from 1 to ..., Primary key needs to connect with another table
    NAME TEXT,                                          --it is like varchar(255)
    AGE INT,                                            --how to input integer
    JOB VARCHAR(30),                                    --you have an access to input the string up to 30 signs
    JOIN_DATE TIMESTAMP DEFAULT CURRENT_TIMESTAMP       --writes current_time by default, if you do not input
);

INSERT INTO table_name(NAME, AGE, JOB) VALUES           --how to insert information to table
    ('Lionel Messi', 37, 'Football player'),
    ('Burak Dakkak', 31, 'Actor'),
    ('Gareth Bale', 34, 'Golf player'),
    ('Indila', 35, 'Singer'),
    ('Ronnie Coleman', 56, 'Bodybuilder');

SELECT ID, NAME || ' ' || AGE as GENERAL FROM table_name;       -- || ' ' || helps to combine multiple columns

SELECT ID, NAME "ISMLAR" FROM table_name;                       -- it is also one way to print column with another name, name --> ismlar

CREATE TABLE students(
    ID SERIAL PRIMARY KEY,
    NAME VARCHAR(30),
    COUNTRY VARCHAR(30),
    AGE INT,
    COURSE INT,
    FACULTY VARCHAR(30),
    IS_SCHOLARSHIP BOOLEAN
);