CREATE TABLE users (
    id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role TEXT NOT NULL, -- admin, user, etc
    created_at TIMESTAMP DEFAULT now()
);

SELECT * FROM users
-- INSERT INTO public.users(
-- 	id, email, password, role, created_at)
-- 	VALUES ('f07378d0-3fd4-42a0-b8fd-dde73b0601a6', 'travelxon@gmail.com', 'Tester@123', 'admin', now());