//
//  AddTodoViewController.swift
//  FifthTodo
//
//  Created by Jean de Klerk on 7/12/14.
//  Copyright (c) 2014 Jean de Klerk. All rights reserved.
//

import UIKit

class AddTodoViewController : UIViewController {
    var todoItem:TodoItem = TodoItem(itemName: "")
    
    init(coder aDecoder: NSCoder!) {
        super.init(coder: aDecoder)
    }
    
    override func prepareForSegue(segue: UIStoryboardSegue, sender: AnyObject) {
        if sender as UIBarButtonItem != self.doneButton {
            return
        }
        
        if self.textField.text.utf16count > 0 {
            self.todoItem = TodoItem(itemName: self.textField.text)
        }
    }
    
    
    @IBOutlet var textField : UITextField
    
    @IBOutlet var doneButton : UIBarButtonItem
}