





print "hello", "world"

a = "hello world \n \'put in apostrophes\' here is an escaped quote: \" and here is an escaped single-quote: \'"
b = 'hello world \n \"put in quotes\" here is an escaped quote: \" and here is an escaped single-quote: \''

print "this shouldnt be compiled AT ALL: :: is isnt a?"

player\jump! 
player\move fast 

if 1 == 1 
  print "obvious much?"

if 1 != 2 
  print "*sigh*"

if (myAwesomeVar != "" and myAwesomeVar != 0) 
  print "myAwesomeVar is #{myAwesomeVar + }" 

(player or {}).position\addX 3 

print 'i am a single quote string'


if type(@) == "string"
  @ = @\gsub "\n", "\\n"
  if (string.match(@\gsub(@, "[^\"]", ""), '$')) 
    return "#{@}"
  return "#{@::gsub , \\}"
