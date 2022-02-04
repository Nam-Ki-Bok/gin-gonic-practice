👑 Gin framework 👑

## Database initialization (MariaDB)

```mariadb
CREATE DATABASE Database_Name
```

```mariadb
CREATE TABLE Database_Name.users (
	idx int(10) unsigned auto_increment NOT NULL,
	email varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
	password varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
	name varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
	CONSTRAINT `PRIMARY` PRIMARY KEY (idx),
	CONSTRAINT users_un UNIQUE KEY (email)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb3
COLLATE=utf8mb3_general_ci
COMMENT='';
```

```mariadb
CREATE TABLE Database_Name.posts (
	idx int(10) unsigned auto_increment NOT NULL,
	email varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
	title varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
	content varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL NULL,
	name varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
	CONSTRAINT `PRIMARY` PRIMARY KEY (idx)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb3
COLLATE=utf8mb3_general_ci
COMMENT='';
CREATE INDEX posts_FK USING BTREE ON Database_Name.posts (email);
```
