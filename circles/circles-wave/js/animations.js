var main = function() {
    var timeout = 200;
    var allCirclesFinished = [];

    moveCircles(false);
    
    function moveCircles(up) {
        $(".circle-div").each(function(j) {
            var self = $(this);
            var promiseToFinish = new $.Deferred();
            allCirclesFinished.push(promiseToFinish);
            setTimeout(function() {moveCircle(self, 0, up, promiseToFinish);}, timeout*j);
        });

        $.when.apply($, allCirclesFinished).then(function() {
            allCirclesFinished = [];
            moveCircles(!up);
        });
    }

    function moveCircle(elem, i, up, promise) {
        var marginTop = elem.css('margin-top').substring(0, elem.css('margin-top').length-2);

        if (up) {
            elem.css('margin-top', parseInt(marginTop)-20+'px');
        } else {
            elem.css('margin-top', parseInt(marginTop)+20+'px');
        }
        
        i++;

        if (i < 15) {
            setTimeout(function() {moveCircle(elem, i, up, promise);}, 50)
        } else {
            promise.resolve();
        }
    }
};

$(document).ready(main);