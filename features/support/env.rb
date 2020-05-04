# frozen_string_literal: true

require 'json'
require 'net/http'
require 'securerandom'

def do_req(endpoint, method, body)
  uri = URI(endpoint)
  http = Net::HTTP.new(uri.host, uri.port)

  req = method.new(uri)
  req.body = body.to_json unless body.nil?
  resp = http.request(req)

  return resp
end

def post(endpoint, body)
  return do_req(endpoint, Net::HTTP::Post, body)
end
