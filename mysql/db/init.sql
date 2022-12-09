DROP SCHEMA IF EXISTS odoro_tree;
CREATE SCHEMA odoro_tree;
USE odoro_tree;

DROP TABLE IF EXISTS trees;
DROP TABLE IF EXISTS branches;
DROP TABLE IF EXISTS leaves;

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
