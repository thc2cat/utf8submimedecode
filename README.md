Simple utf8 decode filter using go libs
========================================

Because reading utf8 encoded subjects in mail headers is no intuitive,
this filter translate utf8 encoded text.

utf8submimedecode was needed for another project [subayes](https://git.dsi.uvsq.fr/thiecail/subayes),
a mail subject ham/spam discriminator.

Exemple
-------

```shell
$ echo "Re: =?UTF-8?B?w4l0YXQgQ0VU?=" | utf8submimedecode
Re: État CET
```