#include <iostream>
#include <cassert>
#include <queue>
#include <stack>
#include <stdlib.h>

int main(int argc, char *argv[])
{
    if(argc < 2)
    {
        std::cerr << "Too few arguments" << std::endl;
        return 1;
    }

    std::queue<int> clocks;

    for(int i = 1; i < argc; i++)
    {
        int v = atoi(argv[i]);
        if(!v) //stopping condition
            break;

        if(v < 27 || v > 127)
        {
            std::cerr << "arguments out of range. Should be within [27,127]." << std::endl;
            return 1;
        }

        clocks.push(v);
    }

    while(clocks.size())
    //for(int j=27; j < 127; j++)
    {
        int input_size = clocks.front();
        clocks.pop();
        std::queue<int> original;
        std::queue<int> queue;
        std::stack<int> min, fmin, hour;

        for(int i = 0; i < input_size;i ++)
        {
            original.push(i);
            queue.push(i);
        }

        int hours = 0;
        do
        {
            //assert(queue.size());

            int v = queue.front();
            queue.pop();
            //std::cout << "queue=" << queue.size() << " min=" << min.size() << " fmin=" << fmin.size() << " hour=" << hour.size() << std::endl;
            if(min.size() == 4)
            {
                while(min.size())
                {
                    queue.push(min.top());
                    min.pop();
                }

                if(fmin.size() == 11)
                {
                    while(fmin.size())
                    {
                        queue.push(fmin.top());
                        fmin.pop();
                    }

                    if(hour.size() == 11)
                    {
                        while(hour.size())
                        {
                            queue.push(hour.top());
                            hour.pop();
                        }
                        queue.push(v);
                    }
                    else
                        hour.push(v);

                    //std::cout << "Hours=" << hours << std::endl;
                    hours++;
                }
                else
                    fmin.push(v);
            }
            else
                min.push(v);
        } while(original != queue);

        //std::cout << input_size << " ball cycle after " << hours/24 << " days." << std::endl;
        std::cout << input_size << " " << hours/24 << std::endl;
    }

   return 0;
}
