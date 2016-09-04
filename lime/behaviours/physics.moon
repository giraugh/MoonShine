--Physics
--DOES: Velocity and collision
--DEPENDS: sprite

phySpawn = =>
   @velocity = Vector 0, 0

phyUpdate = (dt)=>
   @position += Vector @velocity.x*dt, @velocity.y*dt

NewBehaviour "physics", phySpawn, phyUpdate, nil
