import task1
import numpy as np

def compute_omega_and_phi_estimates(SNR_dB: float, FFT_size: float):
    rng = np.random.default_rng()
    sigma = task1.get_sigma(task1.get_SNR(SNR_dB))
    omega_estimates = np.zeros(task1.iterations)
    phi_estimates = np.zeros(task1.iterations)

    for n in range(task1.iterations) :
        x_samples = task1.sample(sigma, rng)
        FFT = np.fft.fft(x_samples, FFT_size)
        m_max = np.argmax(np.abs(FFT))

        omega_est_rough = 2 * np.pi * m_max / (FFT_size * task1.T)
        # omega_estimates_rough[n] = omega_est_rough
        
        def optimizer(x, samples) -> float:
            return -np.abs(F(x, samples))
        
        phi_est = task1.get_phi_est(omega_est_rough, x_samples)

        phi_estimates[n] = phi_est



    return omega_estimates_rough, phi_estimates
