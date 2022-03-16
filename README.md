# Book 2 Ships
Reproducing the original Traveller Book 2 Ship Design System in a simple UI with some automation - around crew and staterooms, tonnage and eventually costs.

Playing with Go UI code to learn more is the primary motivation, but I might as well make a useful tool (well, useful for a nerd like me) while I am at it.

## Book 2 Notes
Book 2 uses letter selections for drives - A-Z skipping I and O, so there 24 available drives of assorted sizes and costs. Only a few of those drives can be used for any given hull tonage.

Hull selections go from 1-9 then A-Z, again skipping I and O, so therre are 33 hull sizes. I've added a hull size 0 for < 100 tons, not sure if that's useful.

# Design
I'm attempting to invert dependencies, kinda. Main should pass the detail & box widgets to the different UI oriented bits - one for weapons, one for drives, the hull, staterooms, etc.

On intializatioon each UI component adds it's settings components (selections available, maybe drive A-Z or number of staterooms) to the passed in settings panel and adds its text output from those settings to the box on it's own line or lines as a text box so it can update and refresh the output as needed. Speaking of which it also sets up the onChange methods so that it is notified as needed.

On the onChanged method call the UI component processes the user selection and updates the UI component's bit of it put into the passed-in box and refreshes the detail UI.

There's complexity around side effects - go to a bigger engine and the required number of engineers increses so additional staterooms may be required...
##Building
Simple golang executable, just do this from bash or Windows/Mac equivalent:

``` bash
git clone https://github.com/DavidLDawes/Book-2-Ships.git
cd Book-2-Ships
go build .
```
The result should be a sbip-ui executable. Mac is the same. Windows probably makes it ship-ui.exe, haven't checked. Not sure if the UI works on Linux or Windows, have not tried it. Most likely Windows needs some compatibility stuff installed.

## Running
If the build step above works, simply run ship-ui from the same top level project directory. For bash:
``` ./ship-ui```

To exit hit < Command >q on Mac, < Alt > F4 on Windows, and perhaps <Ctrl>C from the bash script in Linux? Only been running on Mac so far.

##Current State
Drives and hulls are pretty well plumbed in, I think. Limited to available options per hull, showing tonnage, cost and performance.   The UI properly limits you to selecting usable drives, it would be better to process things and change the drop down to only show selectable drives for the given hull, but for now it ignores impossible requests.

Vehicles can be fiddled with, not reporting any tonnage or cost yet.

Berths are partially there, Staterooms slider not working though. Low and Emergency Low berths can be configured and I think it shows tonnage, need to add costs.

I don't think armored bulkheads are supported properly.

Weapons can be sleected in various numbers per weapon, not dsure what it reports for those.
## Goals
I'd like to get all of the selectable options enhanced to show tonnage and cost, and ideally effect (i.e., drives show P-2 etc., Lasers show cost, tonnage and USP value, diutto for missiles).

I'd like to limnit selections (weapons, vehicles) to what can actually fit into the remaining available tonnage.

Vehicles need the adjustment for berth space.

Need work on crews, I think gunners and pilots & service for vehicles are listed. Need to make sure all are totalled and the number of Stewards is automatically figured.

Version 1 limit berths for Crew to what is needed to handle staffing, plus a drop down that limits selections to what can fit for passengers and Low & Emergency Low berths.

Version 2 allows berths for Crew to add in x% (up to 15% in increments of 1, rounded up) to allow overstaffing.

Version 2.5 Adds in Staff Low Berths of correct size (60%?) to qualifyh for frozen watch.

Version 3 breaks military out from civilian/commercial. limiting Hard Points and weapons on larger commercial vessels and allowing minimal fighters on commercial vessles too.

Military vessels need Launch Tube options. Etc.