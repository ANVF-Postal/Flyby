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
//bools
 celebrateMil = true // Celebrate breaking over one million
 celebrate1k  = true // Celebrate breaking a thousand
 celebrate1hk = true // Celebrate breaking a hundred thousand
 ach69   = false
 ach404  = false
 ach420  = false
 ach808  = false
 ach1337 = false
 jamming = true //Enable or disable jamming
 jammed  = false // This is set if the machine jams
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

    colorMap = map[int]string{ // set the color of each number
    0: Gray,
    1: White,
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
//intro movie
        cash = 100
        score = cash
        winStreak = 0
        rand.Seed(time.Now().UnixNano())
        progWrite("---FLYBY------V------------------------------------------------\n", 50)
        sw := rand.Intn(11) // RNG to pick a quote. Value stored in "sw" variable
        //fmt.Print(BRed) // color the following quotes
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
        //fmt.Print(Reset) //reset terminal color
        color(Reset)
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
 if (bet == cash && cash >= 10000) {
  rc := rand.Intn(5) // rc (random case)
  color(BBlue)
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
   tauntWrite("Greed...", 50, 500)
  }
  color(Reset)
 }
 if (bet <= 0) { // Quit sequence
  progWrite("You decided to quit!", 50)
  time.Sleep(2 * time.Second)
  fmt.Println(BYellow,"\n--------------------------------------", Reset)
  time.Sleep(1000 * time.Millisecond)
  fmt.Println("Final balance ",BYellow, "$",cash, Reset)
  time.Sleep(1000 * time.Millisecond)
  fmt.Println("Highest score ",BGreen,"$",score, Reset)
  time.Sleep(1000 * time.Millisecond)
  fmt.Println("Bet count      ",bets)
  time.Sleep(1000 * time.Millisecond)
  fmt.Println("Win/loss       ",wins,"/",losses)
  time.Sleep(2 * time.Second)
  color(BCyan)
  progWrite("Thank you for playing!",50)
  color(Reset)
  time.Sleep(1 * time.Second)
  fmt.Print("\n")
  os.Exit(0)
 }
 bets++
 funcRoll(bet)
}

func funcRoll(bet int) {
 clean(5)
 rate = 10 // old value: 8
 frameMax := 256
 if (jamming == true) { // jamming function
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
  arrSlice = append(arrSlice, result) // stick random number on the end of the slice (array)
  arrSlice = arrSlice[1:] //remove first element
  slicePrint(arrSlice)
  fmt.Println("--------------^------------------------------------------------")
  fmt.Println("      / \\  ", (frames+1), "/", frameMax)
  time.Sleep(time.Duration(rate) * time.Millisecond)
  if (frames > frameMax/2) {
   rate = rate + 1
  }
  if (frames < jam - 1) { //replace jam with frameMax if it breaks!
    clean(3)
  }
 }
 if (jammed == true) {
  progWrite("JAMMED!", 50)
  time.Sleep(1 * time.Second)
  progDel(last, 50)
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
   time.Sleep(500 * time.Millisecond)
   progDel(last, 10)
   }
   if (cash >= 1000000 && celebrateMil == true) { // Celebrate hitting 1 mil
   color(BYellow)
   progWrite("YEEEEEEEAAAAAAAAAAAAAAH!", 50)
   color(Reset)
   time.Sleep(500 * time.Millisecond)
   progDel(last, 10)
   color(BGreen)
   progWrite("BROKE ONE MILLION!", 50)
   color(Reset)
   time.Sleep(500 * time.Millisecond)
   progDel(last, 10)
   celebrateMil = false
   }
   if (cash >= 100000 && celebrate1hk == true) { // Celebrate hitting 100k
   color(BGreen)
   progWrite("BROKE A HUNDRED THOUSAND!", 50)
   color(Reset)
   time.Sleep(500 * time.Millisecond)
   progDel(last, 10)
   celebrate1hk = false
   }
   if (cash >= 1000 && celebrate1k == true) { // Celebrate hitting 1k
   fmt.Print(BGreen)
   progWrite("BROKE ONE THOUSAND!", 50)
   fmt.Print(Reset)
   time.Sleep(500 * time.Millisecond)
   progDel(last, 10)
   celebrate1k = false
   }
  //cash += (bet*arrSlice[15])
  fmt.Print("Cash: $", cash, BGreen, " WON! ",BYellow,"+$", bet*arrSlice[15], Reset)
  if (cash > score) {
   score = cash
   if (winStreak >= 2) {
   fmt.Print(BMagenta)
   fmt.Print(" ", winStreak, " WIN STREAK!")
   fmt.Print(Reset)
   }
   if (cash >= 10000) {
    fmt.Print(BCyan, " NEW HIGHSCORE!",Reset)
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
  progWrite("Buddy, yer shockin'", 50)
  fmt.Print("\n", Reset)
  os.Exit(0)
 }
 if (cash <= 0) { //lose sequence
  fmt.Print("\n")
  fmt.Print(BRed)
  progWrite("GAME OVER!\n", 200)
  fmt.Print(Red)
  progWrite("You have lost everything\n", 100)
  fmt.Print(Reset)
  progWrite("-------------------------------\n",50)
  progWrite("Highest score: $",50)
  time.Sleep(1 * time.Second)
  fmt.Print(score,"\n")
  time.Sleep(1 * time.Second)
  progWrite("Bet count: ",20)
  time.Sleep(1 * time.Second)
  fmt.Print(bets,"\n")
  time.Sleep(1 * time.Second)
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

func slicePrint (slice []int) { //prints a slice without commas, spaces or []
     for i := 1; i < len(slice); i++ {
        //fmt.Print(slice[i])
        fmt.Printf("%s%d%s", colorMap[slice[i]], slice[i], Reset)
    }
    fmt.Print("\n")
}
