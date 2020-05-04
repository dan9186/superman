Feature:
  As a client
  I want to submit information about a user's login
  So I can see locational details and additional info about the user's activity

  Scenario: A login event is sent
    Given a login event
    When the event is submitted
    Then I can see the contextual info about the event
