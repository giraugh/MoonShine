





print "hello", "world"

a = "hello world \n 'put in apostrophes' here is an escaped quote: \" and here is an escaped single-quote: \'"
b = 'hello world \n "put in quotes" here is an escaped quote: \" and here is an escaped single-quote: \''

print "this shouldnt be compiled AT ALL: :: is isnt a?"

player\jump! 
player\move fast 

if 1 == 1 
  print "obvious much?"

if 1 != 2 
  print "*sigh*"

if (myAwesomeVar != "" and myAwesomeVar != 0) 
  print "myAwesomeVar is #{myAwesomeVar + '"'}" 

(player or {}).position\addX 3 

print 'i am a single quote string'

myArr[#myArr+1] = 3  
myArr[#myArr+2] = 3
myArr[1+#myArr] = 1
a.myArr[#a.myArr+1] = 4
myArr[#myArr+1] = 1 
@myArr[#@myArr+1] = 3



if type(@) == "string"
  @ = @\gsub "\n", "\\n"
  if (string.match(@\gsub(@, "[^'\"]", ""), '$')) 
    return "'#{@}'"
  return "'#{@\gsub '"', '\\"'}'" 
