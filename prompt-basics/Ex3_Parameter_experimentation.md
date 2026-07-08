
# Exercise 3: Parameter Experimentation
# Objective: Understand how parameters affect output diversity.
Parameters like temperature and top-p control how random or predictable the model’s word choices are, so increasing them makes outputs more varied and creative while lowering them makes responses more consistent and focused.

# Use the same prompt with temperature = 0.2 and then with temperature = 0.9.
- The argument demonstrates that keeping the same prompt but changing the temperature shows how lower values produce precise, consistent responses, while higher values generate more creative and varied outputs.

# Repeat with different top-p values (e.g., 0.5 vs 1)
Changing the same prompt with different top-p values shows that lower top-p (0.5) limits word choice for focused, predictable responses, while higher top-p (1) allows more varied and expressive language with greater creativity but slightly less consistency.

# Record how the style, randomness, and focus of responses change.
Using the same prompt shows that lower temperature (e.g., 0.2) produces short, precise, and predictable answers, while higher temperature (e.g., 0.9) creates longer, more creative, and varied responses; similarly, a lower top-p value (0.5) restricts the model to the most likely words and keeps replies focused and repetitive, whereas a higher top-p value (1.0) allows a wider choice of words, resulting in richer language, more detail, and greater variation but slightly less consistency.