BEGIN;

INSERT INTO Permission (name)
VALUES ("create-drivers"), ("get-drivers"), ("create-admins");

INSERT INTO Profile (name)
VALUES ("admin"), ("driver");

INSERT INTO Profiles_Permissions(profile_id, permission_id)
VALUES (1, 1), (1, 2), (1, 3);

INSERT INTO User (username, hash, profile_id)
VALUES ("root", "$2a$10$KSeGPuBJMCl1Krs92itmcOsj6CVhwHLfb6y4UaXRV7OSCDDhYKbnC", 1);
-- Hash is "admin"

COMMIT;