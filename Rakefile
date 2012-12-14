# -*- ruby -*-
# There's no make on my Mac, but every Mac comes with Ruby... :)
PROJECTROOT=File.dirname(__FILE__)

task :default do
	oldpath=ENV["GOPATH"]
	ENV["GOPATH"]="#{PROJECTROOT}:#{oldpath}"
	puts "go build -o mpw main"
	`go build -o mpw main`
	ENV["GOPATH"]=oldpath
end

