CREATE DATABASE Tasks_DB
/**
DROP DATABASE Tasks_DB

DROP TABLE dbo.Category
DROP TABLE dbo.Priority
DROP TABLE dbo.User
DROP TABLE dbo.Task
*/

-- Crear Tabla Category
CREATE TABLE "Category" (
	"Code"	INTEGER,
	"Description"	INTEGER NOT NULL UNIQUE,
	PRIMARY KEY("Code" AUTOINCREMENT)
);

-- Crear Tabla Priority
CREATE TABLE [dbo].[Priority](
	[Code] [tinyint] NOT NULL,
	[Description] [nvarchar](25) NOT NULL,
 CONSTRAINT [PK_Priority] PRIMARY KEY CLUSTERED 
(
	[Code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO

-- Crear Tabla User
CREATE TABLE [dbo].[User](
	[Id] [uniqueidentifier] NOT NULL,
	[Name] [nvarchar](30) NOT NULL,
	[Email] [nvarchar](100) NOT NULL,
	[Password] [nvarchar](32) NOT NULL,
	[CreateAt] [datetime] NOT NULL,
	[LastSession] [datetime] NOT NULL,
 CONSTRAINT [PK_User] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO

-- Crear Tabla Task
CREATE TABLE [dbo].[Task](
	[Id] [uniqueidentifier] NOT NULL,
	[Title] [nvarchar](50) NOT NULL,
	[Description] [nvarchar](max) NOT NULL,
	[CategoryCode] [tinyint] NOT NULL,
	[PriorityCode] [tinyint] NOT NULL,
	[CreateAt] [datetime] NOT NULL,
	[Done] [bit] NOT NULL,
	[UserId] [uniqueidentifier] NOT NULL,
 CONSTRAINT [PK_Task] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO

ALTER TABLE [dbo].[Task]  WITH CHECK ADD  CONSTRAINT [FK_Task_Category] FOREIGN KEY([CategoryCode])
REFERENCES [dbo].[Category] ([Code])
GO

ALTER TABLE [dbo].[Task] CHECK CONSTRAINT [FK_Task_Category]
GO

ALTER TABLE [dbo].[Task]  WITH CHECK ADD  CONSTRAINT [FK_Task_Priority] FOREIGN KEY([PriorityCode])
REFERENCES [dbo].[Priority] ([Code])
GO

ALTER TABLE [dbo].[Task] CHECK CONSTRAINT [FK_Task_Priority]
GO

ALTER TABLE [dbo].[Task]  WITH CHECK ADD  CONSTRAINT [FK_Task_User] FOREIGN KEY([UserId])
REFERENCES [dbo].[User] ([Id])
GO

ALTER TABLE [dbo].[Task] CHECK CONSTRAINT [FK_Task_User]
GO
