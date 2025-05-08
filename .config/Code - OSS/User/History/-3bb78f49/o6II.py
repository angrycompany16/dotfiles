import matplotlib.pyplot as plt
import numpy as np
import argparse, sys

LAMBDA = -2
DT = 0.4
X_0 = 1
T_START = 0
T_END = 2
NUM_STEPS = (T_END - T_START) / DT

def f(t, x):
    return LAMBDA * x

def RK1(x_dot: callable):
    for i in range(NUM_STEPS):
        pass

def RK2(x_dot: callable):
    pass

def RK4(x_dot: callable):
    pass

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--method")
    args = parser.parse_args()

    if args.method == "RK1":
        RK1(f)
    elif args.method == "RK2":
        RK1(f)
    elif args.method == "RK4":
        RK4(f)

if __name__ == "__main__":
    main()