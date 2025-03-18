contest="xueersi"
cd ./$contest/

name="I"
# name="BE"
cd ./$name/

file="./${name}.cpp"
g++ -g $file -o $name

# printf "enter the index of test case:"
# read i
# in="${i}.in"
# out="${name}.out"
# touch $out

# ./$name < $in

./$name
