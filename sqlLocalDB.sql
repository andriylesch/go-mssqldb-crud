USE [LocalDb]
GO
/****** Object:  Table [dbo].[Users]    Script Date: 27/12/2016 11:47:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Users](
	[Id] [int] IDENTITY(1,1) NOT NULL,
	[Nick] [nvarchar](30) NOT NULL,
	[EMail] [nvarchar](60) NOT NULL,
	[Nom] [nvarchar](30) NOT NULL,
	[Prenom] [nvarchar](30) NOT NULL,
	[Adr1] [nvarchar](60) NOT NULL,
	[Ville] [nvarchar](40) NOT NULL,
 CONSTRAINT [PK_Customers] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET IDENTITY_INSERT [dbo].[Users] ON 

INSERT [dbo].[Users] ([Id], [Nick], [EMail], [Nom], [Prenom], [Adr1], [Ville]) VALUES (15000010, N'testUser', N'valbonne_15000010@valbonne.fr', N'Schelbert                     ', N'Peter                         ', N'Adr_517503cd-88a2-483c-bc55-d3a24dfce81d', N'Kiev')
INSERT [dbo].[Users] ([Id], [Nick], [EMail], [Nom], [Prenom], [Adr1], [Ville]) VALUES (15000016, N'Jits', N'valbonne_15000016@valbonne.fr', N'Jovanovic', N'Radovan', N'Bibenlosstrasse                                             ', N'Bremgarten AG                           ')
INSERT [dbo].[Users] ([Id], [Nick], [EMail], [Nom], [Prenom], [Adr1], [Ville]) VALUES (15000017, N'A906301L', N'valbonne_15000017@valbonne.fr', N'Emmenegger                    ', N'David                         ', N'Stockmattstrasse                                            ', N'Baden                                   ')
INSERT [dbo].[Users] ([Id], [Nick], [EMail], [Nom], [Prenom], [Adr1], [Ville]) VALUES (15000029, N'Aah', N'valbonne_15000029@valbonne.fr', N'Hüttmann', N'Aaron', N'Parkettistrasse                                             ', N'Buochs                                  ')
INSERT [dbo].[Users] ([Id], [Nick], [EMail], [Nom], [Prenom], [Adr1], [Ville]) VALUES (15000030, N'Aahhxxuu', N'valbonne_15000030@valbonne.fr', N'Bonassoli                     ', N'Pietro                        ', N'strasse                                                     ', N'BoraBora                                ')
INSERT [dbo].[Users] ([Id], [Nick], [EMail], [Nom], [Prenom], [Adr1], [Ville]) VALUES (15000033, N'Aare42', N'valbonne_15000033@valbonne.fr', N'Mollet', N'Alfred', N'Postfach', N'Aegerten')
INSERT [dbo].[Users] ([Id], [Nick], [EMail], [Nom], [Prenom], [Adr1], [Ville]) VALUES (15000035, N'Aargauer', N'valbonne_15000035@valbonne.fr', N'Byland', N'Wolfgang', N'Laerchenweg 13                                              ', N'Buchs                                   ')
INSERT [dbo].[Users] ([Id], [Nick], [EMail], [Nom], [Prenom], [Adr1], [Ville]) VALUES (15000036, N'Aaron', N'valbonne_15000036@valbonne.fr', N'Stuebi', N'Christian', N'oberer Stutz', N'Guggisberg')
SET IDENTITY_INSERT [dbo].[Users] OFF
/****** Object:  StoredProcedure [dbo].[GetUserById]    Script Date: 27/12/2016 11:47:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetUserById]
(
	@id int
)
AS
BEGIN
	SET NOCOUNT ON;

	SELECT [Id]
		,[Nick]
		,[EMail]
		,[Nom]
		,[Prenom]
		,[Adr1]
		,[Ville]
	FROM [dbo].[Users] WITH(NOLOCK)
	WHERE [id] = @id;

END


GO
/****** Object:  StoredProcedure [dbo].[GetUsers]    Script Date: 27/12/2016 11:47:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetUsers]
AS
BEGIN
	SET NOCOUNT ON;

	SELECT [Id]
		,[Nick]
		,[EMail]
		,[Nom]
		,[Prenom]
		,[Adr1]
		,[Ville]
	FROM [dbo].[Users] WITH(NOLOCK)

END


GO
