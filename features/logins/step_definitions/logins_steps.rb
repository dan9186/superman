Given('a login event') do
  @expected_event = {
    "username": "amy_pond",
    "unix_timestamp": Time.now.to_i,
    "event_uuid": SecureRandom.uuid,
    "ip_address": "206.81.252.6"
  }
end

When('the event is submitted') do
  @response = post('http://app:4567/v1/event', @expected_event)
end

Then('I can see the contextual info about the event') do
  expect(@response.code.to_i).to(eql(201))
  expect(@response.body).not_to(be_nil(), 'expected: body not nil\ngot: body nil')

  body = JSON.parse(@response.body)
  expect(body).not_to(be_nil(), 'expected: json body not nil\ngot: json body nil')

  expect(body['currentGeo']).not_to(be_nil(), "expected: currentGeo field\ngot: field missing\nbody: #{body.inspect}")
  expect(body['currentGeo']['lat']).not_to(be_nil(), "expected: lat field\ngot: field missing\nbody: #{body.inspect}")
  expect(body['currentGeo']['lon']).not_to(be_nil(), "expected: lon field\ngot: field missing\nbody: #{body.inspect}")
  expect(body['currentGeo']['radius']).not_to(be_nil(), "expected: radius field\ngot: field missing\nbody: #{body.inspect}")
end
