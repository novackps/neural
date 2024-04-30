import numpy as np
import matplotlib.pyplot as plt


# Define the logistic function
def logistic(x):
    return 1 / (1 + np.exp(-x))


# Generate values for x from -10 to 10
x = np.linspace(-10, 10, 100)
y = logistic(x)

# Plot the logistic function
plt.figure(figsize=(8, 5))
plt.plot(x, y, label='Logistic Function', color='blue')
plt.title('Logistic Function (S-Curve)')
plt.xlabel('x')
plt.ylabel('f(x)')
plt.legend()
plt.grid(True)
plt.show()
