*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***

User types employee username
    Input Text    name:username    emp

User types employee password
    Input Text    name:password    emp

User add product to cart
    Click Element    id:product-add-button

User select product
    Click Element    id:product-name

User checkout a product
    Click Element    id:checkoutBtn
    Click Element    id:paidBtn

User go to homepage
    Click Element    id:Home-BTN

User open stock's page
    Mouse Over    id:adminAction-BTN
    Mouse Over    id:stock-BTN
    Mouse Down    id:stock-BTN
    Mouse Up      id:stock-BTN

User success the checkout's process
    Wait Until Element Is Visible    alertBox    15
    Element Text Should Be    id:successBox    Thank you for your purchase, please come again.

User can sees product
    Wait Until Element Is Visible    alertBox    15
    Element Text Should Be    id:product-in-cart-name   ${PRODUCT_NAME}

User can sees product decreasing
    Wait Until Element Is Visible    alertBox    15
    Element Text Should Be    id:name-in-stock   ${PRODUCT_NAME}
    Element Text Should Be    id:amount-in-stock   49



*** Test Cases ***
Test checkout on PC
    User opens a home page on PC
    User sees that he is not logged in on PC
    User opens a login page on PC
    User types employee username
    User types employee password
    User clicks login button
    User opens a product page
    User select product
    User add product to cart
    User can sees product
    User checkout a product
    User success the checkout's process
    sleep    5
    User open stock's page
    User can sees product decreasing
    End of test
