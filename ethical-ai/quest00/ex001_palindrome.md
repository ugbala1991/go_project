# PSEUDOCODE

1. FUNCTION is_palindrome(ogba)

2. reversedogba equal reverse of ogba

3. IF ogba equal reversedogba 

4. THEN

5. RETURN True

6. ELSE
7. RETURN False
8. END IF
8. END FUNCTION



## ispalindrome Solution

```python

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


# Test
print(is_palindrome("racecar"))   # True
print(is_palindrome("hello"))     # False
print(is_palindrome("A man a plan a canal Panama"))  # False
```

# STEP 2: USE AI TO LEARN
# Now that your function works, use AI to go deeper:

"I wrote a palindrome function. Here's my code:

```python 

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

```

# What's the time complexity?
    The time complexity is O(n)   # (n) is the length of the string

# What edge cases might I miss?
    The edge cases i might missed include:
    -Empty String
    -Case sensitivity 
    -Space
    -Punctuation

# Are there better approaches?
  - There might be better approaches depending on what you are optimizing your code for. Example;

# Normalized Phrase Palindrome approach

import re

def is_palindrome(s):
    s = s.lower()
    s = re.sub(r'[^a-z0-9]', '', s)
    return s == s[::-1]



# Step 3 - REFLECTION
# WHAT I LEARNT BEFORE USING AI
- I learn to solve the challenge on my own
- I write the logic of the code in plain English before implementing with python
- I discovered the code was error prone after implementation
- I research to learn how to correct syntax errors

# How is your understanding different now?
# MY UNDERSTANDING NOW:
- I can now resolve syntax issues
- I know how functions works now
- i can now write a better pseudocode
- i now understand indentation and case issues in python code

# Could you now write similar functions (e.g., reverse a string) without help?
As a beginner, i cannot accurately write similar function (reverse a string) without a help but i can confidently write the basic for implementation.
