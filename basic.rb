#! /usr/bin/env ruby

require 'agent'
require 'time'
require 'nokogiri'
require 'open-uri'

t = Time.now
DIR = t.strftime('%Y-%m')
Dir.mkdir(DIR) unless File.exists?(DIR)

# go routines
go! do
   puts dateString

   date = dateString
   filename = date + '.md'

   puts filename

   createMarkDown(date, filename)

   scrape("ruby", filename)
   scrape("go", filename)
   scrape("python", filename)
   scrape("php", filename)
   scrape("javascript", filename)

   gitPull()
   gitAddAll()
   gitCommit(date)
   gitPush()

	 sleep(1.days)
	 # or using time
	 # sleep(Time.parse("23:00:02") - Time.now)
	 puts "Sleeping one day"
end

def scrape(language, filename)
  target = File.open("#{DIR}/#{filename}", "a")

  target.write("\n####")

  target.write(language)

  target.write("\n\n")

  uri = "https://github.com/trending?l=#{language}"

  doc = Nokogiri::HTML(open(uri))
  rows = doc.css('ol.repo-list li')

  rows[0..-1].each do |row|
  
    hrefs = row.css("h3 a").map{ |a| 
      a['href']
    }.compact.uniq
   
    hrefs.each do |href|
      remote_url = "https://github.com"+href
      titles = row.css('h3 span').text()
      title = titles.delete "/"
      description = row.css('p').text()
      puts "Fetching #{remote_url}..."
      target.write("* [" + title.strip! + "](" + remote_url + "): " + description.strip! + "\n")
    end # done: hrefs.each

  end # done: rows.each
  target.close
end

def gitPull
  system "git pull origin master"
end

def gitAddAll
  system "git add -A"
end

def gitCommit(date)
  system "git commit -am #{date}"
end

def gitPush
  system "git push origin master"
end

def dateString
  t = Time.now
  return t.strftime('%Y-%m-%d')
end

def createMarkDown(date, filename)
  s = "###" + date + "\n"
  local_fname = "#{DIR}/#{filename}"
  open(local_fname, 'w'){ |file|
    file.write(s)
  }
end

loop do
end
