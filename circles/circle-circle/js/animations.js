function adjustPosition(elem, degree) {
    var radius = 100;
    var marginX = radius*Math.cos(degree);
    var marginY = radius*Math.sin(degree);
    
    $(elem).css("margin-top", marginY);
    $(elem).css("margin-left", marginX);
}

function circleOnce() {
    var fullRotation = 314*2;
    var animationTime = 2;

    for (i = 0; i < fullRotation; i++) {
        (function(i) {
            setTimeout((function() {
                adjustPosition($(".circle-div"), i/100);
            }), i*animationTime);
        })(i);
    }

    setTimeout(function() {
        circleOnce();
    }, animationTime*fullRotation);
}

// (x0,y0) and whose radius is r is (x0 + r cos theta, y0 + r sin theta)
var main = function() {
    circleOnce();
};

$(document).ready(main);