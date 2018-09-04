<?php
include "ghca.php";
use ghca\Ghca;
$username = $_GET["username"];
$password = $_GET["password"];
echo Ghca::encrypt($username, $password);