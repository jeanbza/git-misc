class TodosController < ApplicationController
  def index
    # @todos = Todo.find(:all)
    @todos = Todo.all
  end

  def post
    print "\n\n\n\n\n\n\n"
    print params[:todo]
    print "\n\n\n\n\n\n\n"
    redirect_to({ action: 'index' }, notice: "Saved!")
  end
end