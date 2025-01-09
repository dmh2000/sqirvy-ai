# !/bin/sh
# this script will find all the files in the current directory and its subdirectories
# that match one of the following extensions
# this one looks for code files
extensions=(
    ".go"			#:Go
    ".py"			#:Python
    ".js"			#:JavaScript
    ".java"			#:Java
    ".cpp"			#:C++
    ",cxx"			#:C++
    ".cc"			#:C++
    ".c"			#:C
    ".cs"			#:C			#
    ".php"			#:PHP
    ".rb"			#:Ruby
    ".swift"		#:Swift
    ".ts"			#:TypeScript
    ".html"			#:HTML
    ".htm"			#:HTML
    ".css"			#:CSS
    ".sql"			#:SQL
    ".pl"			#:Perl
    ".sh"			#:Shellscript
    ".r"			#:R
    ".m"			#:MATLABorObjective-C
    ".scala"		#:Scala
    ".kt"			#:Kotlin
    ".rs"			#:Rust
    ".hs"			#:Haskell
    ".lua"			#:Lua
    ".vb"			#:VisualBasic
    ".f"			#:Fortran
    ".f90"			#:Fortran
    ".jl"           #:Julia
)
current_dir=$(pwd)
find . -type f | while read -r file; do
    # Extract the file extension
    extension="${file##*.}"
    
    # Check if the extension is in our list
    for ext in "${extensions[@]}"; do
        if [[ ".$extension" == "$ext" ]]; then
            echo " $file" | sed "s|^$current_dir/||"
            break
        fi
    done
done