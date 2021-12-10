# SC Project 3

## structures

- gateway
- dev
- sensor

**see usercases folder for more details**

## run

``` bash
# compile and upload the binary to remote

# compile
make

# upload
scp -J macneill build/project3.tar pi19:~  
``` 

``` bash
# run this following code remotely

# upzip
tar -xf project3.tar 

# move to the folder which includes binary and run.sh
cd ./build/package/

# run <working dir> <local> <init index>
# see run.sh in ./scripts folder

# on pi19
./run.sh ~/pi19 rasp-019.scss.tcd.ie rasp-020.scss.tcd.ie

# on pi20
./run.sh ~/pi20 rasp-020.scss.tcd.ie rasp-019.scss.tcd.ie
```