local http = require("resty.http")
local json = require("cjson")

local api = {}

api.getRedirectConfig = function()
    local httpc = http.new()
    local res, err = httpc:request_uri("http://router-api:3000/redirects", {
        method = "GET",
    })
    if not res then
        ngx.log(ngx.ERR, "request failed: ", err)
        return
    end

    ngx.log(ngx.INFO, "loading redirects")

    local data = json.decode(res.body)

    local redirects = {}
    ngx.log(ngx.INFO, #data)
    for i,redirect in ipairs(data) do 
        ngx.log(ngx.INFO, "loading redirect")
        redirects[redirect["from"]] = redirect["to"]
    end

    return redirects
end

return api
