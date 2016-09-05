
require "conf"
game_setup!


require "lib.console"
require "lib.loader"
require "lib.math"
require "lib.world"
require "lib.sprites"


ReqAll "behaviours"
ReqAll "species"


love.load = ()->
   player = NewPlayer!


love.update = (dt)->
   UpdateWorld dt


love.draw = ()->
   love.graphics.scale game_conf().upscale
   trace.draw!
   DrawWorld!


player\Jump!
if a != b
  if b == c
    if ((b+c) != "" and (b+c) != 0) 
      print d

c = (player or {}).b


