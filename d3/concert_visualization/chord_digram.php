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
<script>

var musiciansList = [
  {name: 'OutKast', venues: ['Coachella Day 1']},
  {name: 'Muse', venues: ['Coachella Day 2']},
  {name: 'Beck', venues: ['Coachella Day 3']},
  {name: 'Skrillex', venues: ['Coachella Day 2']},
  {name: 'Lorde', venues: ['Coachella Day 2']},
  {name: 'The Knife', venues: ['Coachella Day 1']},
  {name: 'The Replacements', venues: ['Coachella Day 1']},
  {name: 'Broken Bells', venues: ['Coachella Day 1']},
  {name: 'Zedd', venues: ['Coachella Day 1']},
  {name: 'Girl Talk', venues: ['Coachella Day 1']},
  {name: 'Ellie Goulding', venues: ['Coachella Day 1']},
  {name: 'Chromeo', venues: ['Coachella Day 1']},
  {name: 'HAIM', venues: ['Coachella Day 1']},
  {name: 'Neko Case', venues: ['Coachella Day 1']},
  {name: 'AFI', venues: ['Coachella Day 1']},
  {name: 'Martin Garrix', venues: ['Coachella Day 1']},
  {name: 'Bonobo', venues: ['Coachella Day 1']},
  {name: 'Bryan Ferry', venues: ['Coachella Day 1']},
  {name: 'The Glitch Mob', venues: ['Coachella Day 1']},
  {name: 'The Afghan Whigs', venues: ['Coachella Day 1']},
  {name: 'Queens of the Stone Age', venues: ['Coachella Day 2']},
  {name: 'Pharrell Williams', venues: ['Coachella Day 2']},
  {name: 'Foster the People', venues: ['Coachella Day 2']},
  {name: 'Pet Shop Boys', venues: ['Coachella Day 2']},
  {name: 'MGMT', venues: ['Coachella Day 2']},
  {name: 'Empire of the Sun', venues: ['Coachella Day 2']},
  {name: 'Fatboy Slim', venues: ['Coachella Day 2']},
  {name: 'Nas', venues: ['Coachella Day 2']},
  {name: 'Kid Cudi', venues: ['Coachella Day 2']},
  {name: 'The Head and the Heart', venues: ['Coachella Day 2']},
  {name: 'Sleigh Bells', venues: ['Coachella Day 2']},
  {name: 'Arcade Fire', venues: ['Coachella Day 3']},
  {name: 'Calvin Harris', venues: ['Coachella Day 3']},
  {name: 'Neutral Milk Hotel', venues: ['Coachella Day 3']},
  {name: 'Disclosure', venues: ['Coachella Day 3']},
  {name: 'Lana Del Ray', venues: ['Coachella Day 3']},
  {name: 'Motorhead', venues: ['Coachella Day 3']},
  {name: 'Alesso', venues: ['Coachella Day 3']},
  {name: 'Duck Sauce', venues: ['Coachella Day 3']},
  {name: 'Little Dragon', venues: ['Coachella Day 3']},
  {name: 'Flosstradmus', venues: ['Coachella Day 3']},
  {name: 'The Toy Dolls', venues: ['Coachella Day 3']},
  {name: 'Adventure Club', venues: ['Coachella Day 3']},
];

var imports = convertToJSONFlare(musiciansList);

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