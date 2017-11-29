#! /usr/bin/env ruby

require 'time'
require 'nokogiri'
require 'open-uri'

t = Time.now
DIR = t.strftime('%Y-%m')
Dir.mkdir(DIR) unless File.exists?(DIR)

def job 

   puts 'Starting...'

   date = dateString
   filename = date + '.md'

   createMarkDown(date, filename)
  
   scrape('elixir', filename)
   scrape('erlang', filename)
   scrape('ruby', filename)
   scrape('go', filename)
   scrape('python', filename)
   scrape('php', filename)
   scrape('javascript', filename)
   scrape('java', filename)

   git_add_commit_push(date)
   
   puts 'Done.!'
end

def scrape(language, filename)
  local_fname = "#{DIR}/#{filename}"
  target = File.open(local_fname, "a")

  target.write("\n####")

  target.write(language)

  target.write("\n\n")

  uri = "https://github.com/trending?l=#{language}"

  doc = Nokogiri::HTML(open(uri))
  rows = doc.css('ol.repo-list li')

  rows.each do |row|
  
    hrefs = row.css("h3 a").map{ |a| 
      a['href']
    }.compact.uniq
   
    hrefs.each do |href|
      remote_url = "https://github.com"+href
      title = row.css('h3 a').text()
      description = row.css('p').text()
      puts "Fetching #{remote_url}..."
      target.write("* [#{title.strip!}" "](#{remote_url}" "):  #{description.strip!}" "\n")
      #target.write("* [" + title.strip! + "](" + remote_url + "): " + description.strip! + "\n")
    end # done: hrefs.each

  end # done: rows.each
  target.close
end

def git_add_commit_push(date)
  system "git add -A"
  system "git commit -m #{date}"
  system "git push -u origin master"
end

def dateString
  t = Time.now
  return t.strftime('%Y-%m-%d')
end

def createMarkDown(date, filename)
  s = "###" + date + "\n"
  local_fname = "#{DIR}/#{filename}"
  target = File.open(local_fname, 'w')
  target.write(s)
  target.close
end

loop do
  job()
  sleep(86400)
  # or using time
  # sleep(Time.parse("22:00:00") - Time.now)
end

