<?php
	//get predis running
	require 'predis/src/Autoloader.php';
	Predis\Autoloader::register();
    	$client = new Predis\Client([
    'scheme' => 'tcp',
    'host'   => '127.0.0.1',
    'port'   => 8037
]);
	$client->AUTH("IlikeredFROGS865IhateredFROGS865");
	
	//get the json data
	$requestJsonData = file_get_contents('php://input');
	$requestJsonData = json_decode($requestJsonData);
	
	//get id for the redis queue
	$requestIdNumber = $client->INCR("requestsCount");
    
    //using reference lookups to deal with the fact that redis doesn't allow for nested tables
	//set the references
    $client->LPUSH("queueList","request$requestIdNumber");
    $client->HMSET("request$requestIdNumber","endpoint","endpoint$requestIdNumber","data","data$requestIdNumber");
	//plug in the request information
    $client->HMSET("endpoint$requestIdNumber","method",$requestJsonData->endpoint->method,"url",$requestJsonData->endpoint->url);
    foreach($requestJsonData->data as $arrayIndex) {
       foreach ($arrayIndex as $key => $stringValue) {
           $client->HSET("data$requestIdNumber",$key,$stringValue);
       }
    }
    unset($key);
    unset($value);
	
?>
