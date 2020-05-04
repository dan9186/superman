Given('a login event') do
  @expected_event = {
    "username": "amy_pond",
    "unix_timestamp": Time.now.to_i,
    "event_uuid": SecureRandom.uuid,
    "ip_address": "206.81.252.6"
  }
end

When('the event is submitted') do
  pending
end

Then('I can see the contextual info about the event') do
  pending
end
