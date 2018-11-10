*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***

*** Test Cases ***
Test search product with full name of product on mobile
    User opens a home page on mobile phone
    User opens dropdown menu
    User opens a product page on mobile
    Wait Until Element Is Visible    alertBox    30
    User type full name of product
    User clicks search button
    Wait Until Element Is Visible    alertBox    30
    User can sees the product
    End of test