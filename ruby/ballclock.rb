def main()
    if ARGV.length < 2
        puts "Too few arguments"
        return
    end

    clocks = Array.new()
    found = false

    for i in 0..ARGV.length-1
        v = ARGV[i].to_i
        if v == 0
            found = true
            break
        end
        if v < 27 || v > 127
            puts "arguments out of range. Should be within [27,127]."
            return
        end

        clocks << v
    end

    if !found
        puts "Missing end argument 0"
        return
    end

    while clocks.length > 0
        input_size = clocks.shift
        days = ballClockDays(input_size)
        print "#{input_size} ball cycle after #{days} days.\n"
    end
end

def ballClockDays(input_size)
    original = Array.new(input_size) {|i| i}
    queue = Array.new(input_size) {|i| i}
    min = Array.new()
    fmin = Array.new()
    hour = Array.new()
    delta = Array.new()
    hours = 0

    #manually find delta pattern for first 12 hours.
    begin
        v = queue.shift
        #print "queue #{queue.length} min #{min.length} fmin #{fmin.length} hour #{hour.length} #{hours} #{hours/24}\n"
        if min.length == 4
            queue << min[3] << min[2] << min[1] << min[0]
            min.clear
            if fmin.length == 11
                queue << fmin[10] << fmin[9] << fmin[8] << fmin[7] << fmin[6] << fmin[5] << fmin[4] << fmin[3] << fmin[2] << fmin[1] << fmin[0]
                fmin.clear
                if hour.length == 11
                    queue << hour[10] << hour[9] << hour[8] << hour[7] << hour[6] << hour[5] << hour[4] << hour[3] << hour[2] << hour[1] << hour[0] << v
                    #initialize delta array, every 12 hours the pattern repeats.
                    delta = Array.new(queue)
                else
                    hour << v
                end
                hours+=1
            else
                fmin << v
            end
        else
            min << v
        end
    end while hours < 12

    begin
        queue = queue.map.with_index{|x,i| queue[delta[i]]}
        hours += 12
    end while original != queue

    return hours / 24
end

main
