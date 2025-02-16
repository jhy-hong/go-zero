package bug

const issueTemplate=`
<!-- Please answer these questions before submitting your issue. Thanks! -->

### What category of issue (<code>goctl</code> or <code>sdk</code>)?

### What type of issue (<code>feature</code>|<code>bug</code>|<code>suggestion</code>)?

### What version of Goctl are you using (<code>goctl --version</code>)?

<pre>
$ goctl --version
%s
</pre>

### Does this issue reproduce with the latest release?


### What operating system and processor architecture are you using ?
<pre>
%s
</pre>

### What did you do?

<!--
If possible, provide a recipe for reproducing the error.
A complete runnable program is good.
A link on play.golang.org is best.
-->



### What did you expect to see?



### What did you see instead?


`
