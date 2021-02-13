# iotaZeroBalanceHelper


## What is it for?
This program was written to help Iota users who unexpectedly see a zero balance in their Trinity wallet. There are multiple reasons why this could be happening and often it is not easy for non-experts to find out what is going on.

## What does it do?
The goal is that the program lets unexperienced users test what could possibly cause Trinity to not show the expected balance. I plan to add more tests over time.
Currently the tool can test
1. if funds on addresses of a seed were taken into custody by the Iota Foundation in late 2017. If this is the case the affected funds are reclaimable.
2. if there is any balance that Trinity might not show. The program will list all addresses with balance and also allows exporting the generated addresses even if no balance is found. While Trinity will often only show a single address if no funds are found. The exported addresses can be supplied to other more experienced users to check if any funds had been on these addresses. 
3. if the seed contains typo (not implemented yet)
4. if there ever have been funds on addresses of the seed and if so, where they went to (not implemeneted yet). This mode would require an online connection and might not be implemented at all for security reasons

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
