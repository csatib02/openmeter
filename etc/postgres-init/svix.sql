CREATE USER svix WITH PASSWORD 'svix';

CREATE DATABASE svix;

GRANT ALL PRIVILEGES ON DATABASE svix TO svix;

ALTER DATABASE svix OWNER TO svix;