
	##################################################
	                                                 #
	#  author:  Thyme Arc                            
	#  email: thymearc@gmail.com
	#  licence: GNU 3.1
	#  copyright:  Myarkadia Ltd @ all rights reserved
	#
	##################################################



	1. What is it?
	
		a service written in GO.


	2. requirements
	
		do:

			sh env.sh	
			sh dependencies.sh

	3. how to use it?
		enter server folder	
	        you may have to regenerate service with this command:	
		protoc -I CalcService CalcService/CalcService.proto --go_out=plugins=grpc:CalcService
	
	        test it:
		go test main_test.go

	4. Documentation

		grpc api is described in Doc/api.txt
