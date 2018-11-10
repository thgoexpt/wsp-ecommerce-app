*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***

*** Test Cases ***
Test Login on PC failure because of invalid username and password on PC
    User opens a home page on PC
    User sees that he is not logged in on PC
    User opens a login page on PC
    User types invalid username
    User types invalid password
    User clicks login button
    User sees the invalid username or password dialog
    User sees that he is not logged in on PC
    End of test
