# Notes on pip

## Dependency Management

Most Python programs rely on third-party libraries and frameworks. These third-party installations are dependencies. The third-party packages can have their own dependencies too, indirect or transitive dependencies. And of course those dependencies of dependencies can also depend on other things, and so on and so on.

You do not want to handle this chain of dependencies by hand. It is time consuming and error prone. If you do it by hand, you are likely to mess it up yourself, and you make it harder for other people to use your software. Instead, you should use another piece of software, a package manager, which can handle much of this automatically or at least programatically.

The recommended package manager for Python is pip. You should have a version of pip installed as part of your Python installation. You can access pip most easily via a command-line program, also called “pip.”

## Python Software Repositories

The largest software repositories for Python is PyPI (aka, “the cheese shop”).The website is https://pypi.python.org. Anyone can sign up there, and there is no QA or quality control of any kind. So you should be thoughtful before you install things from PyPI.

Warehouse is going to replace PyPI eventually, but it is currently (?) in beta. The change won’t affect users much. It will simply be a different back end. Nope, scratch this. The Warehouse transition has already happened.

## Installing Packages with pip

Basically:

+ `pip install [package name]`
+ `pip show [package name]` shows metadata about a package
+ `pip install` will install the newest version by default, but you can also specify a version `pip install [package name] == [version number]`.
+ Version information can get more complicated. For example, `name>=3,<4`.
+ You can also use `~=` to get compatible versions. This seems to mean basically minor upgrades. So `~=2.1.3` would include `2.1.4, 2.1.5` and so on.

Be smart about global installs. They make sense sometimes, especially when you are installing something with a command-line tool that you want available for the entire Python installation. But usually it is better to install dependencies in a virtual environment.

You can also use pip to install from a specific git repo and branch. You can get very granular: you can specify a branch, commit, or tag.

## Upgrading Packages with pip

+ `pip list --outdated` will tell you what packages are out of date.
+ `pip install --upgrade|-U [name]` will update a package. Note the shortcut of `-U` instead of `--upgrade`.

## Uninstalling Packages with pip

+ `pip uninstall [name]` removes a package from your system.

But what if I want to remove a package and its dependencies? Does `pip uninstall` clean up and remove all of a packages dependencies? No, it does not. If you want to keep a Python installation minimal and clean, you should prefer virtual environments.
