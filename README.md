# iotaZeroBalanceHelper


## What is it for?
This program was written to help Iota users who unexpectedly see a zero balance in their Trinity wallet. There are multiple reasons why this could be happening and often it is not easy for non-experts to find out what is going on.

## What does it do?
The goal is that the program lets unexperienced users test what could possibly cause Trinity to not show the expected balance. I plan to add more tests over time.
The program allows to 
1. check if funds on addresses of a seed were taken into custody by the Iota Foundation in late 2017. If this is the case the affected funds are reclaimable.
2. check if there is any balance that Trinity might not show. The program will list all addresses with balance while Trinity will often only show a single address if no funds are found. 
3. export a list of generated addresses to be tested with the online tool [iotaAddressHistoryChecker](https://github.com/HBMY289/iotaAddressHistoryChecker)
4. check for typos in the seed (not implemented yet)

Seeds that are too short or too long will be adadpted in the same way the Iota Light Wallet did.


## How it works
##### Reclaim
With the given seed the programm calculates a large number of addresses both using the current address generation algorith and the one that was used in 2017 and earlier. These addresses are then matched against a list of reclaimable addresses. Possible matches and reclaimable balance will be shown. If no match is found using the Iota reclaim process will NOT bring back your balance.

##### Balance
If a reclaim case could be ruled out, the balance option will generate addresses using the current algorithm and compare them to an internally stored snapshot of the tangle. This snapshot holds all addresses with balances and by matching against this list the program does not need any connection to the internet or the current tangle to check balances. Currently a snapshot file from Feb 12th, 2021 is included. If you are missing a funds that have been moved after this date you can supply a more recent snapshot file. Ask someone who has access to a Iota node to run the following command and supply you with the resulting file.
```
curl -H 'X-IOTA-API-VERSION: 1' -d '{"command":"getLedgerState"}' localhost:14265 >  snapshot.txt
```
Place the file `snapshot.txt` next to this tool's executable. It will be automatically be used when checking the balances.

##### Export Addresses to check online
This program requires to enter your seed and is designed to run without any internet connection. If all your addresses still do not show any balance you can check the addresses on the Iota tangle explorer [explorer.iota.org](https://explorer.iota.org) for any transactions that show where your funds went. To avoid having to enter all addresses manually I wrote another tool that automates this process ([iotaAddressHistoryChecker](https://github.com/HBMY289/iotaAddressHistoryChecker)). It will require a list of addresses exported by the iotaZeroBalanceHelper and then request all available tranascations from the explorer. It will report any token movements in a short and human readable report.

## Disclaimer
NEVER share your seed with anyone. No community member or member of the Iota Foundation will ever ask for your seed. If someone does it is 100% a scam to steal your money. That said, even entering your seed into a software other than the official Iota wallet should not be handled lightly and can only be recommended as a last resort. For your own safety you should run this software only an an air-gapped computer. 

## How to start the tool
The simplest way is to download the appropriate binary executable for your operating system from [releases](https://github.com/HBMY289/iotaZeroBalanceHelper/releases). You can also build the tool from source, which is rather easy as well. Assuming you have [go](https://golang.org/doc/install) and [git](https://www.atlassian.com/git/tutorials/install-git) installed already you can just execute this command for example in your user folder to get a copy of the source code.
```
git clone https://github.com/HBMY289/iotaZeroBalanceHelper.git
```

Then you change into the new folder and build the excutable.
```
cd iotaZeroBalanceHelper
go build
```
After that you can just start the newly created binary file by typing
```
./iotaZeroBalanceHelper
```
or on Windows
```
iotaZeroBalanceHelper.exe
```
## How to use the tool
Once the program is running you will have to enter the seed and choose from the available options.

##### Number of addresses
The program will ask yoiu how many addresses of your seed you want to check. In most cases 100 addresses will suffice, only if you heavily used your seed for sending back and forth and genereated a lot of addresses a higher number might be necessary.

## Need additonal help?
If you need any additonal help either with the tool itself or with checking the exported addresses you can contact me (HBMY289) or any other community member in the #help channel via the official via the official [Iota Discord server](https://discord.iota.org/).
