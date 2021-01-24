require 'benchmark'

arr = [1,2,3,4,5]

def ifelse(arr)
  if arr.count > 10
    'more than 10'
  elsif arr.count > 5
    'more than 5'
  else
    'equal or less than 5'
  end
end

def ifelse_with_var(arr)
  count = arr.count
  if count > 10
    'more than 10'
  elsif count > 5
    'more than 5'
  else
    'equal or less than 5'
  end
end

puts Benchmark.measure { 1000000.times{ ifelse(arr) } }
puts Benchmark.measure { 1000000.times{ ifelse_with_var(arr) } }
