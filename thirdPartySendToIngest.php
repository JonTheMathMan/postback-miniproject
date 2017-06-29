<?php

    $ch = curl_init();
    
    $data = "{\"endpoint\":{ \"method\":\"GET\",\"url\":\"http://127.0.0.1/thirdPartyReceiveFromDelivery.php/data?title={mascot}&image={location}&foo={bar}\" },\"data\":[ {\"mascot\":\"Gopher\",\"location\":\"https://blog.golang.org/gopher/gopher.png\" }] }";
    
    curl_setopt($ch, CURLOPT_URL,            "http://127.0.0.1/phptest/mockIngest.php" );
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1 );
    curl_setopt($ch, CURLOPT_POST,           1 );
    curl_setopt($ch, CURLOPT_POSTFIELDS, $data); 
    curl_setopt($ch, CURLOPT_HTTPHEADER,     array('Content-type: application/json'));

    $result= curl_exec ($ch);
    
    echo $result;

?>
