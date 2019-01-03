//SPDX-License-Identifier: Apache-2.0

var tuna = require('./controller.js');

module.exports = function(app){

  app.get('/get_cargo/:id', function(req, res){
    tuna.get_tuna(req, res);
  });
  app.get('/add_cargo/:cargo', function(req, res){
    tuna.add_tuna(req, res);
  });
  app.get('/get_all_cargo', function(req, res){
    tuna.get_all_tuna(req, res);
  });
  app.get('/change_holder/:holder', function(req, res){
    tuna.change_holder(req, res);
  });
}
