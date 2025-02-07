import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
from matplotlib.ticker import MaxNLocator


import numpy as np
import csv
import sys

# get a filename from the command line
if len(sys.argv) < 2:
    print("Usage: python diffs.py <filename>")
    sys.exit(1)

filename = sys.argv[1]


fig = plt.figure()
ax = fig.add_subplot(111, projection='3d')

x = np.array([])
y = np.array([])
z = np.array([])

try:
   with open(filename, 'r') as f:
        data = csv.reader(open(filename, 'r'))
        for row in data:
            if int(row[0]) != int(row[1]):
                x = np.append(x,float(row[0]))
                y = np.append(y,float(row[1]))
                z = np.append(z,float(row[2]))
except:
    print("Error reading file")
    sys.exit(1)

zmin = np.min(z)
zmax = np.max(z)

ax.set_xticks(np.arange(np.min(x), np.max(x), 1))
ax.set_yticks(np.arange(np.min(y), np.max(y), 1))
#ax.set_ylim(np.min(y), np.max(y))
#ax.set_zlim(zmin, zmax)
ax.set_xlabel('X axis')
ax.set_ylabel('Y axis')
ax.set_zlabel('Z axis')

ax.xaxis.set_major_locator(MaxNLocator(integer=True))

# normalize z
z = (z - np.min(z)) / (np.max(z) - np.min(z))

ax.scatter(x, y, z, color='r')  
plt.show()