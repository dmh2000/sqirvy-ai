import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
from matplotlib.ticker import MaxNLocator


import numpy as np
import csv
import sys


def create_diff_matrix(fname):

    x = np.array([])
    y = np.array([])
    z = np.array([])

    try:
        with open(fname, "r") as f:
            rdr = csv.reader(open(fname, "r"))
            data = list(rdr)
    except:
        print("Error reading file")
        sys.exit(1)

    # Convert string data to integers/floats
    data = [[int(row[0]), int(row[1]), float(row[2])] for row in data]
    d = len(data)

    # Find the matrix dimensions
    max_row = max(row[0] for row in data)
    max_col = max(row[1] for row in data)

    # Create an empty matrix filled with zeros
    diff_matrix = np.zeros((max_row, max_col))

    # Populate the matrix with values
    for row, col, value in data:
        diff_matrix[row - 1, col - 1] = value

    # invert the matrix to match cosine similarity
    diff_matrix = 1.0 - diff_matrix

    # normalize it to 0..1
    min_val = np.min(diff_matrix)
    max_val = np.max(diff_matrix)
    diff_matrix = (diff_matrix - min_val) / (max_val - min_val)

    # print similarity matrix, mean and standard deviation
    mean = np.mean(diff_matrix)
    std = np.std(diff_matrix)

    return diff_matrix, mean, std


def plot_diff_matrix(ax, bx, diff_matrix):
    # Plot the similarity matrix
    sx = diff_matrix.shape[1]
    sy = diff_matrix.shape[0]

    x = np.linspace(0, 1.0, sx)
    y = np.linspace(0, 1.0, sy)
    X, Y = np.meshgrid(x, y)
    Z = diff_matrix

    # 2D Scatter
    ax.set_xticks(np.arange(0, 1.0, 0.2))
    ax.set_yticks(np.arange(0, 1.0, 0.2))
    ax.set_xlabel("Diffs of files compared to others")
    ax.set_ylabel("Distribution ")
    ax.set_title("Diff Similarity Matrix")
    ax.legend(
        ("Higher Value = More Similar", "Terms are normalized 0..1"),
        loc="upper center",
        shadow=True,
    )

    scatter = ax.scatter(X, Z, c=Z, cmap="viridis")
    plt.colorbar(scatter)

    # 1D histogram
    zf = np.ravel(Z)
    # plot histogram
    bx.set_xlabel("Similarity")
    bx.set_ylabel("Frequency")
    bx.hist(zf, bins=sx * 2)


if __name__ == "__main__":
    # get a filename from the command line
    if len(sys.argv) < 2:
        print("Usage: python diffs.py <csv file>")
        sys.exit(1)
    fname = sys.argv[1]

    diff_matrix = create_diff_matrix(fname)
    print(diff_matrix)
    fig = plt.figure()
    ax = fig.add_subplot(111)
    bx = fig.add_subplot(211)
    plot_diff_matrix(ax, bx, diff_matrix)
    plt.show()
