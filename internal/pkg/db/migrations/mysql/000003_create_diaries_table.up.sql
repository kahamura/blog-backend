CREATE TABLE IF NOT EXISTS diaries(
    id INT NOT NULL UNIQUE AUTO_INCREMENT,
    title VARCHAR (255),
    content TEXT,
    userID INT,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    FOREIGN KEY (userID) REFERENCES users(id) ,
    PRIMARY KEY (id)
)