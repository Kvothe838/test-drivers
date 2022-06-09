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

Status 200: driver saved
Status 400: bad request
Status 403: user unauthorized
Status 409: username is already used
Status 500: internal server error

## GET /drivers?page=[page]

Gets driver by pagination.

Needs authentication and get-users permission (as default in admin profile).

### Query params

- page: number, required

### Returns

Array of:

```
{
  Username: string
  Name: string
  Surname: string
  DNI: string
}
```

Status 200: processed ok
Status 400: bad request
Status 403: user unauthorized
Status 500: internal server error

## GET /drivers/non-travelling

Gets non travelling drivers.

Needs authentication and get-users permission (as default in admin profile).

### Returns

Array of:

```
{
  Username: string
  Name: string
  Surname: string
  DNI: string
}
```

Status 200: processed ok
Status 400: bad request
Status 403: user unauthorized
Status 500: internal server error


## POST /login

Login with profile.


### Body:

```
{
  Username: string, required
  Password: string, required
  Profile: string, required
}
```

Possible profiles:
- admin
- driver

### Returns

Status 200: logged ok
