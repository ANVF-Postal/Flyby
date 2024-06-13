# Flyby
Flyby is a terminal game meant to recreate the "fun" of opening cases in CS:GO.

It uses no engine, stock GO libraries and ANSI escape sequences, so this will only run terminals that support them. Flyby was developed and tested on a Debian server, using Nano through PuTTy.


I'm not a programmer, and this project was mainly an effort to learn the GO programming language, so expect bugs, and don't go too hard on my code okay, I'm trying :)

# Gameplay
```
---FLYBY------V------------------------------------------------
        GREAT   DEBT    AWAITS
--------------^------------------------------------------------
             / \
```
At the beginning of every Flyby game, you're given $100 to play into the machine. Upon betting, a random sequence of numbers ranged 0-9 will _fly by_ the tick arrows to the left of the machine. The number that lands on the tick marks is what your input bet will be multiplied by.
```
---FLYBY------V------------------------------------------------
030026000030006070800000680002002077006400770620309000000002208
--------------^------------------------------------------------
             / \   256 / 256
Cash: $700 WON! +$600
```
Of course, zero values are most common, and larger numbers are rarer than smaller numbers. 

Despite how it may feel, the game is not _rigged_ in any way, it's purely random, based on odds. There is no code to make you lose at critical times, or land on a zero that's right next to a nine. During testing, I was surprised by how rigged _my own game felt._ Technically, you can rig the game by setting either ```rigged``` or ```winner``` to true in the code, recompile, and you'll either get a 9 every time, or a zero right next to a 9 :)

The value of 256 below the machine represents how many frames the machine has left to tick before landing on the final number. This was originally a debug measure, but I found it added an element of suspense where players would try to predict the number as it got close to landing.

However, every bet there is a chance the machine won't complete the full 256 frames, and it will jam randomly on a random frame before 256. This does not help or hurt your odds past the first bet, because the numbers are still random either way. Jamming just adds to the fun.

Can you WIN Flyby?

Yes, the offical way to WIN is to BET ZERO, causing the game to quit. After that your current CASH is your score, and the game will print out your run stats. HIGHSCORE means nothing if it's not the balance you left with.

# Features
-A 64 character long random number sequence, that determines your win

-Rolling animation during bet

-Unique colors for each multiplier

-Win streak counter

-Messages for breaking certain records

-Extensively commented code

-Anti-gambling sentiment

-Dumb achievements

-JAMMING

-And more!

# Pull requests
I will be accepting pull requests for new features, bugfixes and spelling mistakes. I'm open to anything really.

# License
I dunno man, I picked GNU General Public License 3.0 because it sounded cool!

# Review of GO?
GO is not a bad language. I found it to be easier to navigate than C++, and I can actually understand what's going on. It's also got good speed, memory safety through garbage collection and quick compile times. I'd recommend it as a nice and versatile first programming language. The GOpher is also a hilarious mascot.
