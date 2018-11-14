*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***

*** Test Cases ***
Test Login with valid username and password on mobile phone
    User opens a home page on mobile phone
    User opens dropdown menu
    User sees that he is not logged in on mobile phone
    User opens a login page on mobile phone
    User types valid username
    User types valid password
    User clicks login button
    User sees the login successful dialog
    User opens dropdown menu
    User sees that he is logged in on mobile phone
    End of test