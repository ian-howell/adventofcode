import re


def main():
    x_min, x_max, y_min, y_max = get_target_area()
    count = 0
    for orig_dy in range(y_min, 1000):
        for orig_dx in range(x_max+1):
            x_too_high = False
            x, y = 0, 0
            dx, dy = orig_dx, orig_dy
            while True:
                x += dx
                y += dy

                if x > x_max:
                    x_too_high = True
                    break
                if y < y_min:
                    break

                if dx > 0:
                    dx -= 1
                elif dx < 0:
                    dx += 1
                dy -= 1

                if (x_min <= x <= x_max) and (y_min <= y <= y_max):
                    count += 1
                    break
            if x_too_high and y > y_max:
                break

    print(count)


def get_target_area():
    input_re = 'target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)'
    matches = re.search(input_re, input())
    x_min, x_max = int(matches.group(1)), int(matches.group(2))
    y_min, y_max = int(matches.group(3)), int(matches.group(4))
    return x_min, x_max, y_min, y_max


main()
