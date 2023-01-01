CREATE TABLE recipe_ingredients(
    recipe_ID varchar(36) NOT NULL,
    ingredient_ID varchar(36) NOT NULL,
    FOREIGN KEY (recipe_ID) REFERENCES recipes(uuid) ON DELETE CASCADE,
    FOREIGN KEY (ingredient_ID) REFERENCES ingredients(uuid) ON DELETE CASCADE
)
