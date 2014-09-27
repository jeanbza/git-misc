var main = function() {
    var timeout = 200;
    var allCirclesFinished = [];

    moveCircles(false);
    
    function moveCircles(negative) {
        $(".circle-div").each(function(j) {
            var self = $(this);
            var promiseToFinish = new $.Deferred();
            allCirclesFinished.push(promiseToFinish);
            setTimeout(function() {moveCircle(self, 0, negative, false, promiseToFinish);}, timeout*j);
        });

        $.when.apply($, allCirclesFinished).then(function() {
            allCirclesFinished = [];
            moveCircles(!negative);
        });
    }

    function moveCircle(elem, i, negative, horizontally, promise) {
        var marginTop = elem.css('margin-top').substring(0, elem.css('margin-top').length-2);

        if (horizontally) {
            if (negative) {
                elem.css('margin-top', parseInt(marginTop)-20+'px');
            } else {
                elem.css('margin-top', parseInt(marginTop)+20+'px');
            }
        } else {
            if (negative) {
                elem[0].offsetLeft = elem[0].offsetLeft-20;
            } else {
                elem[0].offsetLeft = elem[0].offsetLeft+20;
            }
        }
        
        i++;

        if (i < 15) {
            setTimeout(function() {moveCircle(elem, i, negative, horizontally, promise);}, 50)
        } else {
            promise.resolve();
        }
    }
};

$(document).ready(main);