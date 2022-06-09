# Drivers API

## Description

The API is used by a merchandise service company that is delivered by drivers. This API allows you to enter as an administrator, manage the ABM of the drivers and see a list of them.

## Endpoints

## POST /drivers

Saves driver.

Needs authentication and save-users permission (as default in admin profile).

### Body:

```
{
  Username: string, required
  Password: string, required
  Name: string, required
  Surname: string, required
  DNI: string, required
}
```

### Returns

Status 400: internal bad request
Status 403: user unauthorized
Status 409: username is already used



## GET /drivers?page=[page]

Gets driver by pagination.

Needs authentication and get-users permission (as default in admin profile).

Query params:

- page: number, required

## GET /drivers/non-travelling

Gets non travelling drivers.

Needs authentication and get-users permission (as default in admin profile).
