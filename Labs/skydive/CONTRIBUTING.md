Contributing to Skydive
========================

Your contributions are more than welcome. Please read the following notes to
know about the process.

Making Changes:
---------------

Before pushing your changes, please make sure the tests and the linters pass:

* make fmt
* make test
* make functional

_Please note, make functional will create local network resources
(bridges, namespaces, ...)_

Once ready push your changes to your Skydive fork our CI will check your
pull request.

We're much more likely to approve your changes if you:

* Add tests for new functionality.
* Write a good commit message, please see how to write below.
* Maintain backward compatibility.

How To Submit Code
------------------

We use github.com's Pull Request feature to receive code contributions from
external contributors. See
https://help.github.com/articles/creating-a-pull-request/ for details on
how to create a request.

Commit message format
---------------------

The subject line of your commit should be in the following format:

area: summary

area :
Indicates the area of the Skydive to which the change applies :

* ui
* flow
* capture
* api
* graph
* cmd
* netlink
* ovsdb
* etc.

Feature request or bug report:
------------------------------

Be sure to search for existing bugs before you create another one.
Remember that contributions are always welcome!

https://github.com/skydive-project/skydive/issues

Contact
-------

Contact the project developers via the project's "dev" list or IRC.

* IRC: #skydive-project on irc.freenode.net
* Mailing list: https://www.redhat.com/mailman/listinfo/skydive-dev
