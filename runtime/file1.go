package main

func f1(skip int)(fileName ,funcName string ,line int){
	fileName, funcName, line = getLocation(skip)
	return
}
