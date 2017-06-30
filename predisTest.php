<html>
<body>
<?php
echo 'helloworld';
$dollar = 15.23 + 7.45;
require 'predis/src/Autoloader.php';
Predis\Autoloader::register();
$client = new Predis\Client([
    'scheme' => 'tcp',
    'host'   => '127.0.0.1',
    'port'   => 8037
]);
$client->auth('IlikeredFROGS865IhateredFROGS865');
//$client->set('foo', 'redMonkeysLikeRedis');
$value = $client->get('foo');
echo "<br>";
echo $dollar;
echo $value;
?>

<p>this is a p element</p>
</body>
</html>
