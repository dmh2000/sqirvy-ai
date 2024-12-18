from setuptools import setup
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
    packages=[
        "src/common",
    ],
    entry_points={
        "console_scripts": [
            "sqirvy-doc=sqirvy_doc:main",
            "sqirvy-review=sqirvy_review:main",
        ]
    },
)
