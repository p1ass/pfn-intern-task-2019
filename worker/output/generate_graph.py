import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from pandas.plotting import register_matplotlib_converters

def main():
    df = pd.read_csv("2-1_executing_point.csv", encoding="UTF-8", delimiter=',' ,header=None)
    register_matplotlib_converters()
    df[0] = pd.to_datetime(df[0])

    plt.figure(figsize = (120,7), dpi = 100)
    plt.plot(df[0],df[1])
    plt.xlabel("Timestamp")
    plt.ylabel("Executing Point")
    plt.savefig("2-1.png")
if __name__ == "__main__":
    main()