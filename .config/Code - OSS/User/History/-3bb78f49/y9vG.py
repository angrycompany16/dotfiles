import matplotlib.pyplot as plt
import numpy as np
import argparse, sys

LAMBDA = -2
DT = 0.0001
X_0 = 1
T_START = 0
T_END = 2
NUM_STEPS = int((T_END - T_START) / DT)

def true_solution(t):
    return X_0 * np.exp(LAMBDA * t)

def f(x, t):
    return LAMBDA * x

def RK1(x_dot: callable):
    x_values = np.zeros(NUM_STEPS)
    x_values[0] = X_0
    for i in range(NUM_STEPS - 1):
        x_values[i + 1] = x_values[i] + DT * x_dot(x_values[i], i * DT)

    return x_values

def RK2(x_dot: callable):
    x_values = np.zeros(NUM_STEPS)
    x_values[0] = X_0
    for i in range(NUM_STEPS - 1):
        K_1 = x_dot(x_values[i], i * DT)
        K_2 = x_dot(x_values[i] + 0.5 * DT * K_1, (i + 0.5) * DT)
        x_values[i + 1] = x_values[i] + DT * K_2
    
    return x_values

def RK4(x_dot: callable):
    x_values = np.zeros(NUM_STEPS)
    x_values[0] = X_0
    for i in range(NUM_STEPS - 1):
        K_1 = x_dot(x_values[i], i * DT)
        K_2 = x_dot(x_values[i] + 0.5 * DT * K_1, (i + 0.5) * DT)
        K_3 = x_dot(x_values[i] + 0.5 * DT * K_2, (i + 0.5) * DT)
        K_4 = x_dot(x_values[i] + DT * K_3, (i + 1) * DT)

        x_values[i + 1] = x_values[i] + DT * (1/6 * K_1 + 1/3 * K_2 + 1/3 * K_3 + 1/6 * K_4)
    
    return x_values

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--method")
    args = parser.parse_args()

    x_values = np.zeros(NUM_STEPS)
    t_values = np.linspace(T_START, T_END, NUM_STEPS)
    if args.method == "RK1":
        x_values = RK1(f)
    elif args.method == "RK2":
        x_values = RK2(f)
    elif args.method == "RK4":
        x_values = RK4(f)

    fig, ax = plt.subplots()
    ax.plot(t_values, x_values, label='numerical solution')
    ax.plot(t_values, true_solution(t_values), label='true solution')
    ax.legend()
    plt.show()

if __name__ == "__main__":
    main()