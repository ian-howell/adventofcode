#include <iostream>
#include <sstream>
#include <vector>

#define atoi(x) ( x & 0xf )

std::vector<char> parse_tokens(const std::string& input);
long compute(const std::vector<char>& tokens, size_t& pos);

int main() {
    std::string line;
    std::string token;
    long total = 0;
    while (std::getline(std::cin, line)) {
        std::stringstream in(line);
        line.clear();
        while (in >> token) {
            line += token;
        }
        auto tokens = parse_tokens(line);
        size_t pos = 0;
        total += compute(tokens, pos);
    }
    std::cout << total << std::endl;
}

std::vector<char> parse_tokens(const std::string& input) {
    std::vector<char> tokens;
    std::string _input = input;

    while (!_input.empty()) {
        tokens.push_back(_input[0]);
        _input = _input.substr(1);
    }

    return tokens;
}

long compute(const std::vector<char>& tokens, size_t& pos) {
    if (pos == (tokens.size() - 1)) {
        return atoi(tokens.at(pos));
    }

    long val;
    if (tokens.at(pos) == '(') {
        pos += 1;
        val = compute(tokens, pos);
    } else {
        val = atoi(tokens.at(pos));
    }

    size_t i = pos + 1;
    while (i < tokens.size()) {
        if (tokens.at(i) == ')') {
            pos = i;
            return val;
        }

        if (tokens.at(i) == '+') {
            i++;
            if (tokens.at(i) == '(') {
                i++;
                val += compute(tokens, i);
            } else {
                val += atoi(tokens.at(i));
            }
        } else {
            i++;
            if (tokens.at(i) == '(') {
                i++;
                val *= compute(tokens, i);
            } else {
                val *= atoi(tokens.at(i));
            }
        }
        i++;
    }
    return val;
}
