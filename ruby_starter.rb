str = 'Hello World!'

# Output string
puts str

# arguments
puts "$0: #{$0}"

# each arguments
ARGV.each_with_index do |arg, i|
  puts "ARGV[#{i}]: #{arg}"
end

takahashi, nakahashi, hikuhashi, k = STDIN.gets.chomp.split(' ').map{|n| n.to_i}

# multiple line
# attrs = Array.new
# while line = $stdin.gets
#   attrs << line.chomp.split
# end

puts $stdin.read.split("\n")
