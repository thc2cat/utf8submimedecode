Simple mime decode filter using mime.Decoder go libs
====================================================

Because reading utf8, and other encoded subjects in mail headers is no intuitive,
this filter translate mime encoded text.

utf8submimedecode was needed for another project [subayes](https://git.dsi.uvsq.fr/thiecail/subayes),
a mail subject ham/spam discriminator.

Bonus ! golang mime decoder also decode iso-8859-1

Exemple
-------

```shell
$ echo "Re: =?UTF-8?B?w4l0YXQgQ0VU?=" | utf8submimedecode
Re: État CET

$ echo '=?iso-8859-1?q?this=20is=20some=20text?=' | utf8submimedecode
this is some text

$ echo '=?ISO-8859-1?B?SWYgeW91IGNhbiByZWFkIHRoaXMgeW8=?=' | utf8submimedecode                       >
If you can read this yo

```
