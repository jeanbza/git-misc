function adjustPosition(elem, degree) {
    var radius = 100;
    var marginX = radius*Math.cos(degree);
    var marginY = radius*Math.sin(degree);

    console.dir(" ");
    console.dir(degree);
    console.dir(marginY);
    
    $(elem).css("margin-top", marginY);
    $(elem).css("margin-left", marginX);
}

// (x0,y0) and whose radius is r is (x0 + r cos theta, y0 + r sin theta)
var main = function() {
    var animationTime = 2;

    for (i = 0; i < 100; i++) {
        (function(i) {
            setTimeout((function() {
                adjustPosition($(".circle-div"), i/10);
            }), i*20);
        })(i);
    }
};

$(document).ready(main);