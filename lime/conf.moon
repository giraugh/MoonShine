export game_conf, game_setup
love.conf = (t)->
   t.identity = "Lime"
   t.window.title = "Lime"
   t.window.msaa = 0 --disable antialiasing

game_conf = ->
   t = {}
   t.upscale = 2, 2
   return t

game_setup = ->
   love.graphics.setDefaultFilter "nearest"
