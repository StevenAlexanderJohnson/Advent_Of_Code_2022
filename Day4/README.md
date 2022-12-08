#  Day 4

## Prompt
The elves must clear space to get the rest of the provisions off the boats. Groups are given sections to clear and quickly realize that some groups have some overlaps. You are to determine the number of overlaps.

### Example input
```Text
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
```

# Part 1
First, they only really care about the groups that are completely overlapped by another group. And example of this would be when group A is assigned sections 2-8 and group B is assigned sections 4-6. All of group B's work would be completed when group A cleared their sections.

The assignment is to find how many of groups are completely overlapped.

My approach was to split the line on the comma to separate the two groups. I then split the resulting strings of the dash to get the "From" and "To" sections.

The last step was to check if group one's assignment was completely inside group two's assignment or vice versa and increment a count variable if so.

# Part 2
In the second part they cared not only about the complete overlaps, but if the two groups overlapped at all. The process for this solution is very similar to that of part one, but in this solution, I checked if the From point of assignment two was between the From and To point of assignment one and vice versa.

One error I made was to check the From points for each line which resulted in an answer higher than expected. If the From point of an assignment was equal to From/To point of the other it would have been counted twice. To combat this I simply added an else before the second comparison so it would only be counted once.