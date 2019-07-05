var express = require('express')
var router = express.Router()

var positionController = require("../controllers/position")

router.post("/create", positionController.createPosition)
router.get("/:account", positionController.queryPosition)
router.put("/:id/update", positionController.updatePosition)
router.delete("/:id/delete", positionController.deletePosition)

module.exports = router;