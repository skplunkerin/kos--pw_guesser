# Project Kyzen - K.OS Password Guesser

This project was created to help crack the password through brute forcing all
possible 5-digit pin combinations for the `glitch/travels/labnotes/TOPSECRT/`
directory, where the only hint we had was:

> "Mr. 'But we need more data' is suddenly concerned
> about data. Okay - I'll encrypt the results with
> Shift 5 and move this to my new "TOPSERT" directory.
> Tether, the password is your old 5-digit phone pin.
> Do you still remember that, or did that get fried
> along with your memory of tacos?"

A picture of Tether's old phone was also found, where zooming in showed that 4
of the number keys were dirtier(?) than the others, which helped to restrict
what digits to create a 5-digit pin from.

This project helped do the following:

1. create a list of all possible 5-digit number variations for the given 4 int's
2. pop the first pin off the top of the list of combinations to attempt and
   store it in the clipboard
3. append the pin copied to clipboard to the end of attempted pins file

Then after creating a build (`go build`) of this logic, and with the help of
Stream Deck (click a button to trigger an Automator app), and Automator (a Mac
tool I used to create a looping program), I was able to attempt pin-after-pin
until I found the right combination.

## Stream Deck

Create button to `System - Open` the Automator app. I created different loop
Automator apps _(x10, x25, x50, x100)_ and different buttons for each of them.

## Automator app

### App steps:

This took some trial-and-error as `K.OS` has issues if this goes too fast, as
well as Automator going faster than the clipboard can handle; unfortunately the
`1 second` pause steps are the lowest possible pause amounts _(it took me about_
_~6:20mins to process 100 digits; but there is too much room for error (ie,_
_once the program starts, it has to continue until it finishes; if anything_
_interrupts the programs like a notification, an update, accidentally changing_
_to another app, etc, then you have to manually track what digits were actually_
_attempted, undo changes to the attempts files, and try again. In the end I_
_opted for doing batches of 10-25 attempts at a time after I had multiple 100_
_batch attempts get ruined.))_

1. Run AppleScript

   This will type `dsnd TOPSECRT` and then hit `enter`; make sure you're already
   in the `glitch/travels/labnotes/` dirt before running this.

   - Input:

     ```sh
     on run {input, parameters}
     	tell application "System Events"
     		keystroke "dsnd TOPSECRT"
     		keystroke return
     	end tell
     end run
     ```

   - Options:
     - Check `Ignore this action's input`

2. Pause

   - Duration: `1 second`
   - Options:
     - Check `Ignore this action's input`

3. Run Shell Script

   - Input:

     ```sh
     /full/path/to/your/build/kos--pw_guesser
     ```

   - Options:
     - Check `Ignore this action's input`

4. Pause

   - Duration: `1 second`
   - Options:
     - Check `Ignore this action's input`

5. Run AppleScript

   This will trigger `cmd+v` to paste what's in the clipboard.

   - Input:

     ```sh
     on run {input, parameters}
     	tell application "System Events"
     		keystroke "v" using command down
     	end tell
     end run
     ```

   - Options:
     - Check `Ignore this action's input`

6. Pause

   - Duration: `1 second`
   - Options:
     - Check `Ignore this action's input`

7. Run AppleScript

   This will trigger `enter` to submit the pasted digit.

   - Input:

     ```sh
     on run {input, parameters}
     	tell application "System Events"
     		keystroke return
     	end tell
     end run
     ```

   - Options:
     - Check `Ignore this action's input`

8. Pause

   - Duration: `1 second`
   - Options:
     - Check `Ignore this action's input`

9. Loop

   - Settings:

     - Loop automatically
     - Stop after `x` times _(set this to `10`, `25`, `50`, `75`, or whatever)_
     - Use the current results as input

   - Options:
     - Check `Ignore this action's input`

Then save the program, name it, and test + grant the needed permissions.

### Granting access

In Mac settings -> Privacy & Security -> Accessiblity, add the app(s) to grant
access. Then after running the app for the first time (test this in a text
editor first, reducing the loop to 1), click to approve the additional access
needed.
