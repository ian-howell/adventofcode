import re


def main():
    best_y = 0
    x_min, x_max, y_min, y_max = get_target_area()
    for orig_dy in range(int(1e3)):
        for orig_dx in range(int(1e3)):
            x_too_high = False
            highest_y_this_time = 0
            x, y = 0, 0
            dx, dy = orig_dx, orig_dy
            # print(f"Trying {dx}, {dy}")
            while True:
                x += dx
                y += dy
                highest_y_this_time = max(highest_y_this_time, y)
                # print(f"Velocity: ({dx}, {dy})   Position: ({x}, {y})")

                if x > x_max:
                    # print("Overshot...")
                    x_too_high = True
                    break
                if y < y_min:
                    # print("Probe fell out of range...")
                    break

                if dx > 0:
                    dx -= 1
                elif dx < 0:
                    dx += 1
                dy -= 1

                if (x_min <= x <= x_max) and (y_min <= y <= y_max):
                    best_y = max(highest_y_this_time, best_y)
                    # print(f"Found a new best: {best_y}")
                    break
            if x_too_high:
                break

    print(best_y)


def get_target_area():
    input_re = 'target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)'
    matches = re.search(input_re, input())
    x_min, x_max = int(matches.group(1)), int(matches.group(2))
    y_min, y_max = int(matches.group(3)), int(matches.group(4))
    return x_min, x_max, y_min, y_max


main()
