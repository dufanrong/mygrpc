-- Define the base URL.
local base_url = "http://localhost:8082/"

-- Define the different cases to test.
local cases = {1, 2, 3, 4}

-- Function to send HTTP requests for each case.
request = function()
  -- Randomly select a case from the cases table.
  local case = cases[math.random(#cases)]
  
  -- Construct the URL with the selected case.
  local url = base_url .. "?case=" .. case
  
  -- Send the HTTP GET request to the constructed URL.
  return wrk.format("GET", url)
end
