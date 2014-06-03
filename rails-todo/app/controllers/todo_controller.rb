class TodoController < ApplicationController
  before_action :set_user, only: [:show, :edit, :update, :destroy]

  def index
    @todos = ["bam", "boom"]
  end
end