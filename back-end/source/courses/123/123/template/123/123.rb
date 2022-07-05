require "AssessmentBase.rb"

module 123
  include AssessmentBase

  def assessmentInitialize(course)
    super("123",course)
    @problems = []
  end

end
