name="CF845D"
path="./${name}/"
cd $path

file="${name}.go"
go build $file

printf "enter the index of test cases: "
read t
in="${t}.in"
./$name < $in