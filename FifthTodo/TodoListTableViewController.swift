//
//  TodoListTableViewController.swift
//  FifthTodo
//
//  Created by Jean de Klerk on 7/12/14.
//  Copyright (c) 2014 Jean de Klerk. All rights reserved.
//

import UIKit

@objc(TodoListTableViewController) class TodoListTableViewController: UITableViewController {
    var todoItems: [TodoItem]
    
    init(coder aDecoder: NSCoder!) {
        super.init(coder: aDecoder)
    }
    
    func viewDidLoad() {
        todoItems += [TodoItem]
    }
}
