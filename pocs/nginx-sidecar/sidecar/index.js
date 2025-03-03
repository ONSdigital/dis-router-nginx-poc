const express = require("express");
const axios = require("axios");
const fs = require("fs");
const { exec } = require("child_process");

const app = express();
const PORT = process.env.PORT || 5000;
const API_URL = process.env.API_URL || "http://localhost:7777";
const REDIRECT_CONFIG_PATH = process.env.REDIRECT_CONFIG_PATH || "/etc/nginx/conf.d/redirects.conf";
const ROUTE_CONFIG_PATH = process.env.ROUTE_CONFIG_PATH || "/etc/nginx/conf.d/routes.conf";

app.use(express.json());

const restartNginx = function() {
    exec("exec nginx -s reload", (err, stdout, stderr) =>{
        if (err) {
            console.error("error reloading nginx")
            console.error(err.message)
        } else  {
            console.log("nginx reloaded successfully");
        }
    })
}

const updateRouteConfig = async function() {
    try {
        console.log("fetching latest route config");
        const response = await axios.get(`${API_URL}/routes`)
        const routes = response.data;
        let routeConfig = routes.map(route => `location ${route.path} { proxy_pass http://${route.service}; }`).join("\n");
        fs.writeFileSync(ROUTE_CONFIG_PATH, routeConfig);
        console.log("route config updated");
    } catch (err) {
        console.error("failed to update nginx config")
        console.error(err.message)
    }
}

const updateRedirectConfig = async function() {
    try {
        console.log("fetching latest rediect config");
        const response = await axios.get(`${API_URL}/redirects`)
        const redirects = response.data;
        let redirectConfig = redirects.map(redirect => `rewrite ^${redirect.from}$ $scheme://$http_host${redirect.to} redirect;`).join("\n");
        fs.writeFileSync(REDIRECT_CONFIG_PATH, redirectConfig);
        console.log("redirect config updated");

    } catch (err) {
        console.error("failed to update nginx config")
        console.error(err.message)
    }
}

const updateNginxConfig = async function(){
    await updateRedirectConfig();
    await updateRouteConfig();
    restartNginx();
    console.log("nginx config updated")
}

app.post("/update-config", async (req, res) => {
    console.log("webhook received, updating nginx");
    updateNginxConfig();
    res.json({ status: "nginx reloaded"});
})

updateNginxConfig().catch((err) => {
    console.error("failed to update initial config");
    console.error(err.message);

});

app.listen(PORT, () => console.log(`sidecar server listening on port ${PORT}`))
