create table users (
	user_id serial not null primary key,
	username character varying(32) not null,
	password character varying(60) not null,
	email character varying(128) not null,
	created_at timestamp with time zone default current_timestamp
);
