name="LC855"
cd ./$name/

file="${name}.cpp"
g++ -g $file -o $name

./$name

# printf "enter the index of test case: "
# read nth
# in="${nth}.in"
# out="${name}.out"
# touch $out
# ./$name < $in