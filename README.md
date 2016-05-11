# fastpoisson
Faster Poisson for Go / Golang

Generate random Poisson-distributed numbers

Basic Knuth Algorithm from
https://en.wikipedia.org/wiki/Poisson_distribution#Generating_Poisson-distributed_random_variables
with cached L. For lambda < 1 this was a 4x speed up over the fastest implementation for Go we found.

If you need further speed for large lambdas the table based algorithm from
  https://www.jstatsoft.org/article/view/v011b03/v11b03.pdf
should beat almost anything. There's a matlab implementation here: 
http://www.mathworks.com/matlabcentral/fileexchange/7309-randraw

Send me a pull request if you implement the table implementation.
Also at that point some benchmarking code would make a lot of sense.