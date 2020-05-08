Given('a login event') do
  @expected_event = {
    "username": "amy_pond",
    "unix_timestamp": Time.now.to_i,
    "event_uuid": SecureRandom.uuid,
  }
end

When(/^the event is submitted with the (.*)$/) do |ip_address|
  @expected_event['ip_address'] = ip_address
  @response = post('http://app:4567/v1/event', @expected_event)
end

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
