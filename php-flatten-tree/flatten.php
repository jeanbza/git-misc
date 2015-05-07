<?php
include 'data_array.php';

$flatArr = array();

function flatten($arr) {
    global $flatArr;

    foreach ($arr as $node) {
        array_push($flatArr, $node['Unit']['name']);

        if (!empty($node['children'])) {
            flatten($node['children']);
        }
    }
}

flatten($units);
print_r($flatArr);