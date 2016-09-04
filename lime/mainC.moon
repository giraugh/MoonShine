--Setup Game
require !STR!1!STR!
game_setup!

--Require Library
require !STR!1!STR!
require !STR!1!STR!
require !STR!1!STR!
require !STR!1!STR!
require !STR!1!STR!

--Dynamic requires
ReqAll !STR!0!STR!
ReqAll !STR!0!STR!

--Load Event
love.load = ()->
   player = NewPlayer!

--Update Event
love.update = (dt)->
   UpdateWorld dt

--Draw Event
love.draw = ()->
   love.graphics.scale game_conf().upscale
   trace.draw!
   DrawWorld!


