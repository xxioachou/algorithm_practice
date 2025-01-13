# contest="ABC383"
# cd ./$contest/
name="lanqiao3136"
cd ./$name/

out="${name}.out"
touch $out
file="${name}.go"
go build $file

printf "enter the number of test cases: "
read cnt
i=1
while (( $i <= cnt))
do
    in="${i}.in"
    ans="${i}.out"
    ./$name < $in > $out
    if diff -Z $out $ans; then
        printf "AC on test ${i}\n" 
    else 
        printf "WA on test ${i}\n"
    fi
    i=`expr $i + 1`
done