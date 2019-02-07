## Hyperledger Fabric Sample Application

This application demonstrates the track and trace of Cargo & Containers along with change of Custody by leveraging Hyperledger Fabric in the supply chain. In this exercise we will set up the minimum number of nodes required to develop chaincode. It has a single peer and a single organization.

if getting error about running ./startFabric.sh permission 

chmod a+x startFabric.sh

This code is based on code written by the Hyperledger Fabric community. Source code can be found here: (https://github.com/hyperledger/fabric-samples). 

docker rm -f $(docker ps -aq)

./startFabric.sh

node registerAdmin.js
node registerUser.js

If Error, then -->    rm ~/.hfc-key-store/*

node registerAdmin.js
node registerUser.js

node server.js


