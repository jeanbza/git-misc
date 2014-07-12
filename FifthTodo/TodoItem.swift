//
//  TodoItem.swift
//  FifthTodo
//
//  Created by Jean de Klerk on 7/12/14.
//  Copyright (c) 2014 Jean de Klerk. All rights reserved.
//

import Foundation

class TodoItem {
    var itemName: String = ""
    var completed: Bool = false
    var creationDate: NSDate
    
    init() {
        self.creationDate = NSDate()
    }
}