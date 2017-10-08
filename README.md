# golangsandbox

The code contained here is meant to be an educational project of the Go language.

It will read its configuration and navigation waypoint data and calculate distance and bearing between the point defining a segment (leg). It will record this usage in mongo.


This repository contains ~~2~~ ~~3~~ 5 packages:

* dandb - Calculates distance and bearing between 2 geo locations.
* readnav - Reads a json file with pertinent waypoint data. ~~File locations are handled by some hand written code to read a config file, but could be improved on by using something like Viper or gonfig. We'll cross that bridge when needed.~~ Ostensibly this data could be read from some data collection.
* myconfig - Package to read the (albeit minimal here) configuration file. I understand there are more robust packages out there (Viper and Gonfig come to mind).
* logusage - Records user (login) name and time results when results were displayed, in a mongoDB collection.
* gopkg.in - Contains mongo driver and Bson perticulars.


And one main fuction that ties it all together:

* makenavplan - Just puts it all together, if you will; and outputs to console.


Note: _non-standard folders config/ and data/ contain the config.json and navplan.json files. Maybe another thing that could be improved in the real world scenario._


(With this Abe will never lose his way)
