App = Ember.Application.create();

App.ApplicationController = Ember.Controller.extend({
    todoInput: "",
    todos: ['red', 'yellow', 'blue'],
    addTodo: function() {
        console.dir("boom");
    }
});
