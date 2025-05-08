# Lab work for the subject TTK4275 - Estimation, detection and classification

In this lab we experimented with an FFT approach to making an MLE for a complex exponential variable with white noise.

To run the project, open the `project/` folder, install the requirements using 

```
    pip install -r requirements.txt
```

and run the project with 

```
    python3 main.py --fft_size=X --task=Y --mean=Z
```

The three parameters are:
- fft_size: Applies when computing the variance of the pure FFT estimator. Should be a number 0-5.
- task: 1 or 2. Decides whether to use pure FFT or FFT with numerical optimization.
- mean: should be either `True` or `False`. Decides whether to compute the average or variance of the estimates computed. 