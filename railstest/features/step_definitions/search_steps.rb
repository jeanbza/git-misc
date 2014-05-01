Given /^I have opened "([^\"]*)"$/ do |url|
  visit url
end

When /^I click "([^\"]*)"$/ do |text|
  click_link(text)
end

Then /^I should be on "([^\"]*)"$/ do |url|
  URI.parse(current_url) == url
end