import task1

def compute_omega_and_phi_estimates(SNR_dB: float, FFT_size: float):
    rng = np.random.default_rng()
    sigma = get_sigma(get_SNR(SNR_dB))
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
