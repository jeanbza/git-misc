//
//  TodoListTableViewController.swift
//  FifthTodo
//
//  Created by Jean de Klerk on 7/12/14.
//  Copyright (c) 2014 Jean de Klerk. All rights reserved.
//

import UIKit

@objc(TodoListTableViewController) class TodoListTableViewController: UITableViewController {
    var todoItems: TodoItem[] = []
    
    init(coder aDecoder: NSCoder!) {
        super.init(coder: aDecoder)
    }
    
    func loadInitialData() {
        todoItems += [TodoItem(itemName: "first todo")]
        todoItems += [TodoItem(itemName: "second todo")]
    }
    
    override func viewDidLoad() {
        super.viewDidLoad()
        loadInitialData()
    }
    
    override func numberOfSectionsInTableView(tableView: UITableView!) -> Int {
        return 1
    }
    
    override func tableView(tableView: UITableView?, numberOfRowsInSection section: Int) -> Int {
        return todoItems.count
    }
    
    override func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell? {
        let cell = tableView.dequeueReusableCellWithIdentifier("ListPrototypeCell") as UITableViewCell
        let todoItem = todoItems[indexPath.row]
        cell.textLabel.text = todoItem.itemName
        return cell;
    }
}
