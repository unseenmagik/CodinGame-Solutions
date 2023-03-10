import sys
import math

# The while loop represents the game.
# Each iteration represents a turn of the game
# where you are given inputs (the heights of the mountains)
# and where you have to print an output (the index of the mountain to fire on)
# The inputs you are given are automatically updated according to your last actions.


#Game loop.
while True:
    max = 0
    maxIndex = -1
    
    for i in range(8):
        #Read inputs.
        mountainH = int(input())
        
        #Set highest mountain.
        if mountainH > max:
            max = mountainH
            maxIndex = i

    #Output highest mountain.
    print(maxIndex)