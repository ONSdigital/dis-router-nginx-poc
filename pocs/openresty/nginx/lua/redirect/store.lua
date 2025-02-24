local json = require("cjson")
local api = require("redirect/api")
local ngx_shared_cache = ngx.shared.redirect_config_cache

local store = {}

store.get_cached_config = function()
    local cached_config = ngx_shared_cache:get("redirect_config")
    if cached_config ~= nil then
        local data_table = json.decode(cached_config)
        return data_table
    end
    return nil
end

store.set_cached_config = function(config)
    local json_data = json.encode(config)
    ngx_shared_cache:set("redirect_config", json_data, 30)
end

store.refresh_cached_config = function()
    ngx.log(ngx.INFO, "refreshing redirect configuration")
    local config = api.getRedirectConfig()
    store.set_cached_config(config)
    return config
end

store.get_config = function()
    local cached_config = store.get_cached_config()

    if cached_config then
        ngx.log(ngx.INFO, "using cached configuration")
        return cached_config
    end

    local config = store.refresh_cached_config()

    return config
end

return store
