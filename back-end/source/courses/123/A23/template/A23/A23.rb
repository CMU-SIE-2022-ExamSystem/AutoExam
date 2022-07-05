require "AssessmentBase.rb"

module A23
  include AssessmentBase

  def assessmentInitialize(course)
    super("A23",course)
    @problems = []
  end

end
