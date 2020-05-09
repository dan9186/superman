# Givens
Given('a login event') do
  # randomly pick a time in the past year for the focal event
  time = Time.now.to_i - rand(1..31536000)

  @expected_event = {
    "username": "cuketest",
    "unix_timestamp": Time.now.to_i,
    "event_uuid": SecureRandom.uuid,
    "ip_address": "67.171.166.100", #West Linn, OR
  }
end

Given("a preceding login event") do
  time = @expected_event[:unix_timestamp] - 3600

  @expected_preceding_event = {
    "username": "cuketest",
    "unix_timestamp": time,
    "event_uuid": SecureRandom.uuid,
    "ip_address": "66.39.178.22", #Bend, OR
  }

  expected_event = @expected_event
  @expected_event = @expected_preceding_event
  steps %{
    When the event is submitted
  }
  @expected_event = expected_event
end

Given("a subsequent login event") do
  time = @expected_event[:unix_timestamp] + 3600

  @expected_subsequent_event = {
    "username": "cuketest",
    "unix_timestamp": time,
    "event_uuid": SecureRandom.uuid,
    "ip_address": "47.39.62.215", #Seaside, OR
  }

  expected_event = @expected_event
  @expected_event = @expected_subsequent_event
  steps %{
    When the event is submitted
  }
  @expected_event = expected_event
end

Given("multiple preceding login events") do
  time = @expected_event[:unix_timestamp] - 4000

  older_preceding_event = {
    "username": "cuketest",
    "unix_timestamp": time,
    "event_uuid": SecureRandom.uuid,
    "ip_address": "3.39.21.4",
  }

  expected_event = @expected_event
  @expected_event = older_preceding_event
  steps %{
    When the event is submitted
  }
  @expected_event = expected_event

  time = @expected_event[:unix_timestamp] - 3600

  @expected_preceding_event = {
    "username": "cuketest",
    "unix_timestamp": time,
    "event_uuid": SecureRandom.uuid,
    "ip_address": "66.39.178.22", #Bend, OR
  }

  expected_event = @expected_event
  @expected_event = @expected_preceding_event
  steps %{
    When the event is submitted
  }
  @expected_event = expected_event
end

Given("multiple subsequent login events") do
  time = @expected_event[:unix_timestamp] + 4000

  older_subsequent_event = {
    "username": "cuketest",
    "unix_timestamp": time,
    "event_uuid": SecureRandom.uuid,
    "ip_address": "4.181.56.3",
  }

  expected_event = @expected_event
  @expected_event = older_subsequent_event
  steps %{
    When the event is submitted
  }
  @expected_event = expected_event

  time = @expected_event[:unix_timestamp] + 3600

  @expected_subsequent_event = {
    "username": "cuketest",
    "unix_timestamp": time,
    "event_uuid": SecureRandom.uuid,
    "ip_address": "47.39.62.215", #Seaside, OR
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

  expect(body['travelToCurrentGeoSuspicious']).to(eql(false), "expected: false\ngot: true\nfield: travelToCurrentGeoSuspicious\n")
  expect(body['travelFromCurrentGeoSuspicious']).to(eql(false), "expected: false\ngot: true\nfield: travelFromCurrentGeoSuspicious\n")
end

Then("I can see the contextual info about the event") do
  expect(@response.code.to_i).to(eql(201))
  expect(@response.body).not_to(be_nil(), 'expected: body not nil\ngot: body nil')

  body = JSON.parse(@response.body)
  expect(body).not_to(be_nil(), 'expected: json body not nil\ngot: json body nil')

  expect(body['currentGeo']).not_to(be_nil(), "expected: currentGeo field\ngot: field missing\nbody: #{body.inspect}")

  lat = body['currentGeo']['lat']
  expect(lat).not_to(be_nil(), "expected: lat field\ngot: field missing\nbody: #{body.inspect}")
  expect(lat).to(eql(45.3642))

  lon = body['currentGeo']['lon']
  expect(lon).not_to(be_nil(), "expected: lon field\ngot: field missing\nbody: #{body.inspect}")
  expect(lon).to(eql(-122.6443))

  rad = body['currentGeo']['radius']
  expect(rad).not_to(be_nil(), "expected: radius field\ngot: field missing\nbody: #{body.inspect}")
  expect(rad).to(eql(5))

  expect(body['travelToCurrentGeoSuspicious']).to(eql(false), "expected: false\ngot: true\nfield: travelToCurrentGeoSuspicious\n")
  expect(body['travelFromCurrentGeoSuspicious']).to(eql(false), "expected: false\ngot: true\nfield: travelFromCurrentGeoSuspicious\n")
end

Then("I can see the preceding access info") do
  body = JSON.parse(@response.body)
  expect(body).not_to(be_nil(), 'expected: json body not nil\ngot: json body nil')

  expect(body['precedingIpAccess']).not_to(be_nil(), "expected: precedingIpAccess field\ngot: field missing\nbody: #{body.inspect}")

  ip = body['precedingIpAccess']['ip_address']
  expect(ip).not_to(be_nil(), "expected: ip field\ngot: field missing\nbody: #{body.inspect}")
  expect(ip).to(eql("66.39.178.22"))

  speed = body['precedingIpAccess']['speed']
  expect(speed).not_to(be_nil(), "expected: speed field\ngot: field missing\nbody: #{body.inspect}")
  expect(speed).to(eql(98))

  lat = body['precedingIpAccess']['lat']
  expect(lat).not_to(be_nil(), "expected: lat field\ngot: field missing\nbody: #{body.inspect}")
  expect(lat).to(eql(44.0185))

  lon = body['precedingIpAccess']['lon']
  expect(lon).not_to(be_nil(), "expected: lon field\ngot: field missing\nbody: #{body.inspect}")
  expect(lon).to(eql(-121.2984))

  radius = body['precedingIpAccess']['radius']
  expect(radius).not_to(be_nil(), "expected: radius field\ngot: field missing\nbody: #{body.inspect}")
  expect(radius).to(eql(10))

  timestamp = body['precedingIpAccess']['timestamp']
  expect(timestamp).not_to(be_nil(), "expected: timestamp field\ngot: field missing\nbody: #{body.inspect}")
  expect(timestamp).to(eql(@expected_preceding_event[:unix_timestamp]))
end

Then("I can see the subsequent access info") do
  body = JSON.parse(@response.body)
  expect(body).not_to(be_nil(), 'expected: json body not nil\ngot: json body nil')

  expect(body['subsequentIpAccess']).not_to(be_nil(), "expected: subsequentIpAccess field\ngot: field missing\nbody: #{body.inspect}")

  ip = body['subsequentIpAccess']['ip_address']
  expect(ip).not_to(be_nil(), "expected: ip field\ngot: field missing\nbody: #{body.inspect}")
  expect(ip).to(eql("47.39.62.215"))

  speed = body['subsequentIpAccess']['speed']
  expect(speed).not_to(be_nil(), "expected: speed field\ngot: field missing\nbody: #{body.inspect}")
  expect(speed).to(eql(50))

  lat = body['subsequentIpAccess']['lat']
  expect(lat).not_to(be_nil(), "expected: lat field\ngot: field missing\nbody: #{body.inspect}")
  expect(lat).to(eql(45.9937))

  lon = body['subsequentIpAccess']['lon']
  expect(lon).not_to(be_nil(), "expected: lon field\ngot: field missing\nbody: #{body.inspect}")
  expect(lon).to(eql(-123.9243))

  radius = body['subsequentIpAccess']['radius']
  expect(radius).not_to(be_nil(), "expected: radius field\ngot: field missing\nbody: #{body.inspect}")
  expect(radius).to(eql(20))

  timestamp = body['subsequentIpAccess']['timestamp']
  expect(timestamp).not_to(be_nil(), "expected: timestamp field\ngot: field missing\nbody: #{body.inspect}")
  expect(timestamp).to(eql(@expected_subsequent_event[:unix_timestamp]))
end

Then("I can see the closest preceding access info") do
    steps %{
      Then I can see the preceding access info
    }
end

Then("I can see the closest subsequent access info") do
    steps %{
      Then I can see the subsequent access info
    }
end

After do
  delete('http://app:4567/v1/cleanup')
end
