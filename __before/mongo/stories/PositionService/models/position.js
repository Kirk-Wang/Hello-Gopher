var mongoose = require("mongoose")
var Schema = mongoose.Schema;

var PositionSchema = new Schema({
  account: {
    type: String,
    require: true,
  },
  stock: {
    type: String,
    require: true,
  },
  quantity: {
    type: Number,
    require: true,
  },
  price: {
    type: Number,
    require: true,
  },
})

module.exports = mongoose.model("Position", PositionSchema)

