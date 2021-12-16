# For this problem, I will define a packet as a 3-tuple, where the first item
# is the packet's version, the second is its type, and the third is its value.
#
# The value for a packet shall either be an integer literal, or a list of
# sub-packets.
def main():
    bin_string = get_bin_string()
    packet = parse(bin_string)[0]
    # print_packet(packet)
    print(calculate(packet))


def get_bin_string():
    raw_string = input()
    total_length = len(raw_string) * 4
    hex_val = int(raw_string, base=16)
    bin_string = bin(hex_val)[2:]
    needed_zeros = total_length - len(bin_string)
    bin_string = '0'*needed_zeros + bin_string
    return bin_string


def parse(bin_string):
    packet_version = int(bin_string[:3], base=2)
    bin_string = bin_string[3:]
    packet_type = int(bin_string[:3], base=2)
    bin_string = bin_string[3:]

    if packet_type == 4:
        value, bits_read = parse_literal(bin_string)
        packet = (packet_version, packet_type, value)
        return packet, bits_read + 6

    sub_packets = []
    length_type = bin_string[0]
    bin_string = bin_string[1:]
    if length_type == '0':
        length = int(bin_string[:15], base=2)
        bin_string = bin_string[15:]
        bits_read_total = 0
        while bits_read_total != length:
            sub_packet, bits_read = parse(bin_string)
            bin_string = bin_string[bits_read:]
            bits_read_total += bits_read
            sub_packets.append(sub_packet)
        return (packet_version, packet_type, sub_packets), bits_read_total + 6 + 1 + 15
    else:
        segments = int(bin_string[:11], base=2)
        bin_string = bin_string[11:]
        bits_read_total = 0
        for _ in range(segments):
            sub_packet, bits_read = parse(bin_string)
            bin_string = bin_string[bits_read:]
            bits_read_total += bits_read
            sub_packets.append(sub_packet)
        return (packet_version, packet_type, sub_packets), bits_read_total + 6 + 1 + 11


def parse_literal(bin_string):
    done = False
    value_string = ''
    chunks_read = 0
    while not done:
        done = bin_string[0] == '0'
        value_string += bin_string[1:5]
        bin_string = bin_string[5:]
        chunks_read += 1
    return int(value_string, base=2), chunks_read * 5


def print_packet(packet, level=0):
    print(level*" " + f"- Version: {packet[0]}")
    print(level*" " + f"  Type:    {packet[1]}")
    if type(packet[2]) is list:
        for sub_packet in packet[2]:
            print_packet(sub_packet, level+2)
    else:
        print(level*" " + f"  Value:   {packet[2]}")


def calculate(packet):
    if packet[1] == 4:
        return packet[2]
    if packet[1] == 0:
        # Packets with type ID 0 are sum packets - their value is the sum of
        # the values of their sub-packets. If they only have a single
        # sub-packet, their value is the value of the sub-packet.
        return sum(calculate(sub_packet) for sub_packet in packet[2])
    if packet[1] == 1:
        # Packets with type ID 1 are product packets - their value is the
        # result of multiplying together the values of their sub-packets. If
        # they only have a single sub-packet, their value is the value of the
        # sub-packet.
        return product(calculate(sub_packet) for sub_packet in packet[2])
    if packet[1] == 2:
        # Packets with type ID 2 are minimum packets - their value is the
        # minimum of the values of their sub-packets.
        return min(calculate(sub_packet) for sub_packet in packet[2])
    if packet[1] == 3:
        # Packets with type ID 3 are maximum packets - their value is the
        # maximum of the values of their sub-packets.
        return max(calculate(sub_packet) for sub_packet in packet[2])
    if packet[1] == 5:
        # Packets with type ID 5 are greater than packets - their value is 1 if
        # the value of the first sub-packet is greater than the value of the
        # second sub-packet; otherwise, their value is 0. These packets always
        # have exactly two sub-packets.
        if calculate(packet[2][0]) > calculate(packet[2][1]):
            return 1
        return 0
    if packet[1] == 6:
        # Packets with type ID 6 are less than packets - their value is 1 if
        # the value of the first sub-packet is less than the value of the
        # second sub-packet; otherwise, their value is 0. These packets always
        # have exactly two sub-packets.
        if calculate(packet[2][0]) < calculate(packet[2][1]):
            return 1
        return 0
    if packet[1] == 7:
        # Packets with type ID 7 are equal to packets - their value is 1 if the
        # value of the first sub-packet is equal to the value of the second
        # sub-packet; otherwise, their value is 0. These packets always have
        # exactly two sub-packets.
        if calculate(packet[2][0]) == calculate(packet[2][1]):
            return 1
        return 0


def product(iterable):
    total = 1
    for x in iterable:
        total *= x
    return total


main()
