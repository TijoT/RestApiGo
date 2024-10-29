<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [REST API Test Instructions](#rest-api-test-instructions)
  - [Prerequisite](#prerequisite)
  - [Test Cases](#test-cases)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# REST API Test Instructions
Set up and verify REST API routes using the extension `REST Client` in VS Code. 
Each test file should use `.http` file extension

## Prerequisite
* Ensure the extension `REST Client` is installed in VS Code

## Test Cases
1. Create a user with email and password
2. Login with created user to get token 
3. Use the token for authenticated `POST`, `PUT`, `DELETE` requests
4. No token is needed for GET requests