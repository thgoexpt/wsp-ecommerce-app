*** Settings ***
Library    SeleniumLibrary
Test Teardown     End of test

*** Variables ***
${INVALID_USERNAME}    fail
${INVALID_PASSWORD}    fail
${VALID_USERNAME}    tester
${VALID_PASSWORD}    123456
${VALID_EMAIL}    Charin.ta@ku.th
${NAME}    Charin
${MOCK_EMAIL}    ta@ku.th

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
    Element Text Should Be    id:welcomeUser    Welcome, tester



# Mobile Phone
User opens a home page on mobile phone
    Open Browser    http://localhost:8000/mock/    ${browser}
    Set Window Size    600    800
    Wait Until Element Is Visible    alertBox    30

User sees that he is not logged in on mobile phone
    Element Text Should Be    id:welcomeUser-mobile    Welcome, Guest

User sees that he is logged in on mobile phone
    Element Text Should Be    id:welcomeUser-mobile    Welcome, tester

User opens dropdown menu
    Click Element    id:dropdownMenu
    Wait Until Element Is Visible    id:welcomeUser-mobile    30

User opens a login page on mobile phone
    Click Element    id:loginIcon-mobile
    Wait Until Element Is Visible    myModal    30

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
