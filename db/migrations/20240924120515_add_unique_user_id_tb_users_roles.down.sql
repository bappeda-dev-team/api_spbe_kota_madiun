
ALTER TABLE users_roles 
DROP FOREIGN KEY fk_user_id;

ALTER TABLE users_roles 
DROP INDEX unique_user_id;

ALTER TABLE users_roles 
ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id);