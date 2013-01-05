# -*- ruby -*-
# There's no make on my Mac, but every Mac comes with Ruby... :)
PROJECTROOT=File.dirname(__FILE__)
#puts PROJECTROOT

# for verbose tests:
#TEST_VERBOSE="-v"
TEST_VERBOSE=""

PKGS="main pwdgen pwdgen/symbol pwdgen/rand"

task :default do
    run_go("build -o mpw main")
end

task :test do
    run_go("test #{TEST_VERBOSE} #{PKGS}")
end

task :doc do
    run_go("doc #{TEST_VERBOSE} #{PKGS}")
end


task :http do
    run_line("godoc -http=:6060 -path=\"#{PROJECTROOT}/src\"")
end

task :clean do
    File.unlink("#{PROJECTROOT}/mpw")
end

def run_line(line)
    with_env do
        #system("echo \"GOPATH is $GOPATH\"")
        puts "#{line}"
        system("#{line}")
        fail unless $? == 0
    end
end

def run_go(args)
    run_line("go #{args}")
end

def with_env(&blk)
    oldpath=ENV["GOPATH"]
    ENV["GOPATH"]="#{PROJECTROOT}:#{oldpath}"
    #puts "gopath: #{ENV['GOPATH']}"
    blk.call

    ENV["GOPATH"]=oldpath
end

