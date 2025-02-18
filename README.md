# AUTH APP

---

### This is an auth web app that is created using Golang as its backend and Next.js as its frontend

#### How to run the project:

1. Create .env for both the backend and frontend using .env.example as an example
2. Create a postgres database for the app (if using the .env.example, the database name is auth_db)
3. Seed the database:

```sql
  DDL

  CREATE TABLE users (
	  id BIGSERIAL PRIMARY KEY,
	  username VARCHAR NOT NULL,
	  email VARCHAR NOT NULL,
	  password VARCHAR NOT NULL,
	  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	  edited_at TIMESTAMP NOT NULL DEFAULT NOW(),
	  deleted_at TIMESTAMP NULL
  );
```

```sql
  DML

  INSERT INTO users(username, email, "password")
  VALUES
  ('Test', 'testing@gmail.com', '$2a$10$XHHCQ.hYfq/36X2WgUUyx.S4TfCZzDv4YNRNX4KJfu.eDecB53Wxe');

  -- User info:
  -- Username : Test
  -- Email: testing@gmail.com
  -- Password: Password123#
```

4. Finally, run the `backend` and `frontend` 
(for backend, main.go is located at /backend/cmd/app/main.go)
---
### Demo video
https://drive.google.com/file/d/1isCNETvHUqwsQ_k1uqYIWdaptIzI1jO1/view?usp=sharing
