# Drivers API Challenge

This is a solution to a challenge for a hiring process.

## Challenge information
### Context
We have a website on the Internet and we would like to get some very simple indication of how visitors navigate the pages. For that purpose, we managed to configure our website to send an event every time a visitor navigates to a page. Our website is capable of generating unique identifiers for visitors as a string of characters.
The system generating that event is able to talk to a REST HTTP interface and represents each individual event as a JSON document containing two attributes: the unique identifier of the visitor and the URL of the visited page.
Our product team is starting a new sprint. We are picking the following user story: As a digital marketeer, I need to know how many distinct visitors navigated to a page, knowing its URL.

### Task
Build a GoLang web service capable of:
- Ingesting user navigation JSON events via a REST HTTP endpoint. Each event is
to be ingested via a separate HTTP request (i.e. no batch and no streaming
ingestion).
- Serving the number of distinct visitors for any given page via another REST HTTP
endpoint. The page URL we are interested in should be a query parameter of the
HTTP request. The number of distinct visitors for that URL is returned in a JSON
object.

### Constraints
- There is no need for persistence to a database. Everything can be kept in memory.
- The web service must be capable of handling concurrent requests on both
endpoints.
- Don't solve the data access concurrency problem using an external library

## Description

The API is used by a merchandise service company that is delivered by drivers. This API allows you to enter as an administrator, manage the ABM of the drivers and see a list of them.

[See more details about requirements.](https://docs.google.com/document/d/1jMgaHhiBAJurhoOI6lDN5Z3ChWY-CjUKQKrfthf9wkU/edit)

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
