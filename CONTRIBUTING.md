<!--
 Copyright (c) 2020 DistAlchemist
 
 This software is released under the MIT License.
 https://opensource.org/licenses/MIT
-->

## Contributing

* First clone the repo:

```shell
cd $mg 
git clone https://github.com/DistAlchemist/Mongongo.git
# or git clone git@github.com:DistAlchemist/Mongongo.git
```

or sync with the remote:

```shell
git fetch origin
git checkout master
git rebase origin/master
```

* Create a new branch `dev-featurename` 

```shell
git checkout -b dev-test
```

* After you have made some progess, first commit it locally:

```shell
git status 
# make sure to add unwanted files to .gitignore
git add . # add all change files 
git commit -m "rewrite sql parser" # commit locally
```

* You may commit many times locally. Once you feel good about your branch, push it to remote.

```shell
git push origin dev-test # push local branch to origin with branch name `dev-test`
```

* Then view [https://github.com/DistAlchemist/Mongongo](https://github.com/DistAlchemist/Mongongo), Click the Compare & Pull Request button next to your `dev-test` branch.
