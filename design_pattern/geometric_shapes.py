class GeometricShapes:

    def __init__(self):
        pass

    def get_area(self):
        pass
        
    def get_perimeter(self):
        pass
    
    @staticmethod
    def show_text(name, shape, *args, **kwargs):
        data = ''.join([f"{_key} = {_val} " for _key, _val in kwargs.items() if _key !='self'])
        print(f"Finding {name} of {shape} with {data}")

    def to_string(self):
        return [f"{_key}: {_val}" for _key, _val in self.locals.items()] + [f"area: {self.get_area():.1f}", f"perimeter: {self.get_perimeter():.1f}"]


class Rectangle(GeometricShapes):

    def __init__(self, length, width):
        self.locals = locals()
        self.locals.pop('self')
        self.length = length
        self.width = width

    def __str__(self):
        return "rectangle"

    def get_area(self):
        self.show_text(__class__.__name__, 'area', **self.locals)
        return (self.length + self.width) * 2

    def get_perimeter(self):
        self.show_text(__class__.__name__, 'perimeter', **self.locals)
        return self.length * self.width


class Square(Rectangle):

    def __init__(self, side):
        self.locals = locals()
        self.locals.pop('self')
        self.length = side
        self.width = side
    
    def __str__(self):
        return "square"


class Circle(GeometricShapes):

    def __init__(self, radius):
        self.locals = locals()
        self.locals.pop('self')
        self.radius = radius

    def __str__(self):
        return "circle"

    def get_area(self):
        self.show_text(__class__.__name__, 'area', **self.locals)
        return self.radius * 2 * 3.14
        
    def get_perimeter(self):
        self.show_text(__class__.__name__, 'perimeter', **self.locals)
        return self.radius ** 2 * 3.14

print("-"*10, "Rectangle","-"*10)
rectangle = Rectangle(2.0, 3.0)
print(rectangle.get_area())
print(rectangle.get_perimeter())
print(rectangle.to_string())

square = Square(2.0)
print("-"*10, "Square","-"*10)
print(square.get_area())
print(square.get_perimeter())
print(square.to_string())

circle = Circle(3.0)
print("-"*10, "Circle","-"*10)
print(circle.get_area())
print(circle.get_perimeter())
print(circle.to_string())
