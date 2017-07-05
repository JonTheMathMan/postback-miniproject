<html>
<body>
	<?php
		if($_GET["title"]!="" && $_GET["image"]!="") {
			echo '<h1>'.$_GET["title"].'</h1>';
    			echo '<img src='.$_GET["image"].'/>';
		}

		if($_POST!=null) {
			foreach($_POST as $key => $value) {
				echo "<h1>$key</h1>";
				echo "<p>$value</p>";
			}
		unset($key);
		unset($value);
		}
	?>
</body>
</html>
