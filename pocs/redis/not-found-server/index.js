const express = require("express");

const app = express();
const PORT = process.env.PORT || 5000;

app.use(express.json());

app.get("/releases/wagtailrelease", async (req, res) => {
    console.log(`received request to ${req.path}`);
    res.status(200).send("here's the wagtail release");
})

app.get("/releases/babbagerelease", async (req, res) => {
    console.log(`received request to ${req.path}`);
    res.sendStatus(404);
})

app.get("/consumer-price-inflation/*", async (req, res) => {
    console.log(`received request to ${req.path}`);
    res.status(200).send("request received at fake wagtail");
})

app.get("/*", async (req, res) => {
    console.log(`received request to ${req.path}`);
    res.sendStatus(404);
})

app.listen(PORT, () => console.log(`not-found-server listening on port ${PORT}`))
