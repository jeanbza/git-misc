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
  {name: 'OutKast', venues: ['Coachella']},
  {name: 'Muse', venues: ['Coachella']},
  {name: 'Beck', venues: ['Coachella']},
  {name: 'Skrillex', venues: ['Coachella', 'Bonnaroo']},
  {name: 'Lorde', venues: ['Coachella']},
  {name: 'The Knife', venues: ['Coachella']},
  {name: 'The Replacements', venues: ['Coachella']},
  {name: 'Broken Bells', venues: ['Coachella']},
  {name: 'Zedd', venues: ['Coachella', 'Bonnaroo', 'Ultra']},
  {name: 'Girl Talk', venues: ['Coachella']},
  {name: 'Ellie Goulding', venues: ['Coachella']},
  {name: 'Chromeo', venues: ['Coachella']},
  {name: 'HAIM', venues: ['Coachella']},
  {name: 'Neko Case', venues: ['Coachella']},
  {name: 'AFI', venues: ['Coachella']},
  {name: 'Martin Garrix', venues: ['Coachella']},
  {name: 'Bonobo', venues: ['Coachella']},
  {name: 'Bryan Ferry', venues: ['Coachella']},
  {name: 'The Glitch Mob', venues: ['Coachella', 'Bonnaroo', 'Ultra']},
  {name: 'The Afghan Whigs', venues: ['Coachella']},
  {name: 'Queens of the Stone Age', venues: ['Coachella']},
  {name: 'Pharrell Williams', venues: ['Coachella']},
  {name: 'Foster the People', venues: ['Coachella']},
  {name: 'Pet Shop Boys', venues: ['Coachella']},
  {name: 'MGMT', venues: ['Coachella', 'Ultra']},
  {name: 'Empire of the Sun', venues: ['Coachella', 'Ultra']},
  {name: 'Fatboy Slim', venues: ['Coachella']},
  {name: 'Nas', venues: ['Coachella']},
  {name: 'Kid Cudi', venues: ['Coachella']},
  {name: 'The Head and the Heart', venues: ['Coachella']},
  {name: 'Sleigh Bells', venues: ['Coachella']},
  {name: 'Arcade Fire', venues: ['Coachella']},
  {name: 'Calvin Harris', venues: ['Coachella']},
  {name: 'Neutral Milk Hotel', venues: ['Coachella', 'Bonnaroo']},
  {name: 'Disclosure', venues: ['Coachella', 'Bonnaroo']},
  {name: 'Lana Del Ray', venues: ['Coachella']},
  {name: 'Motorhead', venues: ['Coachella']},
  {name: 'Alesso', venues: ['Coachella']},
  {name: 'Duck Sauce', venues: ['Coachella']},
  {name: 'Little Dragon', venues: ['Coachella']},
  {name: 'Flosstradmus', venues: ['Coachella']},
  {name: 'The Toy Dolls', venues: ['Coachella']},
  {name: 'Adventure Club', venues: ['Coachella']},
  {name: "Elton John", venues: ['Bonnaroo']},
  {name: "Kanye West", venues: ['Bonnaroo']},
  {name: "Jack White", venues: ['Bonnaroo']},
  {name: "Lionel Richie", venues: ['Bonnaroo']},
  {name: "Vampire Weekend", venues: ['Bonnaroo']},
  {name: "The Avett Brothers", venues: ['Bonnaroo']},
  {name: "Arctic Monkeys", venues: ['Bonnaroo']},
  {name: "Frank Ocean", venues: ['Bonnaroo']},
  {name: "The Flaming Lips", venues: ['Bonnaroo']},
  {name: "Nick Cave & The Bad Seeds", venues: ['Bonnaroo']},
  {name: "Kaskade", venues: ['Bonnaroo', 'Ultra']},
  {name: "Wiz Khalifa", venues: ['Bonnaroo']},
  {name: "Damon Albarn", venues: ['Bonnaroo']},
  {name: "Cut Copy The Head and the Heart", venues: ['Bonnaroo']},
  {name: "Ms. Lauryn Hill", venues: ['Bonnaroo']},
  {name: "Funkiest Dancer", venues: ['Bonnaroo']},
  {name: "Chromeo", venues: ['Bonnaroo']},
  {name: "Broken Bells Tedeschi Trucks Band", venues: ['Bonnaroo']},
  {name: "James Blake", venues: ['Bonnaroo']},
  {name: "Bobby Womack", venues: ['Bonnaroo']},
  {name: "Umphrey's McGee", venues: ['Bonnaroo']},
  {name: "Ice Cube", venues: ['Bonnaroo']},
  {name: "Ben Howard Slightly Stoopid", venues: ['Bonnaroo']},
  {name: "Fitz & The Tantrums", venues: ['Bonnaroo']},
  {name: "Cake", venues: ['Bonnaroo']},
  {name: "Janelle Monáe", venues: ['Bonnaroo']},
  {name: "Grouplove", venues: ['Bonnaroo']},
  {name: "Amos Lee", venues: ['Bonnaroo']},
  {name: "CHVRCHES Cage the Elephant", venues: ['Bonnaroo']},
  {name: "Die Antwoord", venues: ['Bonnaroo']},
  {name: "Drive-By Truckers", venues: ['Bonnaroo']},
  {name: "Andrew Bird & the Hands of Glory", venues: ['Bonnaroo']},
  {name: "Mastodon Capital Cities", venues: ['Bonnaroo']},
  {name: "Jake Bugg", venues: ['Bonnaroo']},
  {name: "Chance The Rapper", venues: ['Bonnaroo']},
  {name: "Dr. Dog", venues: ['Bonnaroo']},
  {name: "Yonder Mountain String Band John Butler Trio", venues: ['Bonnaroo']},
  {name: "Little Dragon", venues: ['Bonnaroo']},
  {name: "City and Colour", venues: ['Bonnaroo']},
  {name: "The Naked and Famous Taran Killam", venues: ['Bonnaroo']},
  {name: "Phosphorescent", venues: ['Bonnaroo']},
  {name: "Washed Out", venues: ['Bonnaroo']},
  {name: "Danny Brown", venues: ['Bonnaroo']},
  {name: "Warpaint", venues: ['Bonnaroo']},
  {name: "Sam Smith", venues: ['Bonnaroo']},
  {name: "A$AP Ferg", venues: ['Bonnaroo']},
  {name: "Darkside Seasick Steve", venues: ['Bonnaroo']},
  {name: "Shovels & Rope", venues: ['Bonnaroo']},
  {name: "Lucero", venues: ['Bonnaroo']},
  {name: "Real Estate", venues: ['Bonnaroo']},
  {name: "Carolina Chocolate Drops", venues: ['Bonnaroo']},
  {name: "The Wood Brothers", venues: ['Bonnaroo']},
  {name: "Meshuggah", venues: ['Bonnaroo']},
  {name: "Poliça", venues: ['Bonnaroo']},
  {name: "DakhaBrakha", venues: ['Bonnaroo']},
  {name: "Goat", venues: ['Bonnaroo']},
  {name: "ZZ Ward", venues: ['Bonnaroo']},
  {name: "Seun Kuti & Egypt 80", venues: ['Bonnaroo']},
  {name: "Blackberry Smoke MS MR", venues: ['Bonnaroo']},
  {name: "Hannibal Buress", venues: ['Bonnaroo']},
  {name: "First Aid Kit", venues: ['Bonnaroo']},
  {name: "Rudimental", venues: ['Bonnaroo']},
  {name: "A Tribe Called Red", venues: ['Bonnaroo']},
  {name: "Omar Souleyman The Bouncing Souls", venues: ['Bonnaroo']},
  {name: "Greensky Bluegrass", venues: ['Bonnaroo']},
  {name: "Ty Segall", venues: ['Bonnaroo']},
  {name: "Sarah Jarosz", venues: ['Bonnaroo']},
  {name: "Vintage Trouble", venues: ['Bonnaroo']},
  {name: "Okkervil River White Denim", venues: ['Bonnaroo']},
  {name: "Jonathan Wilson", venues: ['Bonnaroo']},
  {name: "J. Roddy Walston & the Business", venues: ['Bonnaroo']},
  {name: "Robert Delong", venues: ['Bonnaroo']},
  {name: "Cloud Nothings Typhoon", venues: ['Bonnaroo']},
  {name: "Thao & the Get Down Stay Down", venues: ['Bonnaroo']},
  {name: "Valerie June", venues: ['Bonnaroo']},
  {name: "King Khan & the Shrines", venues: ['Bonnaroo']},
  {name: "Cherub", venues: ['Bonnaroo']},
  {name: "BANKS Break Science", venues: ['Bonnaroo']},
  {name: "The Black Lillies", venues: ['Bonnaroo']},
  {name: "The Lone Bellow", venues: ['Bonnaroo']},
  {name: "Caveman", venues: ['Bonnaroo']},
  {name: "Big Sam's Funky Nation Jon Batiste and Stay Human", venues: ['Bonnaroo']},
  {name: "La Santa Cecilia", venues: ['Bonnaroo']},
  {name: "Classixx", venues: ['Bonnaroo']},
  {name: "Allah-Las", venues: ['Bonnaroo']},
  {name: "Cass McCombs", venues: ['Bonnaroo']},
  {name: "Vance Joy", venues: ['Bonnaroo']},
  {name: "Haerts Those Darlins", venues: ['Bonnaroo']},
  {name: "Deafheaven", venues: ['Bonnaroo']},
  {name: "Lake Street Dive", venues: ['Bonnaroo']},
  {name: "St. Paul & The Broken Bones", venues: ['Bonnaroo']},
  {name: "The Wild Feathers The Preatures", venues: ['Bonnaroo']},
  {name: "Blank Range", venues: ['Bonnaroo']},
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