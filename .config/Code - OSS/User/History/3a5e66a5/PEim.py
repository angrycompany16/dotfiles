import numpy as np
import cmath
import os
import time
import matplotlib.pyplot as plt
import scipy.optimize as opt
import math

T = 1e-6
__F_s = 1e6
omega_0 = 2 * np.pi * 1e5
phi = np.pi / 8
A = 1
N = 513
__n_0 = -256
iterations = 30


P = N * (N - 1) / 2
Q = N * (N - 1) * (2 * N - 1) / 6

def __get_sigma_sqr(SNR: float) -> float :
    return np.pow(A, 2) / (2 * SNR)

def get_sigma(SNR: float) -> float :
    return A / np.sqrt((2 * SNR))

def __omega_CRLB(sigma_sqr: float) -> float:
    return (12  * sigma_sqr) / (np.pow(A, 2) * np.pow(T, 2) * N * ((np.pow(N, 2) - 1)))

def __phi_CRLB(sigma_sqr: float) -> float:
    return (12 * sigma_sqr) * (np.pow(__n_0, 2) * N + 2 * __n_0 * P + Q) / (np.pow(A, 2) * np.pow(N, 2) * ((np.pow(N, 2) - 1)))

def F(omega: float, x: np.typing.ArrayLike) -> complex:
    s = 0
    for n in range(N):
        s += x[n] * np.exp(-1j * omega * n * T)
    return s / N

def get_phi_est(omega_est: float, x: np.typing.ArrayLike) -> float:
    return cmath.phase(np.exp(-1j * omega_est * __n_0 * T) * F(omega_est, x))

def get_snr(snr_dB: float) -> float:
    return np.pow(10, snr_dB / 10)

SNRs = [get_snr(SNR_dB) for SNR_dB in range (-10, 61, 10)]
SNRs_in_dB = [SNR_dB for SNR_dB in range (-10, 61, 10)]
FFT_sizes = [np.pow(2, k) for k in range(10, 21, 2)]
sigma_sqrs = [__get_sigma_sqr(SNR) for SNR in SNRs]
omega_CRLBs = [__omega_CRLB(sigma_sqr) for sigma_sqr in sigma_sqrs]
phi_CRLBs = [__phi_CRLB(sigma_sqr) for sigma_sqr in sigma_sqrs]

def sample(sigma: float, rng: np.random.Generator) -> np.typing.ArrayLike:
    w_r = rng.normal(loc=0, scale=sigma, size=N)
    w_i = rng.normal(loc=0, scale=sigma, size=N)
    x = np.zeros(N, dtype=complex)
    for n in range(__n_0, __n_0 + N):
        x[n - __n_0] = A * np.exp(1j * (omega_0 * n * T + phi)) + w_r[n - __n_0] + 1j * w_i[n - __n_0]
    return x

def estimate_omega_phi(SNR_dB: float, FFT_size: float):
    rng = np.random.default_rng()
    sigma = get_sigma(get_snr(SNR_dB))
    omega_estimates = np.zeros(iterations)
    phi_estimates = np.zeros(iterations)

    for n in range(iterations) :
        x_samples = sample(sigma, rng)
        FFT = np.fft.fft(x_samples, FFT_size)
        m_max = np.argmax(np.abs(FFT))

        omega_est = 2 * np.pi * m_max / (FFT_size * T)
        phi_est = get_phi_est(omega_est, x_samples)

        omega_estimates[n] = omega_est
        phi_estimates[n] = phi_est

    return omega_estimates, phi_estimates

def get_variance(estimates: np.typing.ArrayLike):
    return np.var(estimates)

def plot_variances_omega_against_snr(omega_variances: np.typing.ArrayLike):
    epsilon = 1e-2
    clipped_omega_variances = np.clip(omega_variances, epsilon, None)
    
    plt.figure(figsize=(10, 6))
    plt.plot(SNRs_in_dB, clipped_omega_variances, marker='o', linestyle='-', label='Frequency Estimator Variance')
    plt.plot(SNRs_in_dB, omega_CRLBs, marker='s', linestyle='--', label='Frequency CRLB')

    plt.yscale('log')
    plt.xlabel('SNR (dB)')
    plt.ylabel('Frequency Estimator Variance')
    plt.title('Frequency Estimator Variance and CRLB vs SNR')
    plt.grid(True, linestyle='--', alpha=0.7)
    plt.legend()
    plt.tight_layout()
    plt.show()

def plot_average_omega_against_snr(omega_averages_array: np.typing.ArrayLike):
    true_value = np.full(len(SNRs_in_dB), omega_0)

    plt.figure(figsize=(10, 6))
    for i, omega_averages in enumerate(omega_averages_array):
        plt.plot(SNRs_in_dB, omega_averages, marker='o', linestyle='-', label=f'Average Frequency Estimate  w/ fft size {math.log2(FFT_sizes[i])}')

    plt.plot(SNRs_in_dB, true_value, label=f'True frequency')

    plt.yscale('log')
    plt.xlabel('SNR (dB)')
    plt.ylabel('Frequency')
    plt.title('Average Frequency Estimate and True Value vs SNR')
    plt.grid(True, linestyle='--', alpha=0.7)
    plt.legend()
    plt.tight_layout()
    plt.show()

def plot_variances_phi_against_snr(phi_variances: np.typing.ArrayLike):
    plt.figure(figsize=(10, 6))
    plt.plot(SNRs_in_dB, phi_variances, marker='o', linestyle='-', label='Phi Variance')
    plt.plot(SNRs_in_dB, phi_CRLBs, marker='s', linestyle='--', label='Phi CRLB')

    plt.yscale('log')
    plt.xlabel('SNR (dB)')
    plt.ylabel('Phase Estimator Variance')
    plt.title('Phase Estimator Variance and CRLB vs SNR')
    plt.grid(True, linestyle='--', alpha=0.7)
    plt.legend()
    plt.tight_layout()
    plt.show()

def plot_average_phi_against_snr(phi_averages_array: np.typing.ArrayLike):
    true_value = np.full(len(SNRs_in_dB), phi)
    plt.figure(figsize=(10, 6))
    for i, phi_averages in enumerate(phi_averages_array):
        plt.plot(SNRs_in_dB, phi_averages, marker='o', linestyle='-', label=f'Average Phase Estimate w/ fft_size {math.log2(FFT_sizes[i])}')

    plt.plot(SNRs_in_dB, true_value, label=f'True phase')

    plt.yscale('log')
    plt.xlabel('SNR (dB)')
    plt.ylabel('Phase')
    plt.title('Average Phase Estimate and True Value vs SNR')
    plt.grid(True, linestyle='--', alpha=0.7)
    plt.legend()
    plt.tight_layout()
    plt.show()