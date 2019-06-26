import pandas as pd
import numpy as np
import sys
import matplotlib.pyplot as plt
from pandas.plotting import register_matplotlib_converters

def main():
    args = sys.argv
    df = pd.read_csv(args[1], encoding="UTF-8", delimiter=',' ,header=None)
    register_matplotlib_converters()
    df[0] = pd.to_datetime(df[0])

    plt.figure(figsize = (120,7), dpi = 100)
    plt.plot(df[0],df[1])
    plt.xlabel("Timestamp")
    plt.ylabel("Executing Point")

    filename = args[1].split(".")[0] + ".png"

    plt.savefig(filename)
if __name__ == "__main__":
    main()