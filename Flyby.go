package main

import ("fmt"; "time"; "math/rand"; "os")

var (
//ints
 lines int // Tracks how many terminal lines are written
 cash int // Starting money and balance (Set in main())
 wins int // How many wins (unused)
 losses int // How many losses (unused)
 bet int // User submitted bet amount
 bets int // How many bets the player made
 last int // Stores how many letters were written by progWrite
 score int // Most money the player had at peak
 frames int // How long the rolling animation wil be
 rate int // Animation roll delay in miliseconds
 arrSize = 64 // Number of cells in the rolling array
 winStreak int // Count how many wins in a row
 jam int
 achievements int
 perceptive int
//bools
 celebrateMil = true // Celebrate breaking over one million
 celebrate1k  = true // Celebrate breaking a thousand
 celebrate1hk = true // Celebrate breaking a hundred thousand
 ach69   = false // has the player gotten acheivements for 69, 404, etc
 ach404  = false
 ach420  = false
 ach808  = false
 ach1337 = false
 jamming = true //Enable or disable jamming
 jammed  = false // This is set if the machine jams
 rigged  = false // ALWAYS lose, right next to a  9
 winner  = false // ALWAYS roll 9
//strings
 tickerMode = "None" //Option for how the ticker behaves. Options are None, Perception and Ticker
//arrays
 arrSlice = make([]int, arrSize) // Where the rolling numbers are stored
//ANSI colors set to strings. Print any of these strings to color the terminal
    Reset   = "\033[0m" //sets color to default again.
    Red     = "\033[31m"
    Green   = "\033[32m"
    Yellow  = "\033[33m"
    Blue    = "\033[34m"
    Magenta = "\033[35m"
    Cyan    = "\033[36m"
    Gray    = "\033[90m"
    BRed    = "\033[91m"
    BGreen  = "\033[92m"
    BYellow = "\033[93m"
    BBlue   = "\033[94m"
    BMagenta= "\033[95m"
    BCyan   = "\033[96m"
    White   = "\033[97m"
    Black   = "\033[30m"

    colorMap = map[int]string{ // set the color of each number
    0: Black,
    1: White, //1 can never occur, but if it did
    2: White,
    3: Cyan,
    4: BBlue,
    5: BCyan,
    6: BMagenta,
    7: BGreen,
    8: BRed,
    9: BYellow,
    }
)

func clean(x int) { //where "X" is the number of lines you want to clean
   fmt.Print("\x1b[",x,"A\x1b[J")
   lines = lines - x
   if (lines < 0) {
   lines = 0
   }
}

func progWrite(s string, delayMS int) { //write the given string out one letter at a time
        last = 0 // reset count each time prog write is called
        for _, letter := range s {
                fmt.Printf("%c", letter)
                last++ //call "last" in progDel to erase these lines
                time.Sleep(time.Duration(delayMS) * time.Millisecond)
        }
}

func progDel(rm int, dly int) { //won't work across lines
        i := 0 // the := tells the compiler "figure this variable out"
        for i < rm {
        fmt.Printf("\b \b") // backspace. needs two to delete the pointer letter
        i++
        last--
        time.Sleep(time.Duration(dly) * time.Millisecond)
        }
}

func main() {
// intro movie
// maybe add a check for terminal size here?
        //begin startup checks
        startupChecks()
        //end startup checks
        cash = 100
        score = cash
        winStreak = 0
        //Reset = Green // weed mode
        color(Reset)
        rand.Seed(time.Now().UnixNano())
        progWrite("---FLYBY------V------------------------------------------------\n", 50)
        introText()
        progWrite("--------------^------------------------------------------------\n", 25)
        fmt.Print("          / \\", "\n") // two backslashes to show it's not an operator
        fmt.Print("Cash: ", cash)
        lines += 4
        funcBet()
        }

func funcBet() {
 fmt.Print("\n") // balance is printed in this new line
 fmt.Print("Place your bet: "); lines += 1
 fmt.Scan(&bet)
 if (bet > cash){
  clean(2)
  fmt.Print("Not enough cash. You have $", cash)
  funcBet()
 }
 switch bet {
  case 69:
   if (ach69 == false) {
    achWrite(BMagenta, "Hehehehehee...", 1000, 100)
    ach69 = true
    achievements++
    }
  case 404:
   if (ach404 == false) {
    achWrite(Cyan, "Bet not found...", 1000, 50)
    ach404 = true
    achievements++
    }
  case 420:
   if (ach420 == false) {
    achWrite(Green, "Green like that", 1000, 70)
    ach420 = true
    achievements++
    }
  case 808:
   if (ach808 == false) {
    achWrite(BYellow, "Drum Machine", 808, 80)
    ach808 = true
    achievements++
    }
  case 1337:
   if (ach1337 == false) {
    achWrite(BGreen, "Gamer!", 1500, 50)
    ach1337 = true
    achievements++
    }
 }
 if (bet == cash && cash >= 10000) { //Taunt player for betting everything
   msgAllIn()
 }
 if (bet <= 0 && bets == 0) { // congradulate dudes for not gambling
  color(BCyan)
  progWrite("You'll never lose", 50)
  wait(1000)
  progWrite(" if you don't play.", 50)
  wait(2000)
  color(Reset)
  quit()
 }
 if (bet <= 0) { // Quit sequence
  color(BGreen)
  progWrite("You decided to quit!", 50)
  wait(2000)
  fmt.Println(BYellow,"\n--------------------------------------", Reset)
  wait(1000)
  fmt.Println("Final balance ",BYellow, "$",cash, Reset)
  wait(1000)
  fmt.Println("Highest score ",BGreen,"$",score, Reset)
  wait(1000)
  if (achievements > 0) { // print achievement count.
   color(BMagenta)
   fmt.Print("Achievements    ")
   wait(700)
   fmt.Print(achievements)
   wait(1000)
   fmt.Print(Reset, "\n")
  }
  fmt.Println("Bet count      ",bets)
  wait(1000)
  fmt.Println("Win/loss       ",wins,"/",losses)
  wait(2000)
  color(BCyan)
  progWrite("Thank you for playing!",50)
  color(Reset)
  wait(1000)
  fmt.Print("\n")
  quit()
 }
 bets++
 clean(2) // wipe balance and be lines
 funcRoll(bet)
}

func funcRoll(bet int) {
 moveUp(3)
 rate = 10 // old value: 8
 frameMax := 256
 if (jamming == true) { // jamming function.
  rand.Seed(time.Now().UnixNano())
  r := rand.Intn(4) // r is just a new variable for this RNG
  if (r == 1) {
   jammed = true
   jam = rand.Intn(frameMax)
 } else {
  jam = frameMax //else, no jamming)
  }
 }
 var result int // set from 0 - 9 depending on number from RNG
 for frames := 0; frames < jam; frames++ { // set jam to frameMax if it breaks!
  rand.Seed(time.Now().UnixNano())
  rng := rand.Intn(100)
   if (rng <= 60) { result = 0; } // winning / odds table
   if (rng >= 61 && rng < 70) { result = 2; }
   if (rng >= 70 && rng < 75) { result = 3; }
   if (rng >= 75 && rng < 80) { result = 4; }
   if (rng >= 80 && rng < 85) { result = 5; }
   if (rng >= 85 && rng < 90) { result = 6; }
   if (rng >= 90 && rng < 95) { result = 7; }
   if (rng >= 95 && rng < 99) { result = 8; }
   if (rng >= 99) { result = 9; }
  //Rigging. DISABLED IF GLOBAL BOOLS "winner" or "rigged" are distabled!
  if (rigged == true && winner == false && frames == 207) { // rigged, place 0
   result = 0
  }
  if (rigged == true && winner == false && frames == 208) { // rigged place nine after 0
   result = 9
  }
  if (winner == true && rigged == false && frames == 207) { //winner, get 9 always
   result = 9
  }
  //End of rigging.
  arrSlice = append(arrSlice, result) // stick random number on the end of the slice (array)
  arrSlice = arrSlice[1:] //remove first element
  slicePrint(arrSlice)
  moveDown(2)
  //Ticker line switch
  switch tickerMode {
  case "None":
  fmt.Print("        / \\  ", (frames+1), "/", frameMax, "   \n")
  case "Perception":
  if frames == 207 {
   perceptive = arrSlice[63]
  }
  if frames < 207 {
   perceptive = 0
  }
  fmt.Print("        /");fmt.Printf("%s%d%s", colorMap[perceptive], perceptive, Reset); fmt.Print("\\ ",(frames+1), "/", frameMax, "   \n") // this line kills me
  case "Ticker":
   fmt.Print("       /");singleSlice(arrSlice,15);fmt.Print("\\  ", (frames+1), "/", frameMax, "   \n")
  }
  //End ticker line switch
  time.Sleep(time.Duration(rate) * time.Millisecond)
  if (frames > frameMax/2) {
   rate = rate + 1
  }
  if (frames < jam - 1) { //replace jam with frameMax if it breaks!
   moveUp(3)
  }
 }
 if (jammed == true) {
    if (arrSlice[15] > 0) { //flashing colors if you won
     flickerLine("JAMMED!", 80, BYellow, BGreen, 10, 500)
    }
    if (arrSlice[15] == 0) {
     flickerLine("JAMMED!", 80, Red, BRed, 10, 500)
    }
  jammed = false
 }
 if (arrSlice[15] > 0) {
  wins++
  winStreak++
  cash += (bet*arrSlice[15])
  //Celebration messages
   if (arrSlice[15] == 9) { // Celebrate hitting a 9
   color(BYellow)
   progWrite("INCREDIBLE 9x WIN!", 50)
   color(Reset)
   wait(500)
   progDel(last, 10)
   }
   if (cash >= 1000000 && celebrateMil == true) { // Celebrate hitting 1 mil
   color(BYellow)
   progWrite("YEEEEEEEAAAAAAAAAAAAAAH!", 50)
   color(Reset)
   wait(500)
   progDel(last, 10)
   color(BGreen)
   progWrite("BROKE ONE MILLION!", 50)
   color(Reset)
   wait(500)
   progDel(last, 10)
   celebrateMil = false
   }
   if (cash >= 100000 && celebrate1hk == true) { // Celebrate hitting 100k
   color(BGreen)
   progWrite("BROKE A HUNDRED THOUSAND!", 50)
   color(Reset)
   wait(500)
   progDel(last, 10)
   celebrate1hk = false
   }
   if (cash >= 1000 && celebrate1k == true) { // Celebrate hitting 1k
   fmt.Print(BGreen)
   progWrite("BROKE ONE THOUSAND!", 50)
   fmt.Print(Reset)
   wait(500)
   progDel(last, 10)
   celebrate1k = false
   }
  fmt.Print("Cash: $", cash, BGreen, " WON! ",BYellow,"+$", bet*arrSlice[15], Reset)
  if (cash > score) {
   score = cash
   if (winStreak >= 2) {
   color(BMagenta)
   fmt.Print(" ", winStreak, " WIN STREAK!")
   color(Reset)
   }
   if (cash >= 10000) {
     //fmt.Print(BCyan, " NEW HIGHSCORE!",Reset)
     color(BCyan) //colors the place bet line too
     wait(500)
     progWrite(" NEW HIGHSCORE!", 50)
    }
  }

 } else { // lose - array hit zero
  cash = cash - bet
  losses++
  winStreak = 0
  fmt.Print("Cash: $", cash, BRed, "   LOST -$", bet, Reset)
 }
 if (cash <= 0 && bets == 1 ) { // lost after one bet
  fmt.Print("\n", Red)
  wait(1500)
  progWrite("Buddy. ", 50)
  wait(700)
  progWrite("Yer shockin'", 50)
  wait(1000)
  fmt.Print("\n", Reset)
  os.Exit(0)
  }
 if (cash <= 0 && score == 100 && bets > 1 ) { // lost, never got over starting cash
  fmt.Print("\n", Red)
  wait(1000)
  progWrite("Walk on home, boy.", 50)
  fmt.Print("\n", Reset)
  os.Exit(0)
 }
 if (cash <= 0) { //lose sequence
  wait(1500) // suspense delay, lol
  fmt.Print("\n")
  color(BRed)
  progWrite("GAME OVER!\n", 100)
  wait(500)
  color(Red)
  progWrite("You have lost everything\n", 50)
  wait(1000)
  color(Reset)
  progWrite("-------------------------------\n",50)
  progWrite("Highest score: $",50)
  wait(1000)
  fmt.Print(score,"\n")
  wait(1000)
  progWrite("Bet count: ",20)
  wait(1000)
  fmt.Print(bets,"\n")
  wait(1000)
  os.Exit(0)
 }
 funcBet()

}

func tauntWrite (text string, speed int, wait int) {
// structured as text you want, speed it's written, time until deletion.
        progWrite(text, speed)
        time.Sleep(time.Duration(wait) * time.Millisecond)
        progDel(last, 10)

}

func color (col string) { // short way of coloring terminal. Reset with color(Reset)
        fmt.Print(col)
}

func moveUp (amt int) {
        fmt.Print("\x1b[",amt,"F")
}

func moveDown (amt int) {
        fmt.Print("\x1b[",amt,"E")
}

func homeCursor() {
        fmt.Print("\033[1G")
}

func wait (msec int) { //short way to execute time.Sleep in Milliseconds
        time.Sleep(time.Duration(msec) * time.Millisecond)
}

func quit() {
        os.Exit(0)
}

func startupChecks() {
        if (winner == true && rigged == true) {
         write("Winner and Rigged are both true.\nDisable one and recompile")
         quit()
        }
}
func write (text string) { //do I hate writing fmt.Println()? Yes.
        fmt.Println(text)
}

func charCount(str string) int {
        return len(str)
}

func flickerLine(text string, speed int, color1 string, color2 string, times int, displayTime int) {
// text is what's written, speed is how fast it flickers, color 1 and 2 will alternate, times is how many flickers
        for i := 0; i <= times ; i++ {
         homeCursor() // make sure she's home
         color(color1) //set text to color 1
         fmt.Print(text) // write it out
         wait(speed) // wait a bit
         homeCursor() // back to the start of the line
         color(color2) // change color
         fmt.Print(text) //reprint text in new color
         wait(speed) // wait some more
        }
    wait(displayTime) // how long the message stays on the last color
    letters :=  charCount(text) // get letters amount from the input string
    progDel(letters, 50) // remove them at usual speed
    color(Reset) // reset color
    //done
}

func achWrite (colour string, achievement string, stop int, speed int) {
//Structured as color you want,  achievement text, pause time, speet it's written
        color(colour)
        progWrite("Achievement unlocked: ", 50)
        wait(stop)
        progWrite(achievement, speed)
        wait(1500)
        progDel(last, 50)
        progDel(22, 50) // remove "Achievement Unlocked: " text
        color(Reset)// not needed, good to have
}

func singleSlice (slice []int, pos int) { //name the slice, name the position
        fmt.Printf("%s%d%s", colorMap[slice[pos]], slice[pos], Reset)
}

func slicePrint (slice []int) { //prints a slice without commas, spaces or []
     for i := 1; i < len(slice); i++ {
        fmt.Printf("%s%d%s", colorMap[slice[i]], slice[i], Reset)
    }
}

func introText() {
        sw := rand.Intn(11) // RNG to pick a quote. Value stored in "sw" variable
        color(BRed)
        switch sw {
        case 0:
        progWrite("     BET     YOUR    MAX     \n", 200)
        case 1:
        progWrite("     DO      YOUR    WORST   \n", 200)
        case 2:
        progWrite("     TIME    HEALS   NOTHING \n", 200)
        case 3:
        progWrite("     IT'S    ALL     OVER    \n", 200)
        case 4:
        progWrite("     MONEY   FOR     NOTHING \n", 200)
        case 5:
        progWrite("     WHO     LAUGHS  LAST?   \n", 200)
        case 6:
        progWrite("     BACK    SO      SOON?   \n", 200)
        case 7:
        progWrite("     GREED   LOSS    DEATH   \n", 200)
        case 8:
        progWrite("     HOUSE   ALWAYS  WINS    \n", 200)
        case 9:
        progWrite("     GREAT   DEBT    AWAITS  \n", 200)
        case 10:
        progWrite("     TO      YOUR    GRAVE!  \n", 200)
        }
        color(Reset)
}

func msgAllIn() {
  rc := rand.Intn(6) // rc (random case)
  color(BRed)
  switch rc {
   case 0:
   tauntWrite("You might regret that...", 50, 500)
   case 1:
   tauntWrite("All in?", 50, 500)
   case 2:
   tauntWrite("Are you sure?", 50, 500)
   case 3:
   tauntWrite("Bad idea...", 50, 500)
   case 4:
   tauntWrite("No going back...", 50, 500)
   case 5:
   tauntWrite("Is that wise?", 50, 500)
  }
  color(Reset)
}
