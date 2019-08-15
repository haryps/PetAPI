**Pet API**

This is web api for Pet store created using Go language.

---

## Chosen Tech Stack

1. Go language
2. MySql
3. GORM: Go ORM

---

## Reasons of Chosen Tech Stack

1. Go has good performance almost like other lower level languages but it is more simple and reliable
2. GORM currently doesn't support MariaDB so i choose MySql
3. This is the most mature Object Relational Mapping (ORM) library for Go

---

## Infrastructure Requirement

1. Go runtime
2. MySql server

---

## Directory Structure

1. Main function: /cmd/petapi/petapi.go
2. Controllers: /cmd/petapi/controllers
3. Models: /cmd/petapi/models
4. Repository: /cmd/petapi/repository
5. Helper method: /cmd/petapi/helpers
6. Test method: /test

---

## Setup Instructions & Deployment

1. Database connection string is still hardcoded, so please modify: cmd/petapi/models/init.go
2. Build project: go build cmd/petapi/petapi.go
3. Run project: go run cmd/petapi/petapi.go

---