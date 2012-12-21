# -*- ruby -*-
# There's no make on my Mac, but every Mac comes with Ruby... :)
PROJECTROOT=File.dirname(__FILE__)
#puts PROJECTROOT

# for verbose tests:
#TEST_VERBOSE="-v"
TEST_VERBOSE=""

PKGS="main pwdgen pwdgen/symbol pwdgen/rand"

task :default do
    with_env do 
        system("echo \"GOPATH is $GOPATH\"")
        puts "go build -o mpw main"
        system("go build -o mpw main")
        fail unless $? == 0
    end

end

task :test do
    with_env do 
        system("echo \"GOPATH is $GOPATH\"")
        puts "go test #{TEST_VERBOSE} #{PKGS}"
        system("go test #{TEST_VERBOSE} #{PKGS}")
        fail unless $? == 0
    end

end

task :http do
    with_env do
        puts "godoc -http=:6060 -path=\"#{PROJECTROOT}/src\""
        system("godoc -http=:6060 -path=\"#{PROJECTROOT}/src\"")
        fail unless $? == 0
    end
end

task :clean do
    File.unlink("#{PROJECTROOT}/mpw")
end

def with_env(&blk)
    oldpath=ENV["GOPATH"]
    ENV["GOPATH"]="#{PROJECTROOT}:#{oldpath}"
    puts "gopath: #{ENV['GOPATH']}"
    blk.call

    ENV["GOPATH"]=oldpath
end

