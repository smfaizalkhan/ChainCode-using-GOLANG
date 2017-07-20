# ChainCode-using-GOLANG

Sublime Text editor to write chaincode

sudo add-apt-repository ppa:webupd8team/sublime-text-3

sudo apt-get update

sudo apt-get install sublime-text-installer

Package Installer

https://packagecontrol.io/installation

Go to installed Packager folder and get the Package controll sublime package using cURL

wget https://packagecontrol.io/Package%20Control.sublime-package

Then restart sublime text

Press Ctrl+Shift+P to install the downloaded Package into Sublime

Type go sublime to install in Sublime editor

Once installed check the Preference --> Package Setting -->GO Sublime and edit the env variable to the GOPATH

Restart the Editor

Then select the Tools-->Build Systems-->GO Loang



Writing chaincode and unit testing  in GoLang

 Testing a chaincode is a cumber-some process as we need to deploy the chaincode in the hyperledger in a docker conatiner.
 To over come that ,we can do an unit testing before deploying using the Fabric provided CustomMockStub and testing package

File naming for test file shoud be *_test.go

 Any Test Function  in should start with TestXXX() followed by name of the function.

  Create a custom MockStub that internally uses shim.MockStub

   cc := new(SampleChainCode)
	stub := shim.NewMockStub("InitTest", cc)
	
Run the	 test using the command using go test

 ![screenshot from 2017-07-20 13-06-22](https://user-images.githubusercontent.com/22238550/28409641-3f2d9e70-6d4c-11e7-9cdc-7539cde84eee.png)

  
	



