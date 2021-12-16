# For this problem, I will define a packet as a 3-tuple, where the first item
# is the packet's version, the second is its type, and the third is its value.
#
# The value for a packet shall either be an integer literal, or a list of
# sub-packets.
def main():
    bin_string = get_bin_string()
    # print(f"Initial: {bin_string=}")
    packet = parse(bin_string)[0]
    # print_packet(packet)
    print(sum_versions(packet))


def get_bin_string():
    hex_val = int(input(), base=16)
    bin_string = bin(hex_val)[2:]
    needed_zeros = 4 - (len(bin_string) % 4)
    if needed_zeros < 4:
        bin_string = '0'*needed_zeros + bin_string
    return bin_string


def parse(bin_string):
    packet_version = int(bin_string[:3], base=2)
    bin_string = bin_string[3:]
    packet_type = int(bin_string[:3], base=2)
    bin_string = bin_string[3:]

    # print(f"{packet_version=}, {packet_type=}, {bin_string=}")
    if packet_type == 4:
        value, bits_read = parse_literal(bin_string)
        packet = (packet_version, packet_type, value)
        # print(f"literal {packet=}")
        # print()
        return packet, bits_read + 6

    sub_packets = []
    length_type = bin_string[0]
    bin_string = bin_string[1:]
    if length_type == '0':
        length = int(bin_string[:15], base=2)
        bin_string = bin_string[15:]
        # print(f"type 0: {length=}")
        bits_read_total = 0
        while bits_read_total != length:
            sub_packet, bits_read = parse(bin_string)
            bin_string = bin_string[bits_read:]
            bits_read_total += bits_read
            # print(f"0: {bits_read=}, {bits_read_total=}, {bin_string=}")
            sub_packets.append(sub_packet)
        # print()
        return (packet_version, packet_type, sub_packets), bits_read_total + 6 + 1 + 15
    else:
        segments = int(bin_string[:11], base=2)
        bin_string = bin_string[11:]
        # print(f"type 1: {segments=}")
        bits_read_total = 0
        for _ in range(segments):
            sub_packet, bits_read = parse(bin_string)
            bin_string = bin_string[bits_read:]
            bits_read_total += bits_read
            # print(f"1: {bits_read=}, {bits_read_total=}, {bin_string=}")
            sub_packets.append(sub_packet)
        # print()
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


def sum_versions(packet):
    # print(f"{packet=}")
    total = packet[0]
    if packet[1] == 4:
        return total
    for sub_packet in packet[2]:
        total += sum_versions(sub_packet)
    return total


def print_packet(packet, level=0):
    print(level*" " + f"- Version: {packet[0]}")
    print(level*" " + f"  Type:    {packet[1]}")
    if type(packet[2]) is list:
        for sub_packet in packet[2]:
            print_packet(sub_packet, level+2)
    else:
        print(level*" " + f"  Value:   {packet[2]}")

main()
