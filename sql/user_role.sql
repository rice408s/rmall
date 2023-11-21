CREATE TABLE UserRole (
                          UserId INT,
                          RoleId INT,
                          PRIMARY KEY (UserId, RoleId),
                          FOREIGN KEY (UserId) REFERENCES User(Id),
                          FOREIGN KEY (RoleId) REFERENCES Role(Id)
);