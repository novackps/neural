import numpy as np
import matplotlib.pyplot as plt


# Define the modified logistic function with a positive exponent
def modified_logistic(x):
    return 1 / (1 + np.exp(x))


# Generate values for x from -10 to 10
x = np.linspace(-10, 10, 100)
y = modified_logistic(x)

# Plot the modified logistic function
plt.figure(figsize=(8, 5))
plt.plot(x, y, label='Modified Logistic Function', color='green')
plt.title('Modified Logistic Function (Decreasing S-Curve)')
plt.xlabel('x')
plt.ylabel('f(x)')
plt.legend()
plt.grid(True)
plt.show()
