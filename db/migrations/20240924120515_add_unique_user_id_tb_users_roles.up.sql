ALTER TABLE users_roles 
ADD CONSTRAINT unique_user_id UNIQUE (user_id);
