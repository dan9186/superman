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

  Scenario: Preceding login events are sent
    Given a login event
    And a preceding login event
    When the event is submitted
    Then I can see the contextual info about the event
    And I can see the preceding access info

  Scenario: Out of order login events are sent
    Given a login event
    And a subsequent login event
    When the event is submitted
    Then I can see the contextual info about the event
    And I can see the subsequent access info

  Scenario: Subsequent and preceding events already exist
    Given a login event
    And a preceding login event
    And a subsequent login event
    When the event is submitted
    Then I can see the contextual info about the event
    And I can see the preceding access info
    And I can see the subsequent access info

  Scenario: Multiple preceding events are sent
    Given a login event
    And multiple preceding login events
    When the event is submitted
    Then I can see the contextual info about the event
    And I can see the closest preceding access info

  Scenario: Multiple subsequent events are sent
    Given a login event
    And multiple subsequent login events
    When the event is submitted
    Then I can see the contextual info about the event
    And I can see the closest subsequent access info

  Scenario: A suspicious preceding event
    Given a login event
    And a suspicious preceding login event
    When the event is submitted
    Then I can see the preceding event is suspicious

  Scenario: A suspicious subsequent event
    Given a login event
    And a suspicious subsequent login event
    When the event is submitted
    Then I can see the suspicious event is suspicious

  Scenario: A suspicious preceding and subsequent events
    Given a login event
    And a suspicious preceding login event
    And a suspicious subsequent login event
    When the event is submitted
    Then I can see the event preceding event is suspicious
