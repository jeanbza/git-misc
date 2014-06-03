Feature: Clicking Links
  I want to make sure clicking links works. Note: these tests occur as
  a non-signed-in user, so if you want to see the same things, open up
  and incognito browser and go to github.com
  
  Scenario: Clicking Explore
    Given I have opened "https://github.com"
    When I click "Explore"
    Then I should be on "https://github.com/explore"

  Scenario: Clicking Features
    Given I have opened "https://github.com"
    When I click "Features"
    Then I should be on "https://github.com/features"

  Scenario: Clicking Enterprise
    Given I have opened "https://github.com"
    When I click "Enterprise"
    Then I should be on "https://enterprise.github.com/"

  Scenario: Clicking Blog
    Given I have opened "https://github.com"
    When I click "Blog"
    Then I should be on "https://github.com/blog"