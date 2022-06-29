Bugfix: Define errors in YAML format

We had been parsing JSON files only which lead to issues if somebody tries to
override the error codes with a YAML file. We had to replace the document parser
to handle that properly.

https://github.com/webhippie/errors/issues/15
