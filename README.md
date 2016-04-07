Online Tests for Interviews
===========================

DATABASE NAME

Development - online_test_dev
Test - online_test_test


# Schema definitions

### Create users table
```
CREATE TABLE USERS (
id SERIAL,
first_name varchar(100),
last_name varchar(100),
email varchar(100),
college varchar(200),
branch varchar(100),
password varchar(100),
password_confirmation varchar(100),
phone_number varchar(100),
role varchar(50),
year_of_passing varchar(4),
PRIMARY KEY(id));
```

### Create sessions table
```
CREATE TABLE sessions(
id int,
start_time timestamptz,
end_time timestamptz,
user_id int,
CONSTRAINT session_id_key FOREIGN KEY (user_id)
REFERENCES users (id),
auth_token varchar(320),
UNIQUE (auth_token));
```

### Create Section table
```
CREATE TABLE SECTIONS (
id int,
name varchar(100),
PRIMARY KEY(id));
```

### Create Questions table
```
CREATE TABLE QUESTIONS (
id int,
title text,
option_1 varchar(100),
option_2 varchar(100),
option_3 varchar(100),
option_4 varchar(100),
answer varchar(100),
section_id int,
CONSTRAINT section_id_key FOREIGN KEY(section_id)
REFERENCES sections(id),
PRIMARY KEY(id));
```

### Create Result table

```
CREATE TABLE RESULTS (
id SERIAL,
user_id int,
CONSTRAINT session_section_key FOREIGN KEY(user_id)
REFERENCES users(id),
first_name varchar(100),
last_name varchar(100),
email varchar(100),
section_1 int,
section_2 int,
section_3 int,
total_score int
);
```

# Inputs data and fields required for running API


### Create User

URL - http://localhost:3000/sign_up

Method POST

Data has to be sent in raw format
```
{"firstname":"steve","lastname":"jobs","email":"steve@example.com","password":"password","password_confirmation":"password","Branch":"Information Science","Year_of_passing":"2014","Phone_number":"9916854300","Auth_token":"039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2"}
```
The user registers for online exam . A session is created as soon as he signs up.

#### JSON Response

```
{
	"Success": "true",
	"Message": "User created Successfully!",
	"User": {
		"Id": 1,
		"Firstname": "steve",
		"Lastname": "jobs",
		"Email": "steve@example.com",
		"Password": "password",
		"Password_confirmation": "password",
		"Branch": "Information Science",
		"Year_of_passing": "2014",
		"Phone_number": "9916854300",
		"Auth_token": "039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2"
	},
	"Session": {
		"SessionId": 1,
		"StartTime": "2016-03-04T00:59:49.784724111+05:30"
	}
}
```
---
