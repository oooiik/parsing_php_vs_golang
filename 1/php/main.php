<?php

const host = "mysql";
const user = "my";
const pass = "my";
const database = "php";
const table = "10e6-row";
const filePath = "/app/resources/" . table . ".csv";


function main()
{
    dbTruncate();

    $start = microtime(true);

    $conn = connectMysql(host, user, pass, database);

    $file = openFile(filePath);


    $row = 1;
    while (($data = fgetcsv($file)) !== FALSE) {
        if ($row == 1) {
            $row++;
            continue;
        }
        dbInsert($conn, table, $data);
        $row++;
    }


    fclose($file);

    $conn = null;

    $execution_time = (microtime(true) - $start) * 1000000;
    echo floor($execution_time) . " µs\n";
}

function connectMysql(string $host, string $user, string $pass, string $database): PDO
{
    return new PDO("mysql:host=$host;dbname=$database;charset=UTF8", $user, $pass);
}

function dbInsert(PDO $conn, $table, $value)
{
    $sql = "INSERT INTO `$table` (`uid`, `manufacturer_part_number`, `manufacturer`, `quantity`) VALUES ('$value[0]', '$value[2]', '$value[3]', '$value[4]')";

    $conn->query($sql);
}

function dbTruncate()
{
    $conn = connectMysql(host, user, pass, database);
    if (!$conn->query("TRUNCATE `" . table . "`")) {
        die("Error: " . $conn->error . PHP_EOL);
    }
}

function openFile($path)
{
    $file = fopen($path, "r");
    if ($file === false) {
        die("Open file!" . PHP_EOL);
    }
    return $file;
}


main();