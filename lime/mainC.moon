--Setup Game
require "conf"
game_setup!

--Require Library
require "lib.console"
require "lib.loader"
require "lib.math"
require "lib.world"
require "lib.sprites"

--Dynamic requires
ReqAll "behaviours"
ReqAll "species"

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


