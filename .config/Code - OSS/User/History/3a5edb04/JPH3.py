import task1
import numpy as np
import scipy.optimize as opt

def estimate_omega_phi(SNR_dB: float, FFT_size: float):
    rng = np.random.default_rng()
    sigma = task1.get_sigma(task1.get_snr(SNR_dB))
    omega_estimates = np.zeros(task1.iterations)
    phi_estimates = np.zeros(task1.iterations)

    for n in range(task1.iterations) :
        x_samples = task1.sample(sigma, rng)
        FFT = np.fft.fft(x_samples, FFT_size)
        m_max = np.argmax(np.abs(FFT))

        omega_est_rough = 2 * np.pi * m_max / (FFT_size * task1.T)

        def optimizer(x, samples) -> float:
            return -np.abs(task1.F(x, samples))
        
        omega_est_refined = opt.minimize(optimizer, omega_est_rough, x_samples, method='Powell').x
        phi_est = task1.get_phi_est(omega_est_refined, x_samples)
        
        omega_estimates[n] = omega_est_refined
        phi_estimates[n] = phi_est

    return omega_estimates, phi_estimates

def plot_average_phi_against_snr(phi_averages_array: np.typing.ArrayLike):
    true_value = np.full(len(SNRs_in_dB), phi)
    plt.figure(figsize=(10, 6))
    for i, phi_averages in enumerate(phi_averages_array):
        plt.plot(SNRs_in_dB, phi_averages, marker='o', linestyle='-', label=f'Phi Variance w/ fft_size {i}')

    plt.plot(SNRs_in_dB, true_value, label=f'True value of phi')

    plt.yscale('log')
    plt.xlabel('SNR (dB)')
    plt.ylabel('Variance')
    plt.title('Phi Variance and CRLB vs SNR')
    plt.grid(True, linestyle='--', alpha=0.7)
    plt.legend()
    plt.tight_layout()
    plt.show()