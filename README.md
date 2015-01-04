# jingu

jingu is brainf\*ck compiler made with golang.

# Getting Started

Install jingu:

~~~sh
$ go get github.com/zakuro9715/jingu
~~~

And run:
~~~sh
$ jingu
>>> +++++++++[>++++++++>+++++++++++>+++++<<<-]>.>++.+++++++..+++.>-.------------.<++++++++.--------.+++.------.--------.>+.>++++++++++.
Hello world!
~~~


# Feature

Source code from file:

~~~sh
$ jingu helloworld.bf
Hello world
~~~

and interactive mode with -i option:

~~~sh
$ jing -i
>>> +++++++++[>++++++++>+++++++++++>+++++<<<-]>.>++.+++++++..+++.>-.------------.<++++++++.--------.+++.------.--------.>+.>++++++++++.
Hello world
~~~

Compile:

~~~sh
$ jingu -c helloworld.bf
$ ./a.out
Hello world
~~~

You can compile with interactive mode:

~~~sh
$ jingu -ci
>>> +++++++++[>++++++++>+++++++++++>+++++<<<-]>.>++.+++++++..+++.>-.------------.<++++++++.--------.+++.------.--------.>+.>++++++++++.
amaterasu
>>> +++++++++[>++++++++>+++++++++++>+++++<<<-]>.>++.+++++++..+++.>-.------------.<++++++++.--------.+++.------.--------.>+.>++++++++++.
susanoo
>>> +++++++++[>++++++++>+++++++++++>+++++<<<-]>.>++.+++++++..+++.>-.------------.<++++++++.--------.+++.------.--------.>+.>++++++++++.
tukuyomi
>>> Ctrl-d
Bye
$ ./amaterasu.out
Hello world
$ ./susanoo.out
Hello world
$ ./tukuyomi.out
Hello world
~~~

# Why Jingu

I went to Jingu(Grand Shrine in Ise) in the day on that this project began.

Jingu worship two goddesses, Amaterasu-Omikami and Toyouke-Omikami.
Amaterasu-Omikami is goddess of sun and progenitor of Tenno(Emperor of Japan) and all Japanese.
Toyouke-Omikami is goddess of agriculture and industry.

I trust that I made this compiler with power of two goddesses 

so, I name jingu this project in gratitude.


# Do you want to learn about Jingu(or Amaterasu-Omikami, Toyouke-Omikami)

I show you websites about Jingu, Amaterasu-omikami, Toyouke-omikami, or Other Japanese Sinto

* [Isejingu](http://www.isejingu.or.jp/english/) - Official website of Jingu
* [Ise Grand Shrine - Wikipedia, the free encyclopedia](http://en.wikipedia.org/wiki/Ise_Grand_Shrine)
