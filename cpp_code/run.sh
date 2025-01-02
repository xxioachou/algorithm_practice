# contest="2024ECFinal"
# cd ./$contest/
name="CF2044H"
cd ./$name/

file="./${name}.cpp"
g++ -g $file -o $name

out="${name}.out"
touch $out

printf "enter the number of test cases: "
read cnt
i=1
while (( $i <= cnt))
do
    in="${i}.in"
    ans="${i}.out"
    ./$name < $in > $out
    if diff -Z $out $ans;then
        printf "AC on test %d\n" $i
    else
        printf "WA on test %d\n" $i
    fi
    i=`expr $i + 1`
done