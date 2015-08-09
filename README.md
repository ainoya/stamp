Stamp
=======

Sign git commit-id on ios app icon when build phase.

![](https://dl.dropboxusercontent.com/u/10177896/4062bb26d867f3b6488bef4abf8dd009.png)



Setup
------

First, to use stamp, you have to install `imagemagick` and `ghostscript` command to your mac:

```
brew install imagemagick ghostscript
```

Then, download `stamp` binary:

```
wget https://github.com/ainoya/stamp/releases/download/v0.1.0/stamp -O /usr/local/bin/stamp
```


How to use
-----------

```
USAGE:
   Stamp [global options] command [command options] [arguments...]

GLOBAL OPTIONS:
   --in                 File path of source icon
   --out                File path of output image
```

For example:

```
stamp --in YourAppRepository/icon-60@2x-origin.png --out assets/icon-60@2x.png
```

If you want to add git commit-id to your app icon with `stamp`,
add below script run script section in "Build phase" of Xcode:

```
export PATH=$PATH:/usr/local/bin

for orig in $(find ${SRCROOT} -name '*-origin.png'); do
    OUTPUT_PATH=$(echo ${orig} | sed "s/-origin//")
stamp --in ${orig} --out ${OUTPUT_PATH}
done
```

In this case, you have to locate your original app icon as named `*-origin.png`. this script strip `-origin` prefix and use it as generated icon name.

for Example, if you name original icon as `icon-60@2x-origin.png`, Generated file will be named `icon-60@2x.png`, then you need to set `icon-60@2x.png` instead of `icon-60@2x-origin.png` in your App icon assets before.

Reference
-----------

- [Overlaying application version on top of your icon - Krzysztof Zabłocki](http://merowing.info/2013/03/overlaying-application-version-on-top-of-your-icon/#.Vbxr_ZOqpBd)
