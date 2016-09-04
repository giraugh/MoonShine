export ^
-- Player
-- DOES: Moves and shoots

NewSpecies "player"
ApplyBehaviour "sprite", "player"
ApplyBehaviour "sprite-states", "player"
ApplyBehaviour "physics", "player"

NewPlayer = ->
   player = SpawnEntity "player"
   player\associateState "walk_left", AnimatedSprite "graphics/player/walk_left.png", 6, 24, 24, false
   player\associateState "walk_right", AnimatedSprite "graphics/player/walk_right.png", 6, 24, 24, false
   player\setState "walk_right"
   return player
