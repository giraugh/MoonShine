export ^

SplitString = (inSplitPattern, outResults )=>

   if not outResults
      outResults = {}
   theStart = 1
   theSplitStart, theSplitEnd = string.find @, inSplitPattern, theStart
   while theSplitStart
      table.insert outResults, string.sub @, theStart, theSplitStart-1
      theStart = theSplitEnd + 1
      theSplitStart, theSplitEnd = string.find @, inSplitPattern, theStart
   table.insert outResults, string.sub @, theStart
   return outResults

recursiveEnumerate = (folder="", fileTree="")->
	lfs = love.filesystem
	filesTable = lfs.getDirectoryItems folder
	for v in *filesTable
		file = folder.."/"..v
		if lfs.isFile(file)
			fileTree = fileTree.."\n"..file
		elseif lfs.isDirectory file
			fileTree = fileTree.."\n"..file.." (DIR)"
			fileTree = recursiveEnumerate file, fileTree
	return fileTree

-- Moonscript
ReqAll = (dir)->
	files = recursiveEnumerate dir
	lines = SplitString files, "\n"
	for l in *lines
		if l != "" and l\gsub("lua", "") != l
			file = l\gsub ".lua", ""
			file = file\gsub "/", "%."
			require file
