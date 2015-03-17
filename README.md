# badgerd

Go daemon that serves badges for various jenkins project statuses.

*Warning - this is a learning project and as such is still subject to massive breaking changes.*

## Usage

When properly configured you can use badgerd to serve status badges for your jenkins projects.

http://jenkins-server.example.com/badge/Jenkins+Project+Name

This will serve one of 4 badges - build passing, build failing, build running, or build error.

You can also serve coverage badges for projects using rcov with the [RubyMetrics plugin](https://wiki.jenkins-ci.org/display/JENKINS/RubyMetrics+plugin).

http://jenkins-server.example.com/badge/Jenkins+Project+Name/rcov

This call provides badges for coverage at 0-100% or error if the coverage stat can't be found.


## Installation

* copy the config.json.example to config.json and update the values to suit your needs
* build badgerd.go
* copy badgerd, config.json and the badges/ folder to your server
* configure your front end webserver to forward the location /badge/ to the port you've configured (:8081 by default)
* ./badgerd

## Todo

* Make badgerd run as a daemon
* Make init scripts
* Improve logging
* Improve configuration (optional params, commandline arguments)

## Thanks

The badge files come from the amazing [shields.io](http://shields.io/) site. If you need a badge they've got you covered.
