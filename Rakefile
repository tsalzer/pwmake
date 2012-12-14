# -*- ruby -*-
# There's no make on my Mac, but every Mac comes with Ruby... :)
PROJECTROOT=File.dirname(__FILE__)

task :default do
    with_env do 
        puts "go build -o mpw main"
        `go build -o mpw main`
    end

end

task :http do
    with_env do
        puts "godoc -http=:6060 -path=\"#{PROJECTROOT}/src\""
        `godoc -http=:6060 -path="#{PROJECTROOT}/src"`
    end
end

def with_env(&blk)
    oldpath=ENV["GOPATH"]
    ENV["GOPATH"]="#{PROJECTROOT}:#{oldpath}"

    blk.call

    ENV["GOPATH"]=oldpath
end

