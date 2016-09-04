--Sprite
--DOES: draws defined sprite
--DEPENDS: none

sprSpawn = =>
   @sprite = nil
   @position = Vector 0, 0
   @setPosition = (x, y)=>
      @position[0] = x
      @position[1] = y
   @AssignSprite = (sprite)=>
      @sprite = sprite

sprUpdate = (dt)=>
   if @sprite != nil
      @sprite\update dt

sprDraw = =>
   if @sprite != nil
      @sprite\get!\draw @position.x, @position.y

NewBehaviour "sprite", sprSpawn, sprUpdate, sprDraw
