# unixday [![Build Status](https://travis-ci.com/xyproto/unixday.svg?branch=master)](https://travis-ci.com/xyproto/unixday)

Utility for showing todays date as a unix day, or convert from a unix day to a date.

Just like regular UNIX time (counted in seconds since 1970), Unix Days does not take leap seconds into account. Be aware that Unix Days may therefore not match regular days exactly.

# examples

~~~bash
$ unixday
16457
~~~

~~~bash
$ unixday -date 16457
22.01.2015
~~~

