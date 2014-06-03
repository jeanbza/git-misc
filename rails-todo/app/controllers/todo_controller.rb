class TodoController < ApplicationController
  before_action :set_user, only: [:show, :edit, :update, :destroy]

  def index
    @todos = [Todo.new("todo 1", "some content"), Todo.new("todo 2", "some more content")]
  end
end

class Todo
  def initialize(title, content)
    @Title = title
    @Content = content
  end

  def Title
    @Title
  end

  def Content
    @Content
  end
end