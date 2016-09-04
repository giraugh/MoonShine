export trace, print
trace =
   visible: true
   texts: {}
   styles: {}
   styles:
      white:
         r: 255
         g: 255
         b: 255
      red:
         r: 255
         g: 127
         b: 127
      green:
         r: 191
         g: 255
         b: 127
      blue:
         r: 127
         g: 159
         b: 255
      default:
         r: 224
         g: 224
         b: 224

   count: 0
   limit: 32
   print: (text, style=trace.styles.default)->
      if trace.count > trace.limit
         table.remove trace.texts, 1
         table.remove trace.styles, 1
      else
         trace.count += 1
      trace.texts[trace.count] = text
      trace.styles[trace.count] = style
   draw: (x=16, y=16)->
      if trace.visible
         --for each text
         for i = 1,trace.count
            s = trace.styles[i]
            t = trace.texts[i]
            --check string type
            if type(t) == "table" then t = "<table>"
            if t == nil then t = "<nil>"
            --pad string
            t = string.rep("\n", i) .. t
            --determine outline
            if (s.r < 160) and (s.g < 160) and (s.b < 160)
      			love.graphics.setColor 255, 255, 255
      		else
      			love.graphics.setColor 0, 0, 0
            -- draw outline:
            love.graphics.print t, x + 1, y
            love.graphics.print t, x - 1, y
            love.graphics.print t, x, y + 1
            love.graphics.print t, x, y - 1
            -- draw color:
            love.graphics.setColor s.r, s.g, s.b
            love.graphics.print t, x, y

-- overload print function to trace.print
print = trace.print
