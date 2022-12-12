
class State:
    """Class to save the possible states in."""
    X_MIN = 0
    Y_MIN = 0
    X_MAX = 170
    Y_MAX = 40

    def __init__(self, x, y, elevation, parent, depth):
        self.x = x
        self.y = y
        self.elevation = elevation
        self._parent = parent
        self._depth = depth

    def parent(self):
        """Return the previous state."""
        return self._parent

    def depth(self):
        """Return node depth."""
        return self._depth

    def is_solved(self):
        """Return True if the current state is the solution (if Khun Phaen is at the goal coordinates)."""
        return self.elevation == 'E'

    def can_move_to(self, to):
        if to == 'E':
            return self.elevation == 'z'
        return (ord(to) - ord(self.elevation)) <= 1

    # def print_solution(self, max_depth):
    #     """Recursively print the solution to the puzzle. Only call once the goal has been reached."""
    #     if self._parent is not None:
    #         self._parent.print_solution(max_depth)
    #         # Overwrite previous step
    #         print('\033[24A', end='\r')
    #     # Print each step with a sleep
    #     print(f"Step {self._depth:>3}/{max_depth}", end="\n\n")
    #     print(self)
