local token = nil
local num_users = nil

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

  num_users = args[2] or 100
end

function request()
  local headers = {
    ["Content-Type"] = "application/json",
    ["Accept"] = "application/json",
    ["Authorization"] = "Bearer " .. token,
  }
  local user_id = string.format("u-%d", math.random(1, num_users))
  return wrk.format("GET", string.format("%s?user_id=%s", wrk.path, user_id), headers)
end
