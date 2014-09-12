var main = function() {
    var timeout = 200;

    for (k = 0; k < 10; k++) {
        (function(i) {setTimeout(function() {
            $(".circle-div").each(function(j) {
                var self = $(this);
                setTimeout(function() {moveUp(self, 0, i%2!=0);}, timeout*j);
            });
        }, $(".circle-div").length*timeout*i); } (k));
    }

    // $(".circle-div").each(function(j) {
    //     var self = $(this);
    //     setTimeout(function() {moveUp(self, 0, false);}, timeout*j);
    // });

    // setTimeout(function() {
    //     $(".circle-div").each(function(j) {
    //         var self = $(this);
    //         setTimeout(function() {moveUp(self, 0, true);}, timeout*j);
    //     });
    // }, $(".circle-div").length*timeout);

    function moveUp(elem, i, up) {
        var marginTop = elem.css('margin-top').substring(0, elem.css('margin-top').length-2);

        if (up) {
            elem.css('margin-top', parseInt(marginTop)-20+'px');
        } else {
            elem.css('margin-top', parseInt(marginTop)+20+'px');
        }
        
        i++;

        if (i < 15) {
            setTimeout(function() {moveUp(elem, i, up);}, 50)
        }
    }
};

$(document).ready(main);