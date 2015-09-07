# Feature: pwgen behavior

The most popular request so far is to behave like pwgen. Without
further reference this seems to refer to the pwgen command line
utility found in Ubuntu, which aims to generate pronounceable
password.

To imitate its behavior, we have to do two things: Provide the
same CLI, and provide the same functionality.


## CLI Changes

The CLI of pwgen uses the usual getopt style found everywhere
(except in golang). So the first order of the day is to somehow
understand this style, probably whip up a kind of getopt (because
I don't believe in depending on unspecified, unversioned external
packages). Once this is done, we have to support the following
options:

    mpw [ OPTION ] [ pwlength ] [ numpw ]
    
    -0, --no-numerals     - no numbers
    -1                    - one password per line
    -A, --no-capitalize   - no need for capital letters?
    -a, --alt-phonics     - simply ignore
    -B, --ambigous        - no letters which might be confused
    -c, --capitalize      - at least one capital letter
    -C                    - print output in columns
    -N, --num-password=NUM    - generate NUM passwords
    -n, --numerals        - at least one number
    -H, --sha1=PATH[#SEED]    - seed with hash of file PATH[#SEED]
    -h, --help            - help
    -s, --secure          - completely random
    -v, --no-vowels       - no vowels (or numbers mistaken for vowels)
    -y, --symbols         - at least one special character

This is extracted from the man page of pwgen by
Theodore Ts'o (<tytso@alum.mit.edu>.)

Note that we will probably not be "compatible" with pwgen, as our
generator is likely to be different from the pwgen generator, so
using the same seed file (the -H option) will most probably not
yield the same passwords.


## Functionality Changes

Some of the functionality here can already be achived by selecting
the proper symbol sets, other functionality needs real changes. From
my original Python implemented password generator stems the idea of
scramblers: filters which mutate a given password, possibly in a loop
until specific criterias are meet.

Another field are pronounceable passwords. My original intend was to
use a symbol set of syllables, which can be concatenated into a
pronounceable "word", and this still seems to be a good way to
implement pronounceable passwords. Since putting lists of syllables
into code seems to be strange to be, we will need a way to read
a configuration file to construct a symbol set of all possible
syllables.

Finally, the column output: This might be hard or easy, depending on
how much of the original look and feel of pwgen we need to mimick.
For starters, it might be ok to just assume the 80 columns terminal.

