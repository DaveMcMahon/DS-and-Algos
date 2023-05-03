# DS-and-Algos
Space for Data Structures and Algorithm type programs in a variety of languages.

As I progress in my learning of DS & Algos, I will update this readme with notes and points of interest, specifically to recall
certain topics at a glance.

The book I am referencing is 'Grokking Algorithms' authored by Aditya Y. Bhargava.

Algorithm
    - A set of instructions for acomplishing a task.

Binary Search
    - Takes an input of a sorted list of elements and returns the element you're searching for.
    - Returns null if that elements doesn't exist
    - Simple Search eliminates only one element at a time (not efficient).
    - Binary Search goes with the middle element, and eliminates half the remaining numbers, every time.
    Example:
        - Guess a number between 1 and 100.
        - Binary Search can guess the number in 7 steps by halving each result set and checking if the element is
            higher or lower.
        - 100 -> 25 -> 13 -> 7 -> 4 -> 2 -> 1
    Run Time
        - log2n
            - How many 2's to get to n.
            - Opposite of exponentials
            - log2(100) = 7     (2 x 2 x 2 x 2 x 2 x 2 x 2)
            




