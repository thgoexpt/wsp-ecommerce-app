*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***

*** Test Cases ***
Test search product with some part of product's name
    User opens a home page on PC
    User opens a product page
    User type some part of product's name
    User clicks search button
    User can sees the product
    End of test
