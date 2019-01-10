//SPDX-License-Identifier: Apache-2.0

var cargo = require('./controller.js');
var bodyParser    = require('body-parser');
var jsonParser = bodyParser.json()
module.exports = function(app){

  app.post('/registerParticipant/', jsonParser, function(req, res){
    cargo.registerParticipant(req, res);
  });
  app.post('/addNewContainer/', jsonParser, function(req, res){
    cargo.addNewContainer(req, res);
  });
  app.get('/getParticipants/:key', function(req, res){
    cargo.getParticipants(req, res);
  });
  app.put('/loadContainerWithPackages/', jsonParser, function(req, res){
    cargo.loadContainerWithPackages(req, res);
  });
  app.post('/createCargoLoadContainers/', jsonParser, function(req, res){
    cargo.createCargoLoadContainers(req, res);
  });
  app.put('/updateCargoAttributes/', jsonParser, function(req, res){
    cargo.updateCargoAttributes(req, res);
  });
  app.put('/updateCargoCoordinates/', jsonParser, function(req, res){
    cargo.updateCargoCoordinates(req, res);
  });
  app.put('/changeCargoCustody/', jsonParser, function(req, res){
    cargo.changeCargoCustody(req, res);
  });
  app.put('/updateContainerAttributes/', jsonParser, function(req, res){
    cargo.updateContainerAttributes(req, res);
  });
  app.put('/changeContainerCustody/', jsonParser, function(req, res){
    cargo.changeContainerCustody(req, res);
  });
  app.put('/unloadContainerFromCargo/', jsonParser, function(req, res){
    cargo.unloadContainerFromCargo(req, res);
  });
  app.get('/traceCargo/:cargoId', function(req, res){
    cargo.traceCargo(req, res);
  });
  app.get('/trackCargoDetails/:cargoId', function(req, res){
    cargo.trackCargoDetails(req, res);
  });
  app.get('/traceContainer/:containerId', function(req, res){
    cargo.traceContainer(req, res);
  });
  app.get('/trackContainerDetails/:containerId', function(req, res){
    cargo.trackContainerDetails(req, res);
  });
  app.get('/getLoadedContainers/', function(req, res){
    cargo.getLoadedContainers(req, res);
  });
  app.get('/getAvilableContainers/', function(req, res){
    cargo.getAvilableContainers(req, res);
  });
}
