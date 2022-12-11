# Day 6

## Prompt
After finishing the preparations, you start your march to the first star fruit grove. On the way one of the elves hands you a broke handheld device for communication. 
To start talking to someone on the device, you have to first connect to their signal. To fix the device you have to add a subroutine to find the __start-of-packet__ 
marker. The elfs use a protocol that says that the marker is four characters long where no character is repeated. The input being a large string. The output for the 
problem was to return the index of the last character in the marker.

# Part 1
This is a classic problem. To do this I simply used a "sliding window". I created a map to track instances of each letter in the window, and once I found a window that had 
no repeating characters I would return the index of the last character from the window in the larger string. I was able to deduce that the character was repeated by checking 
if the character already existed in the map. I don't really care what the value of the map was only if the key existed.

# Part 2
This prompt was exactly the same but the window was now 14 characters long compared to the 4 characters in part 1. I simply changed the window size and it worked. 
I imagine that this problem had to do with optimization. Having a nested loop to check for repeating characters would be very slow. Compared to Part 1's window where 
there are 6 comparisons for each window, Part 2's window has 91 per window.