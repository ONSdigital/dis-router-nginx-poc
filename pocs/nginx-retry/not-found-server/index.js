const express = require("express");

const app = express();
const PORT = process.env.PORT || 5000;

app.use(express.json());

app.get("/economy/*", async (req, res) => {
    console.log(`received request to ${req.path}`);
    res.status(200).send("request proxied to the not-found-server");
})

app.get("/*", async (req, res) => {
    console.log(`received request to ${req.path}`);
    res.sendStatus(404);
})

app.listen(PORT, () => console.log(`not-found-server listening on port ${PORT}`))
