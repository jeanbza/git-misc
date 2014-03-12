<!DOCTYPE html>
<meta charset="utf-8">
<style>

body {
  font: 10px sans-serif;
}

.chord {
  fill-opacity: .67;
}

</style>
<body>
<script src="http://d3js.org/d3.v3.min.js"></script>
<script src="schedule.js"></script>
<script>

var joinedMusicians = joinDuplicates(musiciansList);
var imports = convertToJSONFlare(joinedMusicians);

function joinDuplicates(arr) {
  var joinedHash = {};
  var joinedArr = [];

  arr.forEach(function(musician) {
    if (joinedHash[musician.name]) {
      musician.venues.forEach(function(venue) {
        // Check if the musician already has the venue listed (duplicate entry)
        if (joinedHash[musician.name].indexOf(venue) == -1) {
          joinedHash[musician.name].push(venue);
        }
      });
    } else {
      joinedHash[musician.name] = musician.venues;
    }
  });

  for (var musician in joinedHash) {
    var venues = joinedHash[musician];
    joinedArr.push({name: musician, venues: venues});
  }

  return joinedArr;
}

function convertToJSONFlare(arr) {
  var venues = [];
  var musicians = [];
  var temporaryVenues = {};

  arr.forEach(function(band) {
    var formattedVenues = [];

    band.venues.forEach(function(venue) {
      formattedVenues.push("flare."+venue+".a");

      if (temporaryVenues[venue]) {
        temporaryVenues[venue].push(band.name);
      } else {
        temporaryVenues[venue] = [band.name];
      }
    });

    musicians.push({
      name: "flare."+band.name+".a",
      size: 2,
      imports: formattedVenues
    })
  });

  for (var venue in temporaryVenues) {
    var venueMusicians = temporaryVenues[venue];
    var formattedMusicians = [];

    venueMusicians.forEach(function(musician) {
      formattedMusicians.push("flare."+musician+".a");
    });

    venues.push({
      name: "flare."+venue+".a",
      size: 2,
      imports: formattedMusicians
    });
  }

  return venues.concat(musicians);
}

var outerRadius = 760 / 2,
    innerRadius = outerRadius - 130;

var fill = d3.scale.category20c();

var chord = d3.layout.chord()
    .padding(.04)
    .sortSubgroups(d3.descending)
    .sortChords(d3.descending);

var arc = d3.svg.arc()
    .innerRadius(innerRadius)
    .outerRadius(innerRadius + 20);

var svg = d3.select("body").append("svg")
    .attr("width", outerRadius * 2)
    .attr("height", outerRadius * 2)
  .append("g")
    .attr("transform", "translate(" + outerRadius + "," + outerRadius + ")");

var indexByName = d3.map(),
    nameByIndex = d3.map(),
    matrix = [],
    n = 0;

// Returns the Flare package name for the given class name.
function name(name) {
  return name.substring(0, name.lastIndexOf(".")).substring(6);
}

// Compute a unique index for each package name.
imports.forEach(function(d) {
  if (!indexByName.has(d = name(d.name))) {
    nameByIndex.set(n, d);
    indexByName.set(d, n++);
  }
});

// Construct a square matrix counting package imports.
imports.forEach(function(d) {
  var source = indexByName.get(name(d.name)),
      row = matrix[source];
  if (!row) {
   row = matrix[source] = [];
   for (var i = -1; ++i < n;) row[i] = 0;
  }
  d.imports.forEach(function(d) { row[indexByName.get(name(d))]++; });
});

chord.matrix(matrix);

var g = svg.selectAll(".group")
    .data(chord.groups)
  .enter().append("g")
    .attr("class", "group");

g.append("path")
    .style("fill", function(d) { return fill(d.index); })
    .style("stroke", function(d) { return fill(d.index); })
    .attr("d", arc);

g.append("text")
    .each(function(d) { d.angle = (d.startAngle + d.endAngle) / 2; })
    .attr("dy", ".35em")
    .attr("transform", function(d) {
      return "rotate(" + (d.angle * 180 / Math.PI - 90) + ")"
          + "translate(" + (innerRadius + 26) + ")"
          + (d.angle > Math.PI ? "rotate(180)" : "");
    })
    .style("text-anchor", function(d) { return d.angle > Math.PI ? "end" : null; })
    .text(function(d) { return nameByIndex.get(d.index); });

svg.selectAll(".chord")
    .data(chord.chords)
  .enter().append("path")
    .attr("class", "chord")
    .style("stroke", function(d) { return d3.rgb(fill(d.source.index)).darker(); })
    .style("fill", function(d) { return fill(d.source.index); })
    .attr("d", d3.svg.chord().radius(innerRadius));

d3.select(self.frameElement).style("height", outerRadius * 2 + "px");

</script>