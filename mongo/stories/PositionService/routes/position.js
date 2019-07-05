var express = require('express')
var router = express.Router()

var positionController = require("../controllers/position")

router.post("/create", positionController.createPosition)
router.post("/:account", positionController.queryPosition)
router.post("/:id/update", positionController.updatePosition)
router.post("/:id/delete", positionController.deletePosition)

module.exports = router;