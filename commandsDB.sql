CREATE TABLE movies (
    id SERIAL,
    movieID varchar(50) NOT NULL,
    movieName varchar(50) NOT NULL,
    PRIMARY KEY (id)
)

INSERT INTO movies (
    movieID,
    movieName
)
VALUES
    ('1', 'Matrix Revolutions'),
    ('2 ', 'Matrix Reloaded'),
    ('3', 'Matrix');