export ^

Vector = (x=0, y=0, z=0) ->
   r = x: x, y: y, z: z
   mt =
      __add: (a, b)->
         ret =
            x: a.x + b.x
            y: a.y + b.y
            z: a.z + b.z
         return ret
      __sub: (a, b)->
         ret =
            x: a.x - b.x
            y: a.y - b.y
            z: a.z - b.z
         return ret
      __div: (a, b)->
         ret =
            x: a.x / b.x
            y: a.y / b.y
            z: a.z / b.z
         return ret
      __mul: (a, b)->
         ret =
            x: a.x * b.x
            y: a.y * b.y
            z: a.z * b.z
         return ret
      __eq: (a, b)->
         return (a.x == b.x) and (a.y == b.y) and (a.z == b.z)
      __lt: (a, b)->
         return (a.x < b.x) and (a.y < b.y) and (a.z < b.z)
      __le: (a, b)->
         return (a.x <= b.x) and (a.y <= b.y) and (a.z <= b.z)
      __len: ->3
   setmetatable r, mt
   return r
