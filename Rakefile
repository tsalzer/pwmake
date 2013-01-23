# -*- ruby -*-
# There's no make on my Mac, but every Mac comes with Ruby... :)
PROJECTROOT=File.dirname(__FILE__)
#puts PROJECTROOT

PREFIX="/usr/local"
DST_BIN="#{PREFIX}/bin"

# for verbose tests:
#TEST_VERBOSE="-v"
TEST_VERBOSE=""

TPKGS="pwdgen pwdgen/symbol pwdgen/rand pwdgen/screen"
PKGS="main #{TPKGS}"

desc "default task"
task :default => :build

desc "build #{TARGET} binary"
task :build do
    run_go("build -o #{TARGET} main")
end

desc "run all tests in #{TPKGS}"
task :test do
    run_go("test #{TEST_VERBOSE} #{TPKGS}")
end

desc "run all benchmarks in #{TPKGS}"
task :bench do
    run_go("test -test.bench 'Benchmark.*' #{TEST_VERBOSE} #{TPKGS}")
end

desc "build documentation from #{PKGS}"
task :doc do
    run_go("doc #{TEST_VERBOSE} #{PKGS}")
end

#desc "install #{TARGET} to #{DST_BIN}"
#task :install => [:build, :test] do
#    puts "do as I say: cp -p #{TARGET} #{DST_BIN}"
#end

desc "run a documentation Webserver on port 6060"
task :http do
    run_line("godoc -http=:6060 -path=\"#{PROJECTROOT}/src\"")
end

desc "remove #{TARGET} from #{PROJECTROOT}"
task :clean do
    File.unlink("#{PROJECTROOT}/#{TARGET}")
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

