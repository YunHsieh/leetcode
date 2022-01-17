# Geometric Shapes

## Question
The area and perimeter of different geometric shapes like rectangle, circle, or square are
calculated using different mathematical formulae.
Create the following three classes:

1. The Rectangle class should implement the Shape interface. It should have two class variables of
float type, length and, width. Also, it should implement the following methods:
    - Rectangle(float new_length, float new_width): sets the class variables, length = new_length and
width = new_width respectively.
    - float getArea(): returns the result of length x width, the area of the rectangle. Also prints
"Finding area of rectangle with length = flength] and width = fwidth;", where flength]  and
fwidth] are the respective values of class variables length and width.
    - float getPerimeter0: returns the result of 2 x (length + width), the perimeter of the rectangle.
Also prints "Finding perimeter of rectangle with length = flength] and width = fwidth]"
    - String toString(): returns the string "Rectangle = [length: flength;, width: fwidth;, area: fareas,
perimeter: fperimeter]".
Note: flength] and (width: respectively represent the class variables length and width. farea] and
fperimeter] respectively represent the area and perimeter of the rectangle. For example, given
that length = 2 and width = 3, calling the method toString(will return "Rectangle = [length: 2.0,
width: 3.0, area: 6.0, perimeter: 10.0]'.

2. The Square class that should inherit the Rectangle class and should implement the following
methods:
    - Square(float side): sets the variables of Rectangle class, length = side and width =
side respectively.
    - float getArea(): returns the result of length x width that denotes the area of the square. Also
prints "Finding area of square with side = flength]"
float getPerimeterO: returns the result of 4 x /ength that denotes the perimeter of the square.
Also prints "Finding perimeter of square with side = flength]".
    - String toString(): returns the string 'Square = [side: flength), area: farea], perimeter:
iperimeterg]".
Note: flength) represents the variable length of Rectangle class. farea] and fperimeters
respectively represent the area and perimeter of the square. For example, given that side = 2,
calling the method toString(will return "Square = [side: 2.0, area: 4.0, perimeter: 8.0J'.

3. The Circle class that implements the Shape interface. It should have a class variable, radius.
Also, it should implement the following methods:
    - Circle(float new_ radius): sets the float  type variable of Circle class, radius = new_radius.
    - float getArea(): returns the result of 3.14 x radius x radius that denotes the area of the circle.
Also prints "Finding area of circle with radius = fradius]"
    - float getPerimeter(: returns the result of 6.28 x radius that denotes the perimeter of the circle.
Also prints "Finding perimeter of circle with radius = fradius]".
    -  String toString(: returns the string "Circle = [radius: fradius), area: fareas, perimeter:
iperimeter ]" .

Note: fradius; represents the variable radius of Circle class. farea] and fperimeter] respectively
represent the area and perimeter of the circle. For example, given that radius = 1, calling the
method toString(will return "Circle = [radius: 1.0, area: 3.14, perimeter: 6.28]'.


## My code Output
---------- Rectangle ----------
Finding Rectangle of area with length = 2.0 width = 3.0 
10.0
Finding Rectangle of perimeter with length = 2.0 width = 3.0 
6.0
Finding Rectangle of area with length = 2.0 width = 3.0 
Finding Rectangle of perimeter with length = 2.0 width = 3.0 
['length: 2.0', 'width: 3.0', 'area: 10.0', 'perimeter: 6.0']
---------- Square ----------
Finding Rectangle of area with side = 2.0 
8.0
Finding Rectangle of perimeter with side = 2.0 
4.0
Finding Rectangle of area with side = 2.0 
Finding Rectangle of perimeter with side = 2.0 
['side: 2.0', 'area: 8.0', 'perimeter: 4.0']
---------- Circle ----------
Finding Circle of area with radius = 3.0 
18.84
Finding Circle of perimeter with radius = 3.0 
28.26
Finding Circle of area with radius = 3.0 
Finding Circle of perimeter with radius = 3.0 
['radius: 3.0', 'area: 18.8', 'perimeter: 28.3']
