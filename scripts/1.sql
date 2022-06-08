BEGIN;

CREATE TABLE Permission(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100)
);

CREATE TABLE Profile(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100)
);

CREATE TABLE Profiles_Permissions(
    profile_id INT,
    permission_id INT,
    PRIMARY KEY(permission_id, profile_id),
    CONSTRAINT FK_Permission FOREIGN KEY (permission_id)
    REFERENCES Permission(id),
    CONSTRAINT FK_Profile FOREIGN KEY (profile_id)
    REFERENCES Profile(id)
);

CREATE TABLE User(
    id INT PRIMARY KEY AUTO_INCREMENT,  
    username VARCHAR(100) NOT NULL,
    hash TEXT NOT NULL,
    profile_id INT,
    CONSTRAINT FK_Profile FOREIGN KEY (profile_id)
    REFERENCES Profile(id)
);

CREATE TABLE Driver(
    id INT PRIMARY KEY AUTO_INCREMENT,  
    dni VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    user_id INT,
    CONSTRAINT FK_User FOREIGN KEY (user_id)
    REFERENCES User(id)
);

CREATE TABLE Travel(
    id INT PRIMARY KEY AUTO_INCREMENT,
    driver_id INT NOT NULL,
    start BIGINT,
    end BIGINT,
    CONSTRAINT FK_Driver FOREIGN KEY (driver_id)
    REFERENCES Driver(id)
);

COMMIT;