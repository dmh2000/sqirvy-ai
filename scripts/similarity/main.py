import matplotlib.pyplot as plt
import numpy as np
import sys
from cosine import create_similarity_matrix, plot_similarity_matrix
from diffs import create_diff_matrix, plot_diff_matrix

if __name__ == "__main__":
    if len(sys.argv) != 4:
        print(f"Usage: {sys.argv[0],} <model name> <diff csv file>  <code directory>")
        sys.exit(1)

    model = sys.argv[1]
    diff = sys.argv[2]
    code = sys.argv[3]

    diff_matrix, mean, std = create_diff_matrix(diff)
    print("Diff Mean similarity   : ", mean)
    print("Diff Standard deviation: ", std)

    similarity_matrix, mean, std = create_similarity_matrix(code)
    print("Cosine Mean              : ", mean)
    print("Cosine Standard deviation: ", std)

    fig, ax = plt.subplots(4, 1, constrained_layout=True, figsize=(12, 12))
    plot_similarity_matrix(ax[0], ax[1], similarity_matrix)
    plot_diff_matrix(ax[2], ax[3], diff_matrix)
    fig.suptitle(f"Model: {model}")

    plt.savefig(f"{model}.png")
    plt.show()
    a = 1
