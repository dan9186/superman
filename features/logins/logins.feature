Feature:
  As a client
  I want to submit information about a user's login
  So I can see locational details and additional info about the user's activity

  Scenario Outline: A login event is sent
    Given a login event
    When the event is submitted with the <ip_address>
    Then I can see the contextual info about the event includes <latitude>, <longitude>, and <radius>

    Examples:
      | ip_address     | latitude | longitude | radius |
      | 206.81.252.6   | 38.9206  | -76.8787  | 1000   |
      | 61.171.166.100 | 31.1458  | 121.6821  | 20     |
