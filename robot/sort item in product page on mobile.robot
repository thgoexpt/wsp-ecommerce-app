*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***

User select sort method
    Click Element    id:sort-method

User select low to high option
    Click Element    id:low-to-high

User select high to low option
    Click Element    id:high-to-low

User can sees first product with lowest price
    Wait Until Element Is Visible    alertBox    15
    Element Text Should Be    id:product-price   $300.10

User can sees first product with highest price
    Wait Until Element Is Visible    alertBox    15
    Element Text Should Be    id:product-price   $500.00

*** Test Cases ***
Test sort item by price(low to high) on mobile
    User opens a home page on mobile phone
    User opens dropdown menu
    User opens a product page on mobile
    Wait Until Element Is Visible    alertBox    30
    User select sort method
    User select low to high option
    Wait Until Element Is Visible    alertBox    30
    User can sees first product with lowest price
    End of test

Test sort item by price(high to low) on mobile
    User opens a home page on mobile phone
    User opens dropdown menu
    User opens a product page on mobile
    Wait Until Element Is Visible    alertBox    30
    User select sort method
    User select high to low option
    Wait Until Element Is Visible    alertBox    30
    User can sees first product with highest price
    End of test
