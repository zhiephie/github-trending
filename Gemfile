source "http://rubygems.org"
ruby '2.3.3'

git_source(:github) do |repo_name|
  repo_name = "#{repo_name}/#{repo_name}" unless repo_name.include?("/")
  "https://github.com/#{repo_name}.git"
end

gem 'agent',   '0.12.0'
gem 'nokogiri', github: 'sparklemotion/nokogiri'