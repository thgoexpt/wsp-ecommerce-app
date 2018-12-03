*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***

User add product to cart
    Click Element    id:product-add-button

User select product
    Click Element    id:product-name

User can sees product
    Wait Until Element Is Visible    alertBox    15
    Element Text Should Be    id:product-in-cart-name   ${PRODUCT_NAME}

User click on cart
    Mouse Over    id:cart-icon
    Mouse Down    id:cart-icon
    Mouse Up      id:cart-icon

User check a product is added
    Wait Until Element Is Visible    id:productName-cartModal    15

*** Test Cases ***
Test cart modal have product PC
    User opens a home page on PC
    User opens a login page on PC
    User types valid username
    User types valid password
    User clicks login button
    User opens a product page
    Wait Until Element Is Visible    alertBox    30
    User select product
    User add product to cart
    Wait Until Element Is Visible    alertBox    30
    User can sees product
    Execute JavaScript    window.scrollTo(0,1000)
    User click on cart
    User check a product is added
    End of test
