from roll_the_string import roll_the_string

def test_roll_the_string_scence1():
    result = roll_the_string('abz', [3,2,1])
    assert result == 'dda'


def test_roll_the_string_scence2():
    result = roll_the_string('ksdfjaljfdfdfad', [30,20,10])
    assert result == 'pxikoepnjhigidg'


def test_roll_the_string_scence3():
    result = roll_the_string('', [3,2,1])
    assert result == ''
