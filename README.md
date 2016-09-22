# MoonShine
A precompiler for the moonscript precompiler.

### Major differences from moonscript
* block comments ```---``` to start and end
* use ```is``` and ```isnt``` instead of ```==``` and ```!=```
* use ```::``` instead of ```\```
* use the existential accessor ```?``` (like in coffescript) i.e ```scene?.player?.jumpHeight```
* easily append to tables, ```myTable[+]``` is the same as ```myTable[#myTable+1]```
* or access table lengths easily, ```myTable[##+1]``` is the same as ```myTable[#myTable+1]```
* increment operator (like in every other programming language), ```++``` i.e ```myVar++```
* put two statements on one line, ```statement1 && statement2```

__Moonshine is just a precompiler for [moonscript](https://github.com/leafo/moonscript). You still need a copy of it.__

