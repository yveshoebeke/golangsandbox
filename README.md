# golangsandbox


## Terminal oriented exercise

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


This should run right out of the box with the only requirement that mongod is installed and running.


Note: 
* _non-standard folders config/ and data/ contain the config.json and navplan.json files. Maybe another thing that could be improved in the real world scenario._
* _also: in a "real application" the db session would be made persistent and be more conform to a model, for example like here :_ https://github.com/thylong/mongo_mock_go_example


(With this Abe will never lose his way)


## HTTP oriented excercise

Data manipulation by browser. Yes, took some portions of the golang example and put some of my own twists on it.

bin/testdata will invoke a localhost. Default port is 8080, but can be set by adding a valid port number after the command.

bin/data: contains the data (.txt) files.
bin/tmpl: html templating.
bin/css: local style theme.

These locations can be improved on if this was to be a real-world setting.


## HTTP wrapper around the first excersise

(coming soon)
