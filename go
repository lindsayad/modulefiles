#%Module 1.0

eval set [ array get env HOME ]
setenv GOROOT	 $HOME/goWorkspace/go
setenv GOPATH 	 $HOME/goWorkspace
prepend-path PATH	$HOME/goWorkspace/go/bin
prepend-path PATH	$HOME/goWorskapce/bin