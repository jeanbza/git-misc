var main = function() {
  // fade and animate each circle up together
  //$(".circle-div").animate({opacity:1,'margin-top':0}, 600);

  // fade circle in one at a time
  // all animate up together
  $(".circle-div").each(function(i) { 
    $(this).delay(i*200).animate({opacity:1,'margin-top':0});
  })

  // $(".mint-circle-div").animate({opacity:1,'margin-top':0}, 600);
  // $(".green-circle-div").delay(100).animate({opacity:1,'margin-top':0}, 600);
  // $(".blue-circle-div").delay(200).animate({opacity:1,'margin-top':0}, 600);
  // $(".purple-circle-div").delay(300).animate({opacity:1,'margin-top':0}, 600);
  // $(".navy-circle-div").delay(400).animate({opacity:1,'margin-top':0}, 600);
};

$(document).ready(main);