--Setup Game
require !STR!2!STR!
game_setup!

--Require Library
require !STR!3!STR!
require !STR!4!STR!
require !STR!5!STR!
require !STR!6!STR!
require !STR!7!STR!

--Dynamic requires
ReqAll !STR!0!STR!
ReqAll !STR!1!STR!

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


