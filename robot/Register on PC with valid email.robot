*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***

*** Test Cases ***
Test Register on PC with valid email
    User opens a home page on PC
    User opens a register page on PC
    User types username in register modal
    User types valid password in register modal
    User types name in register modal
    User types valid email in register modal
    User types address in register modal
    User clicks register button
    User sees the already have that email dialog
    End of test