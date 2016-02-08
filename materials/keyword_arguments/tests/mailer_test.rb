require_relative '../mailer'
require 'minitest/autorun'

class TestMailer < Minitest::Test
  def setup
    @mailer = Mailer.new
  end

  def test_that_subject_is_string
    assert_kind_of String, @mailer.subject
  end
end
