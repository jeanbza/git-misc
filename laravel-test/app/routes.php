<?php

Route::get('/', function()
{
    $res = DB::select('select * from test_db');
    echo "<pre>";print_r($res);echo "</pre>";
    return View::make('hello', array('name' => 'Taylor'));
});