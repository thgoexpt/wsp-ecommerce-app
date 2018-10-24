*** Settings ***
Library    SeleniumLibrary
Test Teardown     End of test

*** Variables ***
${INVALID_USERNAME}    fail
${INVALID_PASSWORD}    fail

*** Keywords ***
End of test
    Close Browser

User opens a home page on PC
    Open Browser    http://localhost:8000    ${browser}
    Set Window Size    1920    1080
    Wait Until Element Is Visible    alertBox    15

User opens a login page
    Click Element    id:loginIcon
    Wait Until Element Is Visible    myModal    15

User types invalid username
    Input Text    name:username    ${INVALID_USERNAME}

User types invalid password
    Input Text    name:password    ${INVALID_PASSWORD}

User clicks login button
    Click Element    id:loginBtn

User sees the invalid username or password dialog
    Wait Until Element Is Visible    alertBox
    Element Text Should Be    id:warningBox    Warning: Invalid username/password

User sees that he is not logged in
    Element Text Should Be    id:welcomeUser    Welcome, Guest

*** Test Cases ***
Test Login on PC failture because of invalid username and password
    User opens a home page on PC
    User sees that he is not logged in
    User opens a login page
    User types invalid username
    User types invalid password
    User clicks login button
    User sees the invalid username or password dialog
    User sees that he is not logged in