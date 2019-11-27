# IOST Blockchain Explorer

### Prerequisites
* Golang 1.10.1 (or newer) is required to build this project
* Mongodb 4.0 (or newer) is required, We recommend Redis stable version
* node.js 8.10   (or newer) is required

## Installation
```bash
get repo first:
git clone https://github.com/iost-official/explorer
```

### Frontend
```bash
#  Change Directory
cd explorer/frontend/

# build
npm install

# run in test
npm run dev

# run production build
npm run build
```

### Backend
```bash
#  Change Directory
cd explorer/backend/

# build
make

# run
./explorer
```

### Mongo Cron Task
```bash
#  Change Directory
cd explorer/backend/task

# build
make

# run
./explorer-task
```
## Contribution

Contribution of any forms is appreciated!

Currently, our core tech team is working intensively to develop the first stable version and core block chain structure. We will accept pull request after the first stable release published.

If you have any questions, please email to team@iost.io

## Community & Resources

Make sure to check out these resources as well for more information and to keep up to date with all the latest news about IOST project and team.

[IOSToken on Reddit](https://www.reddit.com/r/IOStoken)

[Telegram](https://t.me/officialios)

[Twitter](https://twitter.com/IOStoken)

[Official website](https://iost.io)

## Disclaimer

- IOS Blockchain is unfinished and some parts are highly experimental. Use the code at your own risk.

- THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

Transaction
{ 
   "txHash":"HBaTMkom5KMwMgJEELzScFeTiHmZCw8checfDDvGzfoT",
   "blockHeight":2026271,
   "from":"admin",
   "to":"token.iost",
   "amount":0,
   "gasLimit":1000000,
   "price":1,
   "age":"24 days ago",
   "utcTime":"2019-11-03 21:29:39.4038 +0900 JST",
   "code":"",
   "statusCode":0,
   "statusMessage":"",
   "contract":"token.iost",
   "actionName":"balanceOf",
   "data":"[\"silver_token\",\"test1\"]",
   "memo":"",
   "expiration":1572784269403800000,
   "Actions":[ 
      { 
         "contract":"token.iost",
         "action_name":"balanceOf",
         "data":"[\"silver_token\",\"test1\"]"
      }
   ],
   "signer":[ 

   ],
   "publisher":"admin",
   "receipt":"",
   "AmountLimit":null,
   "Receipt":{ 
      "tx_hash":"HBaTMkom5KMwMgJEELzScFeTiHmZCw8checfDDvGzfoT",
      "gas_usage":4695,
      "returns":[ 
         "[\"1200\"]"
      ],
      "Receipts":null
   }
}