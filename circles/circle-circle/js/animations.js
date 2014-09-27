function addMargin(elem, marginToAdd) {
    var marginTop = elem.css('margin-top').substring(0, elem.css('margin-top').length-2);
    margin = parseInt(marginTop)+marginToAdd+'px';
    return margin;
}

var main = function() {
    var animationTime = 2;

    for (i = 0; i < 100; i++) {
        $(".circle-div").delay(i*animationTime).animate({
            'margin-top': addMargin($(".circle-div"), 3*i),
            'margin-left': addMargin($(".circle-div"), 1*i)
        }, animationTime);
    }
};

$(document).ready(main);