## CREATING YOUR OWN RAILS SKELETON TO MESS AROUND IN

PLEASE NOTE: You will need some of the stuff below (esp the gems if you don't have them) to create projects,
so probably make sure you have all of those before you come back and make your own project with these two things.

- gem install skeleton-rails
- http://railswizard.org/ (this is awesome! and what I used to generate this project.)

## USEFUL LINKS

- There are none. No, really. Ok, fine, here's what I used, but I promise you they all suck
- https://gist.github.com/zhengjia/428105: a capybara cheatsheet. Except that half the stuff is out of date, so that's pretty useless
- The 'getting started' guide I followed. Except that it doesn't work, so that's also useless. I ended up taking the idea presented there and turning it into what exists now
- General reading at https://github.com/cucumber/cucumber/wiki

## YOU WILL NEED

#### xCode
Go to my applications and delete xcode if you have it. The latest version doesn't work (go figure). Then, go 
to https://developer.apple.com/downloads/index.action and download 4.6.3 (and install it as normal).

#### Nokogiri

Nokogiri is a PAIN IN THE ASS. Here is the install page: http://nokogiri.org/tutorials/installing_nokogiri.html. 
To install this, I recommend you take the macports route. There are other choices on that page. Here is the macports
route:

- Install macports https://www.macports.org/
- sudo port install libxml2 libxslt
- sudo gem install nokogiri

(if you have problems do sudo port upgrade outdated)

Again, nokogiri is a pain, so don't be surprised if something goes wrong. IF YOU GET WEIRD NOKOGIRI ERRORS WHEN 
RUNNING CUCUMBER (not during installation), RUN THE FOLLOWING: sudo bundle exec gem pristine nokogiri.

#### My Gem Environment

These are all things you're going to want to gem install:
- actionmailer (4.1.0)
- actionpack (4.1.0)
- actionview (4.1.0)
- activemodel (4.1.0)
- activerecord (4.1.0)
- activesupport (4.1.0)
- addressable (2.3.6)
- akephalos2 (2.1.3)
- arel (5.0.1.20140414130214)
- builder (3.2.2)
- bundler (1.6.2)
- capybara (2.2.1)
- coffee-rails (4.0.1)
- coffee-script (2.2.0)
- coffee-script-source (1.7.0)
- cucumber (1.3.14)
- cucumber-rails (1.4.0)
- database_cleaner (1.2.0)
- diff-lcs (1.2.5)
- domain_name (0.5.18)
- erubis (2.7.0)
- execjs (2.0.2)
- gherkin (2.12.2)
- hike (2.1.3, 1.2.3)
- http-cookie (1.0.2)
- i18n (0.6.9)
- jbuilder (2.0.6)
- jquery-rails (3.1.0)
- jruby-jars (1.7.12)
- json (1.8.1)
- launchy (2.4.2)
- libxml-ruby (2.7.0)
- mail (2.5.4)
- mechanize (2.7.3, 2.7.2)
- mime-types (2.2, 1.25.1)
- mini_portile (0.6.0, 0.5.3)
- minitest (5.3.3)
- multi_json (1.9.3)
- multi_test (0.1.1)
- net-http-digest_auth (1.4)
- net-http-persistent (2.9.4)
- nokogiri (1.6.1)
- ntlm-http (0.1.1)
- polyglot (0.3.4)
- rack (1.5.2)
- rack-test (0.6.2)
- rails (4.1.0)
- railties (4.1.0)
- rake (10.3.1)
- rdoc (4.1.1)
- rspec (2.14.1)
- rspec-core (2.14.8)
- rspec-expectations (2.14.5)
- rspec-mocks (2.14.6)
- ruby (0.1.0)
- rvm (1.11.3.9)
- sass (3.3.6, 3.2.19)
- sass-rails (4.0.3)
- sdoc (0.4.0)
- skeleton-rails (0.1.0)
- spring (1.1.2)
- sprockets (2.12.1, 2.11.0)
- sprockets-rails (2.1.3)
- sqlite3 (1.3.9)
- thor (0.19.1)
- thread_safe (0.3.3)
- tilt (2.0.1, 1.4.1)
- treetop (1.5.3, 1.4.15)
- turbolinks (2.2.2)
- tzinfo (1.1.0)
- uglifier (2.5.0)
- unf (0.1.4)
- unf_ext (0.0.6)
- webrat (0.7.3)
- webrobots (0.1.1)
- xpath (2.0.0)


