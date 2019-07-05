var express = require("express")
var app = express()

//MongoDB
var mongoose = require("mongoose")
// mongodb://<user>:<pwd>@<host>:<port>/<database>
mongoose.connect('mongodb://localhost:27017/demo', {useNewUrlParser: true});
var db = mongoose.connection;

db.on('error', () => {
  console.log("MongoDB 连接异常")
})

var bodyParser = require("body-parser")
app.use(bodyParser.json())
app.use(bodyParser.urlencoded({extended: false}))

var position = require('./routes/position')
app.use("/position", position)

var port = 8888;
app.listen(port, () => {
  console.log("仓位记录管理服务运行中……")
})
