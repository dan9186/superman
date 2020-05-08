# Givens
Given('a login event') do
  @expected_event = {
    "username": "amy_pond",
    "unix_timestamp": Time.now.to_i,
    "event_uuid": SecureRandom.uuid,
    "ip_address": "4.4.4.4",
  }
end

Given("a preceding login event") do
  @expected_preceding_event = {
    "username": "amy_pond",
    "unix_timestamp": Time.now.to_i,
    "event_uuid": SecureRandom.uuid,
    "ip_address": "56.3.181.4",
  }

  expected_event = @expected_event
  @expected_event = @expected_preceding_event
  steps %{
    When the event is submitted
  }
  @expected_event = expected_event
end

Given("subsequent login event") do
  @expected_subsequent_event = {
    "username": "amy_pond",
    "unix_timestamp": Time.now.to_i,
    "event_uuid": SecureRandom.uuid,
    "ip_address": "36.12.93.24",
  }

  expected_event = @expected_event
  @expected_event = @expected_subsequent_event
  steps %{
    When the event is submitted
  }
  @expected_event = expected_event
end

# Whens
When(/^the event is submitted with the (.*)$/) do |ip_address|
  @expected_event['ip_address'] = ip_address
  steps %{
    When the event is submitted
  }
end

When("the event is submitted") do
  @response = post('http://app:4567/v1/event', @expected_event)
end

# Thens
Then('I can see the contextual info about the event includes {float}, {float}, and {int}') do |latitude, longitude, radius|
  expect(@response.code.to_i).to(eql(201))
  expect(@response.body).not_to(be_nil(), 'expected: body not nil\ngot: body nil')

  body = JSON.parse(@response.body)
  expect(body).not_to(be_nil(), 'expected: json body not nil\ngot: json body nil')

  expect(body['currentGeo']).not_to(be_nil(), "expected: currentGeo field\ngot: field missing\nbody: #{body.inspect}")

  lat = body['currentGeo']['lat']
  expect(lat).not_to(be_nil(), "expected: lat field\ngot: field missing\nbody: #{body.inspect}")
  expect(lat).to(eql(latitude))

  lon = body['currentGeo']['lon']
  expect(lon).not_to(be_nil(), "expected: lon field\ngot: field missing\nbody: #{body.inspect}")
  expect(lon).to(eql(longitude))

  rad = body['currentGeo']['radius']
  expect(rad).not_to(be_nil(), "expected: radius field\ngot: field missing\nbody: #{body.inspect}")
  expect(rad).to(eql(radius))

  expect(body['travelToCurrentGeoSuspicious']).to(eql(false))
  expect(body['travelFromCurrentGeoSuspicious']).to(eql(false))
end

Then("I can see the contextual info about the event") do
  expect(@response.code.to_i).to(eql(201))
  expect(@response.body).not_to(be_nil(), 'expected: body not nil\ngot: body nil')

  body = JSON.parse(@response.body)
  expect(body).not_to(be_nil(), 'expected: json body not nil\ngot: json body nil')

  expect(body['currentGeo']).not_to(be_nil(), "expected: currentGeo field\ngot: field missing\nbody: #{body.inspect}")

  lat = body['currentGeo']['lat']
  expect(lat).not_to(be_nil(), "expected: lat field\ngot: field missing\nbody: #{body.inspect}")
  expect(lat).to(eql(37.751))

  lon = body['currentGeo']['lon']
  expect(lon).not_to(be_nil(), "expected: lon field\ngot: field missing\nbody: #{body.inspect}")
  expect(lon).to(eql(-97.822))

  rad = body['currentGeo']['radius']
  expect(rad).not_to(be_nil(), "expected: radius field\ngot: field missing\nbody: #{body.inspect}")
  expect(rad).to(eql(1000))

  expect(body['travelToCurrentGeoSuspicious']).to(eql(false))
  expect(body['travelFromCurrentGeoSuspicious']).to(eql(false))
end

Then("I can see the preceding access info") do
  body = JSON.parse(@response.body)
  expect(body).not_to(be_nil(), 'expected: json body not nil\ngot: json body nil')

  expect(body['precedingIpAccess']).not_to(be_nil(), "expected: precedingIpAccess field\ngot: field missing\nbody: #{body.inspect}")

  ip = body['precedingIpAccess']['ip']
  expect(ip).not_to(be_nil(), "expected: ip field\ngot: field missing\nbody: #{body.inspect}")
  expect(ip).to(eql("figure me out"))

  speed = body['precedingIpAccess']['speed']
  expect(speed).not_to(be_nil(), "expected: speed field\ngot: field missing\nbody: #{body.inspect}")
  expect(speed).to(eql("figure me out"))

  lat = body['precedingIpAccess']['lat']
  expect(lat).not_to(be_nil(), "expected: lat field\ngot: field missing\nbody: #{body.inspect}")
  expect(lat).to(eql("figure me out"))

  lon = body['precedingIpAccess']['lon']
  expect(lon).not_to(be_nil(), "expected: lon field\ngot: field missing\nbody: #{body.inspect}")
  expect(lon).to(eql("figure me out"))

  radius = body['precedingIpAccess']['radius']
  expect(radius).not_to(be_nil(), "expected: radius field\ngot: field missing\nbody: #{body.inspect}")
  expect(radius).to(eql("figure me out"))

  timestamp = body['precedingIpAccess']['timestamp']
  expect(timestamp).not_to(be_nil(), "expected: timestamp field\ngot: field missing\nbody: #{body.inspect}")
  expect(timestamp).to(eql("figure me out"))
end

Then("I can see the subsequent access info") do
  body = JSON.parse(@response.body)
  expect(body).not_to(be_nil(), 'expected: json body not nil\ngot: json body nil')

  expect(body['subsequentIpAccess']).not_to(be_nil(), "expected: subsequentIpAccess field\ngot: field missing\nbody: #{body.inspect}")

  ip = body['subsequentIpAccess']['ip']
  expect(ip).not_to(be_nil(), "expected: ip field\ngot: field missing\nbody: #{body.inspect}")
  expect(ip).to(eql("figure me out"))

  speed = body['subsequentIpAccess']['speed']
  expect(speed).not_to(be_nil(), "expected: speed field\ngot: field missing\nbody: #{body.inspect}")
  expect(speed).to(eql("figure me out"))

  lat = body['subsequentIpAccess']['lat']
  expect(lat).not_to(be_nil(), "expected: lat field\ngot: field missing\nbody: #{body.inspect}")
  expect(lat).to(eql("figure me out"))

  lon = body['subsequentIpAccess']['lon']
  expect(lon).not_to(be_nil(), "expected: lon field\ngot: field missing\nbody: #{body.inspect}")
  expect(lon).to(eql("figure me out"))

  radius = body['subsequentIpAccess']['radius']
  expect(radius).not_to(be_nil(), "expected: radius field\ngot: field missing\nbody: #{body.inspect}")
  expect(radius).to(eql("figure me out"))

  timestamp = body['subsequentIpAccess']['timestamp']
  expect(timestamp).not_to(be_nil(), "expected: timestamp field\ngot: field missing\nbody: #{body.inspect}")
  expect(timestamp).to(eql("figure me out"))
end
