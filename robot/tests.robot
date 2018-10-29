*** Settings ***
Library    SeleniumLibrary
Test Teardown     End of test

*** Variables ***
${INVALID_USERNAME}    fail
${INVALID_PASSWORD}    fail

*** Keywords ***
# Global
End of test
    Close Browser

User sees the invalid username or password dialog
    Wait Until Element Is Visible    alertBox
    Element Text Should Be    id:warningBox    Warning: Invalid username/password

User types invalid username
    Input Text    name:username    ${INVALID_USERNAME}

User types invalid password
    Input Text    name:password    ${INVALID_PASSWORD}

User clicks login button
    Click Element    id:loginBtn

# PC
User opens a home page on PC
    Open Browser    http://localhost:8000/mock/    ${browser}
    Set Window Size    1920    1080
    Wait Until Element Is Visible    alertBox    15

User opens a login page on PC
    Click Element    id:loginIcon
    Wait Until Element Is Visible    myModal    15

User sees that he is not logged in on PC
    Element Text Should Be    id:welcomeUser    Welcome, Guest

# Mobile Phone
User opens a home page on mobile phone
    Open Browser    http://localhost:8000/mock/    ${browser}
    Set Window Size    600    800
    Wait Until Element Is Visible    alertBox    15

User sees that he is not logged in on mobile phone
    Element Text Should Be    id:welcomeUser-mobile    Welcome, Guest

User opens dropdown menu
    Click Element    id:dropdownMenu
    Wait Until Element Is Visible    id:welcomeUser-mobile    15

User opens a login page on mobile phone
    Click Element    id:loginIcon-mobile
    Wait Until Element Is Visible    myModal    15

*** Test Cases ***
Test Login on PC failture because of invalid username and password on PC
    User opens a home page on PC
    User sees that he is not logged in on PC
    User opens a login page on PC
    User types invalid username
    User types invalid password
    User clicks login button
    User sees the invalid username or password dialog
    User sees that he is not logged in on PC

Test Login on PC failture because of invalid username and password on mobile phone
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