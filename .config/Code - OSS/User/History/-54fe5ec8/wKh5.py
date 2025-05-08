import numpy as np
import cmath
import os
import time
import matplotlib.pyplot as plt
import scipy.optimize as opt
import math

T = 1e-6
F_s = 1e6
omega_0 = 2 * np.pi * 1e5
phi = np.pi / 8
A = 1
N = 513
n_0 = -256
iterations = 30

P = N * (N - 1) / 2
Q = N * (N - 1) * (2 * N - 1) / 2

def get_sigma_sqr(SNR: float) -> tuple[float, float]:
    return (np.pow(A, 2) / (2 * SNR), A / np.sqrt((2 * SNR)))

def omega_CRLB(sigma_sqr: float) -> float:
    return (12  * sigma_sqr) / (np.pow(A, 2) * np.pow(T, 2) * N * ((np.pow(N, 2) - 1)))

def phi_CRLB(sigma_sqr: float) -> float:
    return (12 * sigma_sqr) * (np.pow(n_0, 2) * N + 2 * n_0 * P + Q) / (np.pow(A, 2) * np.pow(N, 2) * ((np.pow(N, 2) - 1)))

def F(omega: float, x: np.typing.ArrayLike) -> complex:
    s = 0
    for n in range(N):
        s += x[n] * np.exp(-1j * omega * n * T)
    return s / N

def get_phi_est(omega_est: float, x: np.typing.ArrayLike) -> float:
    return cmath.phase(np.exp(-1j * omega_est * n_0 * T) * F(omega_est, x))

def get_SNR(snr_dB: float) -> float:
    return np.pow(10, snr_dB / 10)

SNRs = [get_SNR(SNR_dB) for SNR_dB in range (-10, 61, 10)]
FFT_sizes = [np.pow(2, k) for k in range(10, 21, 2)]

# TODO: Compute all n samples at the same time
def sample(sigma: float, rng: np.random.Generator) -> np.typing.ArrayLike:
    w_r = rng.normal(loc=0, scale=sigma, size=N)
    w_i = rng.normal(loc=0, scale=sigma, size=N)
    x = np.zeros(N, dtype=complex)
    for n in range(n_0, n_0 + N):
        x[n - n_0] = A * np.exp(1j * (omega_0 * n * T + phi)) + w_r[n - n_0] + 1j * w_i[n - n_0]
    return x

def solve_expensive(SNR_dB: float):
    seed = time.time()
    rng = np.random.default_rng(seed=int(seed))
    
    SNR = get_SNR(SNR_dB)
    sigma_sqr, sigma = get_sigma_sqr(SNR)

    # Variance
    omegafig, omegaax = plt.subplots()
    phifig, phiax = plt.subplots()

    omega_variances = np.zeros(len(FFT_sizes))
    phi_variances = np.zeros(len(FFT_sizes))

    omega_CRLB_arr = np.ones(len(FFT_sizes)) * omega_CRLB(sigma_sqr)
    phi_CRLB_arr = np.ones(len(FFT_sizes)) * phi_CRLB(sigma_sqr)

    # Error
    omegaerrfig, omegaerrax = plt.subplots()
    phierrfig, phierrax = plt.subplots()

    omega_means = np.zeros(len(FFT_sizes))
    phi_means = np.zeros(len(FFT_sizes))

    omega_true_arr = np.ones(len(FFT_sizes)) * omega_0
    phi_true_arr = np.ones(len(FFT_sizes)) * phi

    for (i, FFT_size) in enumerate(FFT_sizes):
        omega_estimates = np.zeros(iterations)
        phi_estimates = np.zeros(iterations)

        for n in range(iterations):
            x_samples = sample(sigma, rng)
            
            # print(np.abs(x_samples))
            # plt.plot(x_values, np.real(x_samples))
            # plt.plot(x_values, np.imag(x_samples))
            FFT = np.fft.fft(x_samples, n=FFT_size)

            # omega_values = np.linspace(0, FFT_size, FFT_size)
            # plt.plot(omega_values, np.abs(FFT))

            m_max = np.argmax(np.abs(FFT))

            omega_est = 2 * np.pi * m_max / (FFT_size * T)
            phi_est = get_phi_est(omega_est, x_samples)
            
            omega_estimates[n] = omega_est
            phi_estimates[n] = phi_est
        
        omega_variances[i] = np.var(omega_estimates)
        phi_variances[i] = np.var(phi_estimates)

        omega_means[i] = np.mean(omega_estimates)
        phi_means[i] = np.mean(phi_estimates)

    # Variance
    omegaax.plot(np.linspace(0, len(FFT_sizes) - 1, len(FFT_sizes)), omega_variances, label=["Omega variance"])
    omegaax.plot(np.linspace(0, len(FFT_sizes) - 1, len(FFT_sizes)), omega_CRLB_arr, label=["CRLB"])
    omegaax.legend()

    phiax.plot(np.linspace(0, len(FFT_sizes) - 1, len(FFT_sizes)), phi_variances, label=["Phi variance"])
    phiax.plot(np.linspace(0, len(FFT_sizes) - 1, len(FFT_sizes)), phi_CRLB_arr, label=["CRLB"])    
    phiax.legend()

    # Error
    omegaerrax.plot(np.linspace(0, len(FFT_sizes) - 1, len(FFT_sizes)), omega_means, label=["Omega mean"])
    omegaerrax.plot(np.linspace(0, len(FFT_sizes) - 1, len(FFT_sizes)), omega_true_arr, label=["omega true value"])
    omegaerrax.legend()

    phierrax.plot(np.linspace(0, len(FFT_sizes) - 1, len(FFT_sizes)), phi_means, label=["Phi mean"])
    phierrax.plot(np.linspace(0, len(FFT_sizes) - 1, len(FFT_sizes)), phi_true_arr, label=["phi true value"])
    phierrax.legend()

def solve_cheap(SNR_dB: float, FFT_size: float):
    seed = time.time()
    rng = np.random.default_rng(seed=int(seed))
    
    SNR = get_SNR(SNR_dB)
    sigma_sqr, sigma = get_sigma_sqr(SNR)

    omega_rough_estimates = np.zeros(iterations)
    omega_refined_estimates = np.zeros(iterations)
    # phi_estimates = np.zeros(iterations)

    for i in range(iterations):
        x_samples = sample(sigma, rng)
            
        # print(np.abs(x_samples))
        # plt.plot(x_values, np.real(x_samples))
        # plt.plot(x_values, np.imag(x_samples))
        FFT = np.fft.fft(x_samples, n=FFT_size)

        # omega_values = np.linspace(0, FFT_size, FFT_size)
        # plt.plot(omega_values, np.abs(FFT))

        m_max = np.argmax(np.abs(FFT))

        omega_est_rough = 2 * np.pi * m_max / (FFT_size * T)
        omega_rough_estimates[i] = omega_est_rough

        def optimizer(x, samples) -> float:
            return -np.abs(F(x, samples))

        omega_refined_estimates[i] = opt.minimize(optimizer, omega_est_rough, x_samples, method='Nelder-Mead').x


    # phi_est = get_phi_est(omega_est, x_samples)
    
    # omega_estimates[n] = omega_est_refined
    # phi_estimates[n] = phi_est
    
    # omega_est_average = np.mean()
    print("Variance rough")
    print(np.var(omega_rough_estimates))
    print("Mean rough")
    print(np.mean(omega_rough_estimates))

    print("Variance refined")
    print(np.var(omega_refined_estimates))
    print("Mean refined")
    print(np.mean(omega_refined_estimates))
    # print("phi_variance")

def main():
    # solve_expensive(-10)
    solve_cheap(20, FFT_sizes[0])
    plt.show()

if __name__ == "__main__":
    main()