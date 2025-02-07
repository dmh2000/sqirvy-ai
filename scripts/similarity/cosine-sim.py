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
    with open(filename, 'r') as file:
        return file.read()
    

# read multiple documents
def read_documents(filenames):
    documents = []
    for filename in filenames:
        documents.append(read_document(filename))
    return documents


# get filenames from the command line
if len(sys.argv) != 2:
    print("Usage: python cosine-sim.py <directory>")
    sys.exit(1)

# list filenames from directory './code'
try:
    directory = sys.argv[1]
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
print("Cosine Similarity Matrix:")
print(similarity_matrix)
print("Mean similarity   : ", mean)
print("Standard deviation: ", std)

# Plot the similarity matrix
sx = similarity_matrix.shape[1]
sy = similarity_matrix.shape[0]

x = np.linspace(0, 1.1, sx)
y = np.linspace(0, 1.1, sy)
X,Y = np.meshgrid(x,y)
Z = similarity_matrix

fig = plt.figure(figsize=(10, 6))

# 3D Scatter
# ax = fig.add_subplot(gs[0,0], projection='3d')
# catter = ax.scatter(X, Y, Z, c=Z, cmap='viridis')
# ax.set_xticks(np.arange(0,1.2, 0.2))
# ax.set_yticks(np.arange(0,1.2, 0.2))
# ax.set_xlabel('code files 1..n')
# ax.set_ylabel('code files 1..n')
# ax.set_zlabel('Similarity (higher is more similar)')

# 2D Scatter
ax = fig.add_subplot(111)
ax.set_xticks(np.arange(0,1.2, 0.2))
ax.set_yticks(np.arange(0,1.2, 0.2))
ax.set_xlabel('code files 1..n scaled to 0..1')
ax.set_ylabel('Similarity (higher is more similar)')
scatter = ax.scatter(X,Z,c=Z, cmap='viridis')

plt.show()
a = 1
