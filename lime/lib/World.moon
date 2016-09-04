export ^ --export propererly named i.e capital first letter

--Behaviours World, stores behaviours
BehavioursWorld = {}
GetBehaviour = (name)->
   for a in *BehavioursWorld
      if a.name == name
         return a
   return nil

--Species World, stores species
SpeciesWorld = {}
GetSpecies = (name)->
   for a in *SpeciesWorld
      if a.name == name
         return a
   return nil

--Entity World, stores entities
EntityWorld = {}

NewSpecies = (name)->
   sp = {}
   sp.name = name
   sp.behaviours = {}
   SpeciesWorld[#SpeciesWorld+1] = sp

ApplyBehaviour = (behaviour, species)->
   sp = GetSpecies species
   sp.behaviours[#sp.behaviours+1] = GetBehaviour behaviour

NewBehaviour = (name, spawn, update, draw)->
   be = {}
   be.name = name
   be.spawn = spawn
   be.update = update
   be.draw = draw
   BehavioursWorld[#BehavioursWorld+1] = be

SpawnEntity = (species, config)->
   sp = GetSpecies species
   en = {}
   en.species = sp.name
   en.config = config
   en.behaviours = sp.behaviours
   for behaviour in *en.behaviours
      if behaviour.spawn != nil
         behaviour.spawn en
   EntityWorld[#EntityWorld+1] = en
   return en

UpdateWorld = (dt)->
   for entity in *EntityWorld
      for behaviour in *entity.behaviours
         if behaviour.update != nil
            behaviour.update entity, dt

DrawWorld = ->
   for entity in *EntityWorld
      for behaviour in *entity.behaviours
         if behaviour.draw != nil
            behaviour.draw entity
