#pragma once

#include <iostream>
#include <vector>

struct OpCode {
    public:
        const  static  int  UNKNOWN   =  0;
        const  static  int  ADD       =  1;
        const  static  int  MULTIPLY  =  2;
        const  static  int  FINISH    =  99;
};

class IntCode {
    private:
        int pc;
        std::vector<int> memory;

        void add();
        void multiply();
        void finish();

    public:
        IntCode();
        explicit IntCode(const std::string& program);

        void set_program(const std::string& program);

        int execute();
        void step();
        bool is_finished();

        void set_memory(int index, int value);
        int get_memory(int index);

        void dump_memory(std::ostream& out);
};
