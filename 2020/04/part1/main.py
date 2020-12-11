def main():
    passports = get_all_passports()

    valid_passports = 0
    for passport in passports:
        if is_valid_passport(passport):
            valid_passports += 1
    print(valid_passports)


def get_all_passports():
    passports = []
    passport = set()
    while True:
        line = None
        try:
            line = input()
        except EOFError:
            break
        else:
            if line != "":
                passport |= {field[:3] for field in line.split()}
            else:
                passports += [passport]
                passport = set()

    return passports + [passport]


def is_valid_passport(passport, must_contain=['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']):
    for v in must_contain:
        if v not in passport:
            return False
    return True


if __name__ == "__main__":
    main()
