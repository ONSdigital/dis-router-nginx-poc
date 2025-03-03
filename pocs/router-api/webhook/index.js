const chokidar = require("chokidar");
const axios = require("axios");

const WEB_HOOK_URL = process.env.WEB_HOOK_URL || "http://localhost:8080/webhook/redirect-config";

console.log("watching database")

chokidar.watch("../data/db.json").on("change", () => {
    console.log("database has changed")

    console.log(`firing webhook to ${WEB_HOOK_URL}`);
    axios.post(`${WEB_HOOK_URL}`)
        .then(() => console.log("webhook sent"))
        .catch(err => console.error("webhook error", err));
});
