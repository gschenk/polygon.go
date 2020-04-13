# polyGo -- Convex Hull Polygon Area
Returns the area of the smallest convex polygon that encloses a point cloud. Go final version of an algorithm 
prototyped in ES6 Java Scipt for Node https://github.com/gschenk/polygon.node

## Algorithm Outline
Node of the complex hull polygon (CHP) are found through a gift-wrap algorithm with Akl-Toussaint heuristics.
A simple method is used to find polygon of three to eight edges that inscribes the CHP. This
polygon is henceforth called Akl-Toussaint polygon (ATP).

Starting at an arbitrary point of the ATP a gift wrap algorith is only used for points outside
the ATP. The gift wrap algorithm selects a point with an extreme angle it forms with a node and
a point that is also within the CHP polygon. The point identifies thus, a node of the
CHP, is then the starting point for the search for the next node point.

Outside points are identified by the sign of the determinant formed of an edge vector
with the vector from the respective node to the point. The value of this determinant implicitly
contains angle information and is used to identify the next CHP node.

We will go around counter clockwise the CHP polygon starting at point with the smallest x-coordinate value
(designated 'left').

The sum of the determinant of the CHP edge vectors and an arbitrary point within (geometric centre of the CHP)
is the area of the CHP.


## Useage
Pipe a JSON formated collection of x-y coordinates as a single line to STDIN. Example:
`[{"x": 1.1, "y":0.2},{"x": -1e-3, "y":-1}, ... ]`

The programme will return a scalar value (CHP area) to STDOUT.

## maybe useful
convex hull
https://en.wikipedia.org/wiki/Convex_hull

Andrew's monotone chain algorithm
http://geomalgorithms.com/a10-_hull-1.html

Akl-Toussaint Heuristics
https://arxiv.org/abs/1304.2676
