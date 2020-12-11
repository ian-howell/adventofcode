def main():
    passports = get_all_passports()

    valid_passports = 0
    for passport in passports:
        if is_valid_passport(passport):
            valid_passports += 1
    print(valid_passports)


def get_all_passports():
    passports = []
    passport = {}
    while True:
        line = None
        try:
            line = input()
        except EOFError:
            break
        else:
            if line != "":
                passport.update({field[:3]: field[4:] for field in line.split()})
            else:
                passports += [passport]
                passport = {}

    return passports + [passport]


def is_valid_passport(passport):
    validation_funcs = {
            'byr': lambda x: 1920 <= int(x) <= 2002,
            'iyr': lambda x: 2010 <= int(x) <= 2020,
            'eyr': lambda x: 2020 <= int(x) <= 2030,
            'hgt': lambda x: (x.endswith('in') and 59 <= int(x[:-2]) <= 76) or (x.endswith('cm') and 150 <= int(x[:-2]) <= 193),
            'hcl': lambda x: len(x) == 7 and x[0] == '#' and all([y in '0123456789abcdef' for y in x[1:]]),
            'ecl': lambda x: x in {'amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'},
            'pid': lambda x: len(x) == 9 and all([y in '0123456789' for y in x]),
            }
    for v in validation_funcs:
        if v not in passport or not validation_funcs[v](passport[v]):
            return False
    return True

if __name__ == "__main__":
    main()
