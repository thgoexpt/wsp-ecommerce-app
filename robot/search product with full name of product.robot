*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***

*** Test Cases ***
Test search product with full name of product
    User opens a home page on PC
    User opens a product page
    User type full name of product
    User clicks search button
    User can sees the product
    End of test
