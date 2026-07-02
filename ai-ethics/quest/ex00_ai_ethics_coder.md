# Part A: The Critical Distinction

# HONEST REFLECTION ON AI USE

# How have you used AI for coding so far? 
   - I have used AI for almost all coding task but i try to do it myself first before i amplify with AI.

# Do you usually try problems yourself first? 
   - Yes. I always observe the 15-20 minutes rules to test my knowledge on the challendge before i deepening my understanding with AI.

# Can you explain code you've submitted without AI's help? 
   - Yes, i can explain the basic of the code wihout AI help.

# What would happen if AI was unavailable during an exam or job test? 
- I am not expected to use AI during exam or job test. So therefore, if AI become unavailable, i will provide all answer to the exam or job test questions.. 


# Identify your current pattern:
# MY CURRENT PATTERN is Learner B: uses AI as a learning amplifier.
- Try to solve problem first
- brainstorm with AI for deep learning
- Break task into smaller parts for better understanding
- Undertake task with confidence

# Write a short paragraph describing where you are today and what type of learner you want to become.

For now, i am a learner who uses AI to amplify learning but with strong belief to be a learner who is independent of AI in the nearest future. A learner who sees or ideat with AI for deeper learning and for
- Enhance efficiency
- improve understanding of concept
- Optimizing code built
- simplifying learning



# Part B: The Wrong Way vs. The Right Way

# Track B — The Right Way (DO THIS): Step 1: Attempt independently

## 1. Write pseudocode for a palindrome check

MY PSUEDOCODE

FUNCTION is_palindrome(ogba)

1. SET reverseogba = reverse of ogba

2. IF ogba = reversedogba THEN

3. RETURN True

4. ELSE

5. RETURN False

6. END IF

7. END FUNCTION

 
## 2. Implement your solution in Python

# PYTHON IMPLEMENTATION 


def is_palindrome(ogba):
    """
    Checks if a string is a palindrome.
    Ignores case, spaces, and punctuation.
    """
    # Step 1: Convert to lowercase
    ogba = ogba.lower()

    # Step 2: Remove non-alphanumeric characters
    cleaned = ""
    for char in ogba:
        if char.isalnum():  # keep only letters and numbers
            cleaned += char

    # Step 3: Reverse the cleaned string
    reversed_cleaned = cleaned[::-1]

    # Step 4: Compare
    return cleaned == reversed_cleaned

Test_cases    
print(is_palindrome("racecar"))                     # True
print(is_palindrome("hello"))                       # False
print(is_palindrome("A man a plan a canal Panama")) # True


FULL IMPLEMENTATION WITH PYTHON TEST CASES

def is_palindrome(ogba):
    """
    def is_palindrome(ogba):
    """
    Checks if a string is a palindrome.
    Ignores case, spaces, and punctuation.
    """
    # Step 1: Convert to lowercase
    ogba = ogba.lower()

    # Step 2: Remove non-alphanumeric characters
    cleaned = ""
    for char in ogba:
        if char.isalnum():  # keep only letters and numbers
            cleaned += char

    # Step 3: Reverse the cleaned string
    reversed_cleaned = cleaned[::-1]

    # Step 4: Compare
    return cleaned == reversed_cleaned

test_cases
print(is_palindrome("racecar"))                     # True
print(is_palindrome("hello"))                       # False
print(is_palindrome("A man a plan a canal Panama")) # True



# Strategic AI use After i have my working solution, ask AI:

"I wrote a function to check palindromes. Here's my approach: [paste YOUR code]

# MODIFICATION ANALYSIS

    #What's the time complexity?
	 O(n)

    #What edge cases am I missing?
	 Punctuation, Empty string, only spaces, numbers mixed wit letters, non-string input 

   # Alternatives and trade-offs?

    def is_palindrome(ogba):

    ogba = ogba.lower().replace(" ", "")
    
    left = 0
    right = len(text) - 1

    while left < right:
        if text[left] != text[right]:
            return False, right
        left += 1
        right -= 1

    return True


  # How does it perform on very long strings?"
- Uses constant extra space
- More memory efficient
- Better for very large inputs



# Reflection

  #  What did you learn by struggling first?
     - By strugging first, i think deeper and i am confident i can solve the problem.

 #   How is your understanding different than if you'd just asked for the solution?
     - I understand the code Function, Conditional statement, Syntax and the workflow.

 # Can you now implement similar functions (reverse a string, find duplicates) without AI?
     - Yes, i can implement similar function (reverse a string)

 #   What mental model did you build?
     - This is the mental code i built:

def uche(text):
    text = text.lower().replace(" ", "")
    
    left = 0
    right = len(text) - 1

    while left < right:
        if text[left] != text[right]:
            return False
        left += 1
        right -= 1

    return True




    # Part D: The Fairness Contract

   # I will use AI when:

    - After I've attempted a problem for at least 20 minutes
    - To understand why my solution works/doesn't
    - To explore alternatives after I have a working solution
    - To improve my codebase
    - To learn deeper
    - For code efficiency
    

   # I will NOT use AI when:

   - I haven't made an attempt
   - I'm taking exam or test
   - I am learning fundamentals
   - Learning how code works
   

#    I know I'm using AI fairly when:

   - I can explain my code clearly
   - I feel more confident 
   - I could correct syntax errors
   - I am ideating and not depending on AI 
   - I didn't copy AI code directly
   
# Agbo Ugbala Sunday
# 16/02/2026



 # Part E: Real-World Scenario Analysis 

# Interview: "Explain how you'd implement a caching system." If you always relied on AI, can you answer?
  - If i always depend on AI, I will freeze, so i cannot answer.

# Production bug at 2 AM: AI is unavailable. Can you debug code you don't fully understand?
  - i can't fix code i don't fully understand. 

# New tech with little documentation: If you never learned to read docs and experiment, what happens?
  - Learning new tools requires reading documentation, Therefore, i cannot learn the tool.

# Write a paragraph: How does using AI fairly now prepare you for these scenarios?
  - Using AI fairly now helps me stay independent later or in the near future. Moreso, responsible AI use today will build my confidence an competence ensuring I can operate confidently even when AI is unavailable.
  

# Part F: Building Irreplaceable Skills

Rate yourself 1-5 and write an improvement plan for your lowest area:

 #         Skills                         RATING
    Problem decomposition                   4/5

    Systems thinking                        4/5

    Critical evaluation                     4/5

    Debugging mindset                       3/5

    Conceptual understanding (the "why")    4/5

# DEBUGGING MINDSET
   A debugging mindset is a way of thinking that focuses on finding, understanding, and fixing problems whether in code, systems, or even real life situations. A debugging mindset treats errors as clues, not failures.

# Action plan: 
  3 specific actions this week to improve it without outsourcing thinking to AI.
  
  1. Reproduce the issue consistently: I desire to be consistent in my daily routines in 
     order to attain a high level of understanding thereof.
  2. Isolate the smallest failing part: I desire to inproved my findings to enable me 
     known and understanding the bottleneck in problem solving.
  3. Test one variable at a time: To debug better means, to test one variable after 
     another untill lasting solution is reached.

