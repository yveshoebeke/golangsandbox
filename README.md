# golangsandbox

This reposatory contains ~~2~~ 3 packages:

* dandb - Calculates distance and bearing between 2 geo locations.
* readnav - Reads a json file with pertinent waypoint data. ~~File locations are handled by some hand written code to read a config file, but could be improved on by using something like Viper or gonfig. We'll cross that bridge when needed.~~
* myconfig - Package to read the (albeit minimal here) configuration file. I understand there are more robust processes out there (Viper, Gonfig).

And one main fuction that ties it all together:

* makenavplan - Just puts the 2 and 2 together if you will and outputs to console.

Note: _non-standard folders config/ and data/ contain the config.json and navplan.json files. Maybe another thing that could be improved in the real world._
_I was looking for an excuse to use interface{} but couldn't readily identify a real purpose for one, probably somewhere when myconfig would get more mature_

(With this Abe will never lose his way)
