import sys
import math

# Auto-generated code below aims at helping you parse
# the standard input according to the problem statement.

#Read inputs.
surface_n = int(input())
for i in range(surface_n):
    land_x, land_y = map(int, input().split())

# Game loop
while True:
    x, y, h_speed, v_speed, fuel, rotate, power = map(int, input().split())

    if v_speed <= -44 and power < 4:
        power += 1
    print(0, power)
