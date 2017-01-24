#%Module 1.0

eval set [ array get env HOME ]

set BASE_PATH = $HOME/sage-6.4.1-x86_64-Linux
setenv SAGE_HOME $BASE_PATH
prepend-path PATH $BASE_PATH
