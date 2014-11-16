readtime
========

readtime is an trivially simple command line utility to give you an estimate of
how long it will take you to read a text-based document.

Usage
-----
I often use `wc` to give me an idea of how large a document is before reading
it, but the numbers don't really mean that much to me intuitively. This gives
me a fuzzy reading estimate based on a standard 200 words-per-minute, so I know
what I'm getting into beforehand.

    $ readtime README.md
    2 min read

Or, you could maybe use this to figure out how much of a windbag you are and get
Medium style estimates for all your Jekyll blog posts:

    $ readtime $BLOG/_posts/*.markdown | sort -nr | head -n5
     17 min read	_posts/2003-02-01-a-conversation-with-paul-bausch.markdown
     10 min read	_posts/2012-11-09-the-year-in-side-projects.markdown
      6 min read	_posts/2014-10-02-rubygems-bundler-they-took-our-jobs.markdown
      5 min read	_posts/2011-03-14-on-leaving-flickr.markdown
      5 min read	_posts/2009-07-08-palm-pre-mojo-sdk-experiments.markdown

You can override the default WPM (words-per-minute) calculation with the `-r`
flag.  For example, if you are a really fast reader, and want to practice your
reading skills on the dictionary:

    $ readtime -r 450 /usr/share/dict/words
    525 min read

Installation
------------
For now this is just available as source code, and requires Go to compile.

    go install github.com/mroth/readtime

If there is actually demand, I will make precompiled binaries available, and
package for homebrew, etc.
