-- CREATE DATABASE Tasks_DB

/**
DROP DATABASE Tasks_DB

DROP TABLE dbo.Category
DROP TABLE dbo.Priority
DROP TABLE dbo.User
DROP TABLE dbo.Task
*/

-- Crear Tabla Category
CREATE TABLE "Category" (
	"Code" INTEGER,
	"Description" TEXT NOT NULL UNIQUE,
	PRIMARY KEY("Code" AUTOINCREMENT)
);

CREATE TABLE "Priority" (
	"Code" INTEGER,
	"Description" TEXT NOT NULL UNIQUE,
	PRIMARY KEY("Code" AUTOINCREMENT)
);

CREATE TABLE "User" (
	"Id"	TEXT,
	"Name"	TEXT UNIQUE,
	"Email"	TEXT NOT NULL UNIQUE,
	"Password"	TEXT NOT NULL,
	"CreatedAt"	TEXT NOT NULL,
	"LastSession"	TEXT NOT NULL,
	PRIMARY KEY("Id")
);

CREATE TABLE "Task" (
	"Id"	TEXT,
	"Title"	TEXT NOT NULL,
	"Description"	TEXT NOT NULL,
	"CategoryCode"	INTEGER NOT NULL,
	"PriorityCode"	INTEGER NOT NULL,
	"CreateAt"	TEXT NOT NULL,
	"Done"	INTEGER NOT NULL DEFAULT 1,
	"UserId"	TEXT NOT NULL,
	FOREIGN KEY("CategoryCode") REFERENCES "Category"("Code"),
	FOREIGN KEY("PriorityCode") REFERENCES "Priority"("Code"),
	PRIMARY KEY("Id")
);
