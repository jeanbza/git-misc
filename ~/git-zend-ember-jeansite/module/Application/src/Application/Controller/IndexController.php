<?php

namespace Application\Controller;

use Zend\Mvc\Controller\AbstractActionController;
use Zend\View\Model\ViewModel;
use Zend\Db\Adapter\Adapter;
use Zend\Db\Adapter\Driver\ResultInterface;
use Zend\Db\ResultSet\ResultSet;

class IndexController extends AbstractActionController
{
    public function indexAction()
    {
        $adapter = new Adapter(array(
            'driver' => 'Mysqli',
            'database' => 'jeansite',
            'username' => 'root',
            'password' => '',
            'options' => array('buffer_results' => true)
        ));

        $sql = "SELECT * FROM jeansite.post";

        $statement = $adapter->createStatement($sql);
        $result = $statement->execute();

        if ($result instanceof ResultInterface && $result->isQueryResult()) {
            $resultSet = new ResultSet;
            $resultSet->initialize($result);
        }

        return new ViewModel(array(
            "test" => "This is test data from the controller",
            "db" => $resultSet->toArray()
        ));
    }
}
