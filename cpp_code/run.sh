name="CF2042D"
cd ./$name/

out="${name}.out"
touch $out

file="${name}.cpp"
g++ -g $file -o $name

printf "enter the number of test cases: "
read cnt
i=1
while (( $i <= cnt ))
do
    in="${i}.in"
    ans="${i}.out"
    ./$name < $in > $out
    if diff -Z $out $ans;then
        printf "AC on test %d\n" $i
    else
        printf "WA on test %d\n" $i
        break
    fi
    i=`expr $i + 1`
done