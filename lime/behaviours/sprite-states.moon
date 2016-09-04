--Sprite State Machine
--DOES: set up states and associate sprites
--DEPENDS: sprite

ssmSpawn = =>
   @spriteStates = {}
   @state = "default"
   @setState = (@state)=>
   @associateState = (state, image)=>
      @spriteStates[state] = image

ssmUpdate = =>
   if @spriteStates[@state] != nil
      @sprite = @spriteStates[@state]

NewBehaviour "sprite-states", ssmSpawn, ssmUpdate
