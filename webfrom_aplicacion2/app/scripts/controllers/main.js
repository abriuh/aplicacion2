'use strict';

/**
 * @ngdoc function
 * @name webfromAplicacion2App.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the webfromAplicacion2App
 */
angular.module('webfromAplicacion2App')
  .controller('MainCtrl', function ($scope,$http) {
    
    $http.get('http:///localhost:8080/v1/factura')
    .then(function(response){
    	$scope.captura=response.data;
    })


    $http.get('http:///localhost:8080/v1/cliente')
    .then(function(response){
    	$scope.captura2=response.data;
    })

    $scope.envio={};

    $scope.registrar=function(){
        console.log("hola")
        $http.post('http://localhost:8080/v1/factura',$scope.envio)
        .then(function(response){
            console.log(response.envio)
        })
    }



  });
