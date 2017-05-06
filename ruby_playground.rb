
require "zlib"

require 'fileutils'

# Create content object
content = "what is up, doc?"
header = "blob #{content.length}\0"
store = header + content
require 'digest/sha1'
sha1 = Digest::SHA1.hexdigest(store)
digest  = Digest::SHA1.digest(store)
zlib_content = Zlib::Deflate.deflate(store)
path = '.git/objects/' + sha1[0,2] + '/' + sha1[2,38]
FileUtils.mkdir_p(File.dirname(path))
File.open(path, 'w') { |f| f.write zlib_content }


# Create tree object
#zi = Zlib::Inflate.new
#f1 = File.open(".git/objects/c8/2b154409d60ff285aacd55ff340c0fbb0901d2")
#x1 = f1.read
#s1 = zi.inflate(x1)
content = "100644 .gitignore"+ "\x00" + digest #.map { |b| sprintf(", 0x%02x",b) }.join

store = "tree #{content.length}\0" + content
sha1 = Digest::SHA1.hexdigest(store)
zlib_content = Zlib::Deflate.deflate(store)
# path = ".git/objects/bd/9dbf5aae1a3862dd1526723246b20206e5fc33"
path = '.git/objects/' + sha1[0,2] + '/' + sha1[2,38]

puts path
FileUtils.mkdir_p(File.dirname(path))
File.open(path, 'w') { |f| f.write zlib_content }
#f2 = File.open(path)
#x2 = f2.read
#s2 = zi.inflate(x2)

# Create commit
# echo 'second commit' | git commit-tree 6398107bf9b91a96e55e90959994958705325d06 -p 519b6efacb56a32930a7b22dafe903aea8a76114
#echo "71ba5227f9b7b9d75b7a5a7904486b78000f0318" > .git/refs/heads/master
# commit {size}\0
str = <<STR
tree 6398107bf9b91a96e55e90959994958705325d06
parent 519b6efacb56a32930a7b22dafe903aea8a76114
author Alexandr Kostrikov <alexandr.kostrikov@gmail.com> 1494000784 +0300
committer Alexandr Kostrikov <alexandr.kostrikov@gmail.com> 1494000784 +0300

commit message
STR
puts str
store = "commit #{str.length}\0" + str
puts store
sha1 = Digest::SHA1.hexdigest(store)
puts sha1
zlib_content = Zlib::Deflate.deflate(store)
# path = ".git/objects/bd/9dbf5aae1a3862dd1526723246b20206e5fc33"
path = '.git/objects/' + sha1[0,2] + '/' + sha1[2,38]

puts path
FileUtils.mkdir_p(File.dirname(path))
File.open(path, 'w') { |f| f.write zlib_content }
# echo "41f3e1eac8f408c95562601ea14d0c0587112cc7" > .git/refs/heads/master
# The check
=begin
zi = Zlib::Inflate.new
f1 = File.open(".git/objects/71/ba5227f9b7b9d75b7a5a7904486b78000f0318")
x1 = f1.read

puts
puts "x1"
puts x1
s1 = zi.inflate(x1)
puts
puts "s1"
puts s1
=end