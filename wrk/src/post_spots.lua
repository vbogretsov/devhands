local PAYLOAD_TEMPLATE = [[
{
  "name": "%s",
  "description": "%s",
  "raiting": %s,
  "country": "%s",
  "city": "%s",
  "lat": %s,
  "lng": %s
}
]]

local token = nil

function init(args)
  local token_file = args[1]
  if not token_file then
    error("expected script argument token_file")
  end

  local token_fd = io.open(token_file, "r")
  if not token_fd then
    error("unable to read token file")
  end

  token = token_fd:read("*a"):gsub("%s+", "")
  token_fd:close()
end

function random_string(length)
  local chars = 'abcdefghijkl mnopqrstuvwxyzABC DEFGHIJKL MNOPQRSTUVWX YZ0123456789'
  local result = ''
  for i = 1, length do
    local rand = math.random(#chars)
    result = result .. chars:sub(rand, rand)
  end
  return result
end

local function random_payload()
  local name = string.format("spot-%d-%d", math.random(1, 10000), math.random(1, 10000))
  local description = random_string(math.random(10, 20))
  local raiting = 10 * math.random()
  local lat = -90 + math.random() * 180
  local lng = -180 + math.random() * 360
  local country = "RU"
  local city = "Moscow"
  local payload = string.format(
    PAYLOAD_TEMPLATE,
    name,
    description,
    raiting,
    country,
    city,
    lat,
    lng
  )
  return payload
end

function request()
  local headers = {
    ["Content-Type"] = "application/json",
    ["Accept"] = "application/json",
    ["Authorization"] = "Bearer " .. token,
  }
  return wrk.format("POST", wrk.path, headers, random_payload())
end
