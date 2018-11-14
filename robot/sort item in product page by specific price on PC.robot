*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***

User select sort method
    Click Element    id:sort-specific

User select up to 200 option
    Click Element    id:up-to-200

User can sees product with price in specific range
    Wait Until Element Is Visible    alertBox    15
    Element Text Should Be    id:product-price   $299.00

*** Test Cases ***
Test sort item by specific price on PC
    User opens a home page on PC
    User opens a product page
    Wait Until Element Is Visible    alertBox    30
    User select sort method
    User select up to 200 option
    Wait Until Element Is Visible    alertBox    30
    User can sees product with price in specific range
    End of test

Test sort item by specific price on mobile
    User opens a home page on mobile phone
    User opens dropdown menu
    User opens a product page on mobile
    Wait Until Element Is Visible    alertBox    30
    User select sort method
    User select up to 200 option
    Wait Until Element Is Visible    alertBox    30
    User can sees product with price in specific range
    End of test
