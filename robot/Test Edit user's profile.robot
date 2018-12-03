*** Settings ***
Library    SeleniumLibrary
Test Teardown    End of test
Resource    commons.txt

*** Variables ***

*** Keywords ***
User opens a profile page on PC
    Click Element    id:userProfileIcon

User can see name information the product
    Wait Until Element Is Visible    alertBox    15
    Textfield Value Should Be    id:first_name   Test User

User can see email information the product
    Wait Until Element Is Visible    alertBox    15
    Textfield Value Should Be     id:email   test@example.com

User can see address information the product
    Wait Until Element Is Visible    alertBox    15
    Textfield Value Should Be   id:location   Kasetsart, TH

User open a edit-profile page on PC
    Click Element    id:editProfileBtn

User types name in textfield
    Input Text    id:first_name    Jo

User types email in textfield
    Input Text    id:email    Go@gmail.com

User types address in textfield
    Input Text    id:location   KasetsartV2

User types password in textfield
    Input Text    name:pass-old   test
    Input Text    name:pass   test1
    Input Text    name:pass-repeat   test1

User submit new information
    Click Element    id:submitBtn

User logout on PC
    Click Element    id:logoutIcon

User can see new name information the product
    Wait Until Element Is Visible    alertBox    15
    Textfield Value Should Be    id:first_name   Jo

User can see new email information the product
    Wait Until Element Is Visible    alertBox    15
    Textfield Value Should Be     id:email   Go@gmail.com

User can see new address information the product
    Wait Until Element Is Visible    alertBox    15
    Textfield Value Should Be   id:location   KasetsartV2

User sees the success to change information dialog
    Wait Until Element Is Visible    alertBox    15
    Element Text Should Be    id:successBox    You have successfully edit your profile.

User sees the success to change password dialog
    Wait Until Element Is Visible    alertBox    15
    Element Text Should Be    id:successBox    Update password successfully.




*** Test Cases ***
Test Profile page on PC
    User opens a home page on PC
    User sees that he is not logged in on PC
    User opens a login page on PC
    User types valid username
    User types valid password
    User clicks login button
    User sees the login successful dialog
    User sees that he is logged in on PC
    User opens a profile page on PC
    User can see name information the product
    User can see email information the product
    User can see address information the product
    End of test

Test EditProfile page on PC
    User opens a home page on PC
    User sees that he is not logged in on PC
    User opens a login page on PC
    User types valid username
    User types valid password
    User clicks login button
    User sees the login successful dialog
    User sees that he is logged in on PC
    User opens a profile page on PC
    User can see name information the product
    User can see email information the product
    User can see address information the product
    User open a edit-profile page on PC
    User types name in textfield
    User types email in textfield
    User types address in textfield
    User submit new information
    User can see new name information the product
    User can see new email information the product
    User can see new address information the product
    User sees the success to change information dialog
    End of test

Test Change password page on PC
    User opens a home page on PC
    User sees that he is not logged in on PC
    User opens a login page on PC
    User types valid username
    User types valid password
    User clicks login button
    User sees the login successful dialog
    User sees that he is logged in on PC
    User opens a profile page on PC
    User open a edit-profile page on PC
    User types password in textfield
    User submit new information
    User sees the success to change password dialog
    User logout on PC
    User opens a login page on PC
    User types valid username
    User types new password
    User clicks login button
    User sees the login successful dialog
    User sees that he is logged in on PC
    End of test
