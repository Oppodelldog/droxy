#!/usr/bin/env bats

testRoot="/go/src/github.com/Oppodelldog/droxy/.test"
configFile="${testRoot}/droxy.toml"
droxy="${testRoot}/droxy"
logfile="${testRoot}/droxy.log"
fileCreationTestDir="${testRoot}/test"

function setupTestFileCreation(){
    rm -rf ${fileCreationTestDir}
    mkdir ${fileCreationTestDir}
    cd ${fileCreationTestDir}||exit
    ls |wc -l||exit
    export DROXY_CONFIG=${configFile}
}

function setupTestWithCommands(){
    setupTestFileCreation
    $droxy clones
}

function teardown(){
    rm -rf ${fileCreationTestDir}
    rm -f ${logfile}
}

@test "droxy" {
  $droxy
  
  [ $? -eq 0 ]
}

@test "droxy clones - executes without error" {
    setupTestFileCreation
    
    $droxy clones
    
  [ $? -eq 0 ]
}

@test "droxy clones - creates binary" {
    setupTestFileCreation
    
    $droxy clones
    
  [ $(ls outputs_test123|wc -l) -eq 1 ]
}

@test "droxy clones - does not overwrite existing binary" {
    setupTestFileCreation
    
    fileName="outputs_test123"
    touch $fileName
    fileSizeBefore=$(stat -c%s $fileName)
    $droxy clones
    fileSizeAfter=$(stat -c%s $fileName)  
      
    [ $fileSizeBefore -eq $fileSizeAfter ]
    [ $fileSizeBefore -eq 0 ]
}

@test "droxy clones -f - overwrites existing binary" {
    setupTestFileCreation
    
    fileName="outputs_test123"
    touch $fileName
    fileSizeBefore=$(stat -c%s $fileName)
    $droxy clones -f
    fileSizeAfter=$(stat -c%s $fileName) 
       
    [ $fileSizeBefore -ne $fileSizeAfter ]
    [ $fileSizeBefore -eq 0 ]
}

@test "droxy - binary outputs the expected" {
    setupTestWithCommands

    [[ $(./outputs_test123) == "test123"* ]]
}
