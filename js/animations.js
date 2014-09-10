var main = function() {
    $(".circle-div").each(function(j) {
        var self = $(this);
        setTimeout(function() {moveUp(self, 0);}, 200*j);
    });

    function moveUp(elem, i) {
        var marginTop = elem.css('margin-top').substring(0, elem.css('margin-top').length-2);

        elem.css('margin-top', marginTop-20+'px');
        i++;

        if (i < 15) {
            setTimeout(function() {moveUp(elem, i);}, 50)
        }
    }
};

$(document).ready(main);