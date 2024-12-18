from setuptools import setup
from setuptools import find_packages
setup(
    name="sqirvy_ai",
    version="0.1",
    description="An AI tool for reviewing and documenting code.",
    author="David Howard",
    author_email="dmh2000@gmail.com",
    url="https://github.com/dmh2000/sqirvy_ai",
    install_requires=[
        "anthropic",
        "pyyaml",
    ],
    packages=find_packages(where="src"),
    entry_points={
        "console_scripts": [
            "sqirvy-doc=sqirvy_doc.main:main",
            "sqirvy-review=sqirvy_review.main:main",
        ]
    },
)
