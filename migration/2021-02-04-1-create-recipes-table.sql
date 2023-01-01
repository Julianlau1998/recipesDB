CREATE TABLE recipes(
    uuid varchar(36) NOT NULL,
    title varchar(255),
    description TEXT,
    duration varchar(30),
    category_ID varchar(36) NULL,
    image varchar(255) NOT NULL,
    country varchar(255) NOT NULL,
    is_vegetarian BOOLEAN DEFAULT FALSE,
    is_vegan BOOLEAN DEFAULT FALSE,
    createdDate TIMESTAMP,
    PRIMARY KEY (uuid),
    FOREIGN KEY (category_ID) REFERENCES categories(uuid) ON DELETE SET NULL
)
