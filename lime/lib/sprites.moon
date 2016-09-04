--Utility to load sprites and associate quads
export ^

class Image
   new: (@file, x, y, width, height)=>
      @img = love.graphics.newImage @file
      @usingQuads = (x != nil)
      if @usingQuads
         @quad = love.graphics.newQuad x, y, width, height, @img\getDimensions!
   draw: (x, y)=>
      if @usingQuads
         love.graphics.draw @img, @quad, x, y
      else
         love.graphics.draw @img, x, y


class Sprite
   new: (delay=100)=>
      @images = {}
      @frameIndex = 1
      @timer = 0
      @timerMax = delay --in milliseconds
   get: =>
      return @images[@frameIndex]
   update: (dt)=>
      @timer += dt*1000 --dt is in seconds so we multiply
      if @timer >= @timerMax
         @timer = 0
         @frameIndex += 1
         if @frameIndex > #@images
            @frameIndex = 1
   addFrame: (image)=>
      @images[#@images+1] = image
   fromTilesheet: (file, frameAmount, frameWidth, frameHeight, vertical) =>
      for i = 1, frameAmount
         if not vertical
            @addFrame Image file, (i-1)*frameWidth, 0, frameWidth, frameHeight
         else
            @addFrame Image file, 0, (i-1)*frameHeight, frameWidth, frameHeight

StaticSprite = (file)->
   sp = Sprite!
   sp\addFrame Image file
   return sp

AnimatedSprite = (file, frameAmount, frameWidth, frameHeight, vertical)->
   sp = Sprite!
   sp\fromTilesheet file, frameAmount, frameWidth, frameHeight, vertical
   return sp
