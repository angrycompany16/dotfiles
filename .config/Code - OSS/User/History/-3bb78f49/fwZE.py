import matplotlib.pyplot as plt
import numpy as np
import argparse, sys

LAMBDA = -2
DT = 0.4
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
    pass

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
        RK4(f)

    fig, ax = plt.subplots()
    ax.plot(x_values, t_values, label='numerical solution')
    ax.plot(true_solution(t_values), t_values, label='true solution')
    ax.legend()
    plt.show()

if __name__ == "__main__":
    main()