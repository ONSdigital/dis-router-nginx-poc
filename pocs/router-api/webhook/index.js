const chokidar = require("chokidar");
const axios = require("axios");

console.log("watching database")

chokidar.watch("../data/db.json").on("change", () => {
    console.log("database has changed")
    axios.get("http://router:80/webhook/redirect-config")
        .then(() => console.log("webhook sent"))
        .catch(err => console.error("webhook error", err));
});
