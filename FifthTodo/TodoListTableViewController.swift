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
    
    override func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return todoItems.count
    }
    
    override func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell? {
        let cell = tableView.dequeueReusableCellWithIdentifier("ListPrototypeCell") as UITableViewCell
        let todoItem = todoItems[indexPath.row]
        cell.textLabel.text = todoItem.itemName
        
        if todoItem.completed {
            cell.accessoryType = UITableViewCellAccessoryType.Checkmark
        } else {
            cell.accessoryType = UITableViewCellAccessoryType.None
        }
        
        return cell;
    }
    
    override func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath: NSIndexPath) {
        let todoItem = todoItems[indexPath.row]
        todoItem.completed = !todoItem.completed
        tableView.deselectRowAtIndexPath(indexPath, animated: false)
        tableView.reloadRowsAtIndexPaths([indexPath], withRowAnimation:UITableViewRowAnimation.None)
    }
    
    func unwindToList(segue: UIStoryboardSegue) {
        println("Unwinding")
        let source = segue.sourceViewController as AddTodoViewController
        self.todoItems += source.todoItem
        println("Adding item \(source.todoItem.itemName)")
        self.tableView.reloadData()
    }
}
