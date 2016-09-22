# MoonShine
A precompiler for the moonscript precompiler.

### Major differences from moonscript
* block comments ```---``` to start and end
* use ```is``` and ```isnt``` instead of ```==``` and ```!=```
* use ```::``` instead of ```\```
* use the existential accessor ```?``` (like in coffescript) i.e ```scene?.player?.jumpHeight```
* easily append to tables, ```myTable[+]``` is the same as ```myTable[#myTable+1]```
* or access table lengths easily, ```myTable[##+2]``` is the same as ```myTable[#myTable+2]``` (you can also do ```[##]```)
* increment operator (like in every other programming language), ```++``` i.e ```myVar++```
* put two statements on one line, ```statement1 && statement2```

__Moonshine is just a precompiler for [moonscript](https://github.com/leafo/moonscript). You still need a copy of it.__

There is an atom grammar available for moonshine [here](https://atom.io/packages/language-moonshine). It is based off the [moonscript grammar](https://atom.io/packages/language-moonscript) by OttoRobba.
