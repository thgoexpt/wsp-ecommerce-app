*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***

*** Test Cases ***
Test Login failure because of invalid username and password on mobile phone
    User opens a home page on mobile phone
    User opens dropdown menu
    User sees that he is not logged in on mobile phone
    User opens a login page on mobile phone
    User types invalid username
    User types invalid password
    User clicks login button
    User sees the invalid username or password dialog
    User opens dropdown menu
    User sees that he is not logged in on mobile phone
    End of test
