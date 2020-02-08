let unless = macro(cond, cons, alt) {
    quote(if (!(unquote(cond))) {
        unquote(cons);
    } else {
        unquote(alt);
    });
}

let test = fn(arg) {
    puts(arg)
}

let one = 1+1
puts(one)
puts(one+1)
puts("Hello World")
unless(10 > 5, puts("not greater"), puts("greater"));
