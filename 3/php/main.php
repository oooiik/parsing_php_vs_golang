<?php

define("countNol", $argv[1] ?? die('input not found' . PHP_EOL));
const host = "mysql";
const user = "my";
const pass = "my";
const database = "php";

const table = "10e" . countNol . "-row";
const filePath = "/app/resources/" . table . ".csv";

function main()
{
    dbTruncate();

    $start = microtime(true);

    $conn = connectMysql(host, user, pass, database);

    $file = openFile(filePath);


    $row = 1;

    $per = 1000;

    $datas = [];
    while (($data = fgetcsv($file, 10000)) !== FALSE) {
        if ($row == 1) {
            $row++;
            continue;
        }
        $datas[] = $data;
        if (!($row % $per)) {
            dbInsert($conn, table, $datas);
            $datas = [];
        }
        $row++;
    }
    dbInsert($conn, table, $datas);


    fclose($file);

    $conn = null;

    $execution_time = (microtime(true) - $start) * 1000000;
    echo (10 ** countNol) . " rows: " . floor($execution_time) . " µs\n";
}

function connectMysql(string $host, string $user, string $pass, string $database): PDO
{
    return new PDO("mysql:host=$host;dbname=$database;charset=UTF8", $user, $pass);
}

function dbInsert(PDO $conn, $table, $values)
{
    $sql = "INSERT INTO `$table` (`uid`, `manufacturer_part_number`, `manufacturer`, `quantity`) VALUES ";

    foreach ($values as $key => $value) {
        $sql .= "(\"$value[0]\", \"$value[2]\", \"$value[3]\", \"$value[4]\")";

        if ($key !== array_key_last($values)) {
            $sql .= ", ";
        }
    }

    $conn->exec($sql);
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