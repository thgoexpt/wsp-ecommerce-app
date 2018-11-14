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
Dummy test
    User opens a home page on PC
    End of test
