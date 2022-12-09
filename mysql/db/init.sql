DROP SCHEMA IF EXISTS odoro-tree;
CREATE SCHEMA odoro-tree;
USE odoro-tree;

DROP TABLE IF EXISTS odoro-tree;

CREATE TABLE trees
(
    tree_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name    VARCHAR(10)
);

CREATE TABLE branches
(
    branch_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    parent_tree_id INT NOT NULL
);

CREATE TABLE leaves
(
    leaf_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    front_content VARCHAR(256) NOT NULL,
    back_content VARCHAR(256) NOT NULL,
    rating INT NOT NULL,
    parent_branch_id INT NOT NULL
);
