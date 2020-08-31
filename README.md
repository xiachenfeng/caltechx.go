我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

caltechx.go
===========

This is a [golang](http://golang.org) version of the Machine learning course from caltech: [Learning from data](https://work.caltech.edu/telecourse.html).

* **week 1:**
    * PLA (Perceptron learning Algorithm)
* **week 2:**
    * Hoeffding Inequality
    * Linear Regression
    * Nonlinear Transformation
* **week 3:**
    * Generalization Error
* **week 4:**
    * VC bound
    * Bias and Variance
* **week 5:**
    * Linear Regression Error
    * Gradient Descent
    * Logistic Regression
* **week 6:**
    * Overfitting and Regularization With Weight Decay
    * Neural Networks
* **week 7:**
    * Validation
    * Estimators
    * Cross Validation
    * PLA vs. SVM
* **week 8:**
    * Support Vector Machines With Soft Margins
    * Polynomial Kernels
    * Cross Validation
    * RBF Kernel

##Build:
There is a specific directory `week<x>` for the homework of each week (1 to 8). To build it run the following command where `x = 1`

    go get ./week1
##Run:
Similarly you can run the work of a specific week as follows:

    week1

##Test:
Tests will be slow as they are running the homeworks whom typically have to run multiple "runs" (1000 or more runs) and compute an average.

    go test ./week1


##Todo:
* refactor
* concurrent runs.
* command line animations. [Pretty command line / console output on Unix in Python and Go Lang](http://www.darkcoding.net/software/pretty-command-line-console-output-on-unix-in-python-and-go-lang/)
* refactor PLA and other functions into separate packages.
* linear regression should have a Xn array and an Zn collection when a transformation takes place
* add transpose function.
* transformation function should accept array with param x0 = 1 to transform
* better and consistent print statements.
* catch all error and have all functions send errors.
* add tests

##Current tree:

    $ tree
    .
    ├── LICENSE
    ├── README.md
    ├── biasAndVariance
    │   └── biasAndVariance.go
    ├── data
    │   ├── in.dta
    │   └── out.dta
    ├── generalizationError
    │   └── generalizationerror.go
    ├── gradientDescent
    │   └── gradientDescent.go
    ├── hoeffding
    │   └── hoeffding.go
    ├── linear
    │   └── linear.go
    ├── linreg
    │   ├── linreg.go
    │   └── matrix.go
    ├── logreg
    │   └── logreg.go
    ├── measure
    │   └── measure.go
    ├── pla
    │   └── pla.go
    ├── week1
    │   ├── week1.go
    │   ├── week1_test.go
    ├── week2
    │   └── week2.go
    ├── week3
    │   └── week3.go
    ├── week4
    │   └── week4.go
    ├── week5
    │   └── week5.go
    ├── week6
    │   └── week6.go
    └── week7
        └── week7.go



##Thoughts:

It might be better to divide the packages based on *models* and *methods*.
Here is how the topics are presented in the learning from data web page:
[topics](http://work.caltech.edu/library/)

###models:

* linear classification: PLA
* linear regression
* logistic regression
* non linear transformation
* neural networks
* support vector machines
* nearest neighbors


###methods:
* regularization
* validation
