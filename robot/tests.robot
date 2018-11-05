*** Settings ***
Library    SeleniumLibrary
Test Teardown     End of test

*** Variables ***
${INVALID_USERNAME}    fail
${INVALID_PASSWORD}    fail
${VALID_USERNAME}    test
${VALID_PASSWORD}    test
${VALID_EMAIL}    test@example.com
${NAME}    Charin
${MOCK_EMAIL}    ta@ku.th
${SOME_PRODUCT_NAME}    Kuro
${PRODUCT_NAME}    Kurobuta
${FULL_PRODUCT_NAME}    Kurobuta (Chicken)

*** Keywords ***
# Global
End of test
    Close Browser

User sees the invalid username or password dialog
    Wait Until Element Is Visible    alertBox    30
    Element Text Should Be    id:warningBox    Warning: Invalid username/password

User sees the login successful dialog
    Wait Until Element Is Visible    alertBox    30
    Element Text Should Be    id:successBox    Login successful

User sees the success to registration dialog
    Wait Until Element Is Visible    alertBox    30
    Element Text Should Be    id:successBox    User created successful, please login.

User sees the already have that username dialog
    Wait Until Element Is Visible    alertBox    30
    Element Text Should Be    id:warningBox    Warning: Username already exists

User sees the already have that email dialog
    Wait Until Element Is Visible    alertBox    30
    Element Text Should Be    id:warningBox    Warning: Email already in use

User can sees the product
    Wait Until Element Is Visible    alertBox    30
    Element Text Should Be    id:product-name   ${FULL_PRODUCT_NAME}



User types valid username in register modal
    Input Text    id:regisUsername    ${VALID_USERNAME}

User types valid password in register modal
    Input Text    id:regisPass    ${VALID_PASSWORD}

User types username in register modal
    Input Text    id:regisUsername    KKKKK

User types valid email in register modal
    Input Text    id:regisEmail    ${VALID_EMAIL}

User types email in register modal
    Input Text    id:regisEmail    ${MOCK_EMAIL}

User types name in register modal
    Input Text    name:name    ${NAME}

User types address in register modal
    Input Text    name:address    Kasetsart



User types invalid username
    Input Text    name:username    ${INVALID_USERNAME}

User types invalid password
    Input Text    name:password    ${INVALID_PASSWORD}

User types valid username
    Input Text    name:username    ${VALID_USERNAME}

User types valid password
    Input Text    name:password    ${VALID_PASSWORD}

User clicks login button
    Click Element    id:loginBtn

User clicks register button
    Click Element    id:regisBtn

User clicks search button
    Click Element    id:search-BTN


# PC
User opens a home page on PC
    Open Browser    http://localhost:8000/mock/    ${browser}
    Set Window Size    1920    1080
    Wait Until Element Is Visible    alertBox    30

User opens a login page on PC
    Click Element    id:loginIcon
    Wait Until Element Is Visible    myModal    30

User opens a register page on PC
    Click Element    id:registerIcon
    Wait Until Element Is Visible    myModal_regis    30

User sees that he is not logged in on PC
    Element Text Should Be    id:welcomeUser    Welcome, Guest

User sees that he is logged in on PC
    Element Text Should Be    id:welcomeUser    Welcome, test

User opens a product page
    Click Element    id:Shop-BTN

User type full name of product
    Input Text    id:search    ${PRODUCT_NAME}

User type some part of product's name
    Input Text    id:search    ${SOME_PRODUCT_NAME}



# Mobile Phone
User opens a home page on mobile phone
    Open Browser    http://localhost:8000/mock/    ${browser}
    Set Window Size    600    800
    Wait Until Element Is Visible    alertBox    30

User sees that he is not logged in on mobile phone
    Element Text Should Be    id:welcomeUser-mobile    Welcome, Guest

User sees that he is logged in on mobile phone
    Element Text Should Be    id:welcomeUser-mobile    Welcome, test

User opens dropdown menu
    Click Element    id:dropdownMenu
    Wait Until Element Is Visible    id:welcomeUser-mobile    30

User opens a login page on mobile phone
    Click Element    id:loginIcon-mobile
    Wait Until Element Is Visible    myModal    30

User opens a register page on mobile phone
    Click Element    id:registerIcon-mobile
    Wait Until Element Is Visible    myModal_regis    30

User opens a product page on mobile
    Click Element    id:Shop-BTN-Mobile

*** Test Cases ***
Test Login on PC failure because of invalid username and password on PC
    User opens a home page on PC
    User sees that he is not logged in on PC
    User opens a login page on PC
    User types invalid username
    User types invalid password
    User clicks login button
    User sees the invalid username or password dialog
    User sees that he is not logged in on PC
    End of test

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

Test Login on PC failure because of invalid password on PC
    User opens a home page on PC
    User sees that he is not logged in on PC
    User opens a login page on PC
    User types valid username
    User types invalid password
    User clicks login button
    User sees the invalid username or password dialog
    User sees that he is not logged in on PC
    End of test

Test Login failure because of invalid password on mobile phone
    User opens a home page on mobile phone
    User opens dropdown menu
    User sees that he is not logged in on mobile phone
    User opens a login page on mobile phone
    User types valid username
    User types invalid password
    User clicks login button
    User sees the invalid username or password dialog
    User opens dropdown menu
    User sees that he is not logged in on mobile phone
    End of test


Test Login on PC with valid username and password
    User opens a home page on PC
    User sees that he is not logged in on PC
    User opens a login page on PC
    User types valid username
    User types valid password
    User clicks login button
    User sees the login successful dialog
    User sees that he is logged in on PC
    End of test

Test Login with valid username and password on mobile phone
    User opens a home page on mobile phone
    User opens dropdown menu
    User sees that he is not logged in on mobile phone
    User opens a login page on mobile phone
    User types valid username
    User types valid password
    User clicks login button
    User sees the login successful dialog
    User opens dropdown menu
    User sees that he is logged in on mobile phone
    End of test

Test Register on PC with valid username
    User opens a home page on PC
    User opens a register page on PC
    User types valid username in register modal
    User types valid password in register modal
    User types name in register modal
    User types email in register modal
    User types address in register modal
    User clicks register button
    User sees the already have that username dialog
    End of test

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

Test Register on mobile phone with valid username
    User opens a home page on mobile phone
    User opens dropdown menu
    User opens a register page on mobile phone
    User types valid username in register modal
    User types valid password in register modal
    User types name in register modal
    User types email in register modal
    User types address in register modal
    User clicks register button
    User sees the already have that username dialog
    End of test

Test Register on mobile phone with valid email
    User opens a home page on mobile phone
    User opens dropdown menu
    User opens a register page on mobile phone
    User types username in register modal
    User types valid password in register modal
    User types name in register modal
    User types valid email in register modal
    User types address in register modal
    User clicks register button
    User sees the already have that email dialog
    End of test

Test search product with full name of product
    User opens a home page on PC
    User opens a product page
    User type full name of product
    User clicks search button
    User can sees the product
    End of test


Test search product with some part of product's name
    User opens a home page on PC
    User opens a product page
    User type some part of product's name
    User clicks search button
    User can sees the product
    End of test

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


Test search product with some part of product's name on mobile
    User opens a home page on mobile phone
    User opens dropdown menu
    User opens a product page on mobile
    Wait Until Element Is Visible    alertBox    30
    User type some part of product's name
    User clicks search button
    Wait Until Element Is Visible    alertBox    30
    User can sees the product
    End of test
