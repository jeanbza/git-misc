//
//  ViewController.swift
//  swift-clock
//
//  Created by Jean de Klerk on 9/14/14.
//  Copyright (c) 2014 Jean de Klerk. All rights reserved.
//

import UIKit

class ViewController: UIViewController {
    
    @IBOutlet var myLabel : UILabel!

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view, typically from a nib.
        
        updateTime()
        
        var timer = NSTimer()
        let aSelector : Selector = "updateTime"
        timer = NSTimer.scheduledTimerWithTimeInterval(0.01, target: self, selector: aSelector, userInfo: nil, repeats: true)
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

    func updateTime() {
        var date = NSDate()
        var calendar = NSCalendar.currentCalendar()
        var components = calendar.components(.CalendarUnitHour | .CalendarUnitMinute | .CalendarUnitSecond, fromDate: date)
        var hour = components.hour
        var minutes = components.minute
        var seconds = components.second
        
        myLabel.text = "The time is \(hour):\(minutes):\(seconds)"
    }
}