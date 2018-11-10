*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***

*** Test Cases ***
Test Login on PC with valid username and password
    User opens a home page on PC
    User sees that he is not logged in on PC
    User opens a login page on PC
    User types valid username
    User types valid password
    User clicks login button
    User sees the login successful dialog
    User sees that he is logged in on PC
    End of test
