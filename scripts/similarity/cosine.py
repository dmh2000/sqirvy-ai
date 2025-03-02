from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity
from matplotlib.gridspec import GridSpec
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
import numpy as np
import sys
import os


# read a document from a file
def read_document(filename):
    with open(filename, "r") as file:
        return file.read()


# read multiple documents
def read_documents(filenames):
    documents = []
    for filename in filenames:
        documents.append(read_document(filename))
    return documents


def plot_similarity_matrix(ax, bx, similarity_matrix):
    # Plot the similarity matrix
    sx = similarity_matrix.shape[1]
    sy = similarity_matrix.shape[0]

    x = np.linspace(0, 1.0, sx)
    y = np.linspace(0, 1.0, sy)
    X, Y = np.meshgrid(x, y)
    Z = similarity_matrix

    # 2D Scatter
    ax.set_xticks(np.arange(0, 1.0, 0.2))
    ax.set_yticks(np.arange(0, 1.0, 0.2))
    ax.set_xlabel("File compared to the others (normalized)")
    ax.set_ylabel("Distribution")
    ax.set_title("Cosine Similarity")
    ax.legend(("Higher Value = More Similar"), loc="upper center", shadow=True)

    scatter = ax.scatter(X, Z, c=Z, cmap="viridis")
    plt.colorbar(scatter)

    # 1D histogram
    bx.set_xlabel("Similarity")
    bx.set_ylabel("Frequency")
    zf = np.ravel(Z)
    # plot histogram
    bx.hist(zf, bins=sx * 2)


def create_similarity_matrix(directory):
    # list filenames from directory './code'
    try:
        files = [f for f in os.listdir(directory)]
        files = [os.path.join(directory, f) for f in files]
    except:
        print("Error reading directory")
        sys.exit(1)

    documents = read_documents(files)

    # Create TF-IDF vectorizer
    vectorizer = TfidfVectorizer()

    # Convert documents to TF-IDF arrays
    tfidf_matrix = vectorizer.fit_transform(documents)

    # Convert to dense NumPy arrays
    doc_arrays = tfidf_matrix.toarray()

    # Compute cosine similarity
    similarity_matrix = cosine_similarity(doc_arrays)

    # print similarity matrix, mean and standard deviation
    mean = np.mean(similarity_matrix)
    std = np.std(similarity_matrix)

    return similarity_matrix, mean, std


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python3 cos-sim.py <directory>")
        sys.exit(1)

    similarity_matrix = create_similarity_matrix(sys.argv[1])

    fig = plt.figure()
    ax = fig.add_subplot(111)
    bx = fig.add_subplot(111)
    plot_similarity_matrix(ax, bx, similarity_matrix)
    plt.show()
