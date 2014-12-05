package main

import (
    "launchpad.net/goamz/aws"
    "launchpad.net/goamz/s3"
    "io/ioutil"
)

func main() {
    auth, err := aws.EnvAuth()
    if err != nil {
        println("EnvAuth")
        panic(err.Error())    
    }
    s := s3.New(auth, aws.USWest2)

    bucket := s.Bucket("test-aws-sjpuas");

    println("==== list files ====") 
    list, err := bucket.List("", "", "", 10);
    if err != nil {
        println("List")
        panic(err.Error())    
    }

    println("bucket:"+list.Name)   
    println("Archivos:")
    for _, val := range list.Contents {
        println("\t"+val.Key)
    }   

    
    println("==== get file js/main.js ====") 
    data, err := bucket.Get("js/main.js")
    if err != nil {
        panic(err.Error())    
    } 
    
    err = ioutil.WriteFile("main.js", data, 0x777)
    if err != nil {
        panic(err)
    }
    

    println("==== put new file screen.png ====") 
    data, err = ioutil.ReadFile("screen.png")

    err = bucket.Put("images/1235/screen.png",data,"image/png","")
    if err != nil {
        panic(err)
    }

}