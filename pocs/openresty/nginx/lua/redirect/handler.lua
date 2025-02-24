local store = require("redirect/store")

local function process_redirects()
    local uri = ngx.var.uri
    local redirects = store.get_config()

    ngx.log(ngx.INFO, "incoming uri", uri)
    if redirects ~= nil and redirects[uri] ~= nil then
        ngx.redirect(redirects[uri])
    end

end

process_redirects()
