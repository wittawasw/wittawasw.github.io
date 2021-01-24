require_relative '../post'
require 'minitest/autorun'

class TestPost < Minitest::Test
  def setup
    @post = Post.new(subject: 'RegularReminder',
                     body: 'This is the message',
                    )
  end

  def test_subject_is_string
    assert_kind_of String, @post.subject
  end

  def test_body_is_string
    assert_kind_of String, @post.body
  end

  def test_error_on_new_without_subject
    assert_raises(StandardError) { Post.new(body: 'text') }
  end

  def test_error_on_new_without_body
    assert_raises(StandardError) { Post.new(subject: 'text') }
  end
end
