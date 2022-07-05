require "AssessmentBase.rb"

module Temp
  include AssessmentBase

  def assessmentInitialize(course)
    super("temp",course)
    @problems = []
  end

end
