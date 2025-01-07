# contest="ABC383"
# cd ./$contest/
name="CF1789C"
cd ./$name/

file="${name}.go"
go build $file

printf "enter the index of test case:"
read i
in="${i}.in"

./$name < $in
