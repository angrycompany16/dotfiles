import task1
import task2
import numpy as np
import argparse, sys

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--fft_size")
    parser.add_argument("--task")
    parser.add_argument("--mean")
    args=parser.parse_args()

    if int(args.task) == 1:
        if args.mean == 'True':
            solve_mean_task1()
        else:
            solve_task1(args.fft_size)
    elif int(args.task) == 2:
        if args.mean == 'True':
            solve_mean_task2()                
        else:
            solve_task2()

def solve_mean_task1():
    omega_average_estimates = [np.zeros(len(task1.SNRs_in_dB)) for i in task1.FFT_sizes]
    phi_average_estimates = [np.zeros(len(task1.SNRs_in_dB)) for i in task1.FFT_sizes]
    for i, fft_size in enumerate(task1.FFT_sizes):
        for j, SNR_dB in enumerate(task1.SNRs_in_dB):
            omega_estimates, phi_estimates = task1.estimate_omega_phi(SNR_dB, fft_size)

            omega_average_estimates[i][j] = np.mean(omega_estimates)
            phi_average_estimates[i][j] = np.mean(phi_estimates)
        
    task1.plot_average_omega_against_snr(omega_average_estimates)
    task1.plot_average_phi_against_snr(phi_average_estimates)

def solve_task1(fft_size: int):
    fft_size = task1.FFT_sizes[int(fft_size)]

    omega_variance_array = np.zeros(len(task1.SNRs_in_dB))
    phi_variance_array = np.zeros(len(task1.SNRs_in_dB))
    for i, SNR_dB in enumerate(task1.SNRs_in_dB):
        omega_estimates, phi_estimates = task1.estimate_omega_phi(SNR_dB, fft_size)

        omega_variance, phi_variance = task1.get_variance(omega_estimates), task1.get_variance(phi_estimates)

        omega_variance_array[i] = omega_variance
        phi_variance_array[i] = phi_variance

    task1.plot_variances_omega_against_snr(omega_variance_array)
    task1.plot_variances_phi_against_snr(phi_variance_array)

def solve_mean_task2():
    omega_average_estimates = np.zeros(len(task1.SNRs_in_dB))
    phi_average_estimates = np.zeros(len(task1.SNRs_in_dB))
    for i, SNR_dB in enumerate(task1.SNRs_in_dB):
        omega_estimates, phi_estimates = task2.estimate_omega_phi(SNR_dB, task1.FFT_sizes[0])

        omega_average_estimates[i] = np.mean(omega_estimates)
        phi_average_estimates[i] = np.mean(phi_estimates)
        
    task2.plot_average_omega_against_snr(omega_average_estimates)
    task2.plot_average_phi_against_snr(phi_average_estimates)

def solve_task2():
    print("Solving task 2")
    omega_variance_array = np.zeros(len(task1.SNRs_in_dB))
    phi_variance_array = np.zeros(len(task1.SNRs_in_dB))
    for i, SNR_dB in enumerate(task1.SNRs_in_dB):
        omega_estimates, phi_estimates = task2.estimate_omega_phi(SNR_dB, task1.FFT_sizes[2])

        omega_variance, phi_variance = task1.get_variance(omega_estimates), task1.get_variance(phi_estimates)

        omega_variance_array[i] = omega_variance
        phi_variance_array[i] = phi_variance

    task1.plot_variances_omega_against_snr(omega_variance_array)
    task1.plot_variances_phi_against_snr(phi_variance_array)

if __name__ == "__main__":
    main()
