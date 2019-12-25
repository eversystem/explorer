Iost explorer backend
======

j-LOD STG
  "rpcHost": "172.31.43.158:30002",



### Requirements 

1. golang 1.10+
2. mongodb



### How to deploy

1. clone code to $GOPATH/src/github.com/iost-official/explorer
2. add config file cp backend/config/config.json.sample backend/config/config.json
    update the ```rpcHost```, ```mongodb``` config if need
3. Run blockchain sync task: 
```bash
cd backend/task
make
nohup ./explorer-task &
```
4. Run REST api service
```bash
cd backend
make
nohup ./explorer &
```

```
{ 
   "id":"base.iost",
   "code":"_IOSTInstruction_counter.incr(5.6);const producerPermission = 'active';\n_IOSTInstruction_counter.incr(3);const voteStatInterval = 1200;\n_IOSTInstruction_counter.incr(3);const issueInterval = 172800;\n_IOSTInstruction_counter.incr(189);class Base {\n    init() {\n        _IOSTInstruction_counter.incr(11.5);this._put('execBlockNumber', 0);\n    }\n    initAdmin(adminID) {\n        _IOSTInstruction_counter.incr(7);const bn = block.number;\n        if (_IOSTInstruction_counter.incr(6.3),_IOSTBinaryOp(bn, 0, '!==')) {\n            _IOSTInstruction_counter.incr(62.5);throw new Error('init out of genesis block');\n        }\n        _IOSTInstruction_counter.incr(10.7);this._put('adminID', adminID);\n    }\n    can_update(data) {\n        _IOSTInstruction_counter.incr(13.7);const admin = this._get('adminID');\n        _IOSTInstruction_counter.incr(8);this._requireAuth(admin, producerPermission);\n        return true;\n    }\n    _requireAuth(account, permission) {\n        _IOSTInstruction_counter.incr(11);const ret = blockchain.requireAuth(account, permission);\n        if (_IOSTInstruction_counter.incr(6.3),_IOSTBinaryOp(ret, true, '!==')) {\n            _IOSTInstruction_counter.incr(68.8);throw new Error(_IOSTBinaryOp('require auth failed. ret = ', ret, '+'));\n        }\n    }\n    _get(k) {\n        _IOSTInstruction_counter.incr(11);const val = storage.get(k);\n        if (_IOSTInstruction_counter.incr(8.3),_IOSTBinaryOp(val, '', '===')) {\n            return null;\n        }\n        _IOSTInstruction_counter.incr(8);return JSON.parse(val);\n    }\n    _put(k, v) {\n        _IOSTInstruction_counter.incr(16);storage.put(k, JSON.stringify(v));\n    }\n    _vote() {\n        _IOSTInstruction_counter.incr(18.200000000000003);blockchain.callWithAuth('vote_producer.iost', 'stat', _IOSTTemplateTag`[]`);\n    }\n    _bonus(data) {\n        _IOSTInstruction_counter.incr(16.6);blockchain.callWithAuth('bonus.iost', 'issueContribute', [data]);\n    }\n    _issue() {\n        _IOSTInstruction_counter.incr(17.9);blockchain.callWithAuth('issue.iost', 'issueIOST', _IOSTTemplateTag`[]`);\n    }\n    _saveBlockInfo() {\n        _IOSTInstruction_counter.incr(14.8);let json = storage.get('current_block_info');\n        _IOSTInstruction_counter.incr(29.200000000000003);storage.put(_IOSTBinaryOp('chain_info_', block.parentHash, '+'), JSON.stringify(json));\n        _IOSTInstruction_counter.incr(19.8);storage.put('current_block_info', JSON.stringify(block));\n    }\n    _saveWitnessInfo() {\n        _IOSTInstruction_counter.incr(27.8);let map = JSON.parse(storage.get('witness_produced') || '{}');\n        _IOSTInstruction_counter.incr(28.1);map[block.witness] = _IOSTBinaryOp(map[block.witness] || 0, 1, '+');\n        _IOSTInstruction_counter.incr(19.6);storage.put('witness_produced', JSON.stringify(map));\n    }\n    _clearWitnessInfo() {\n        _IOSTInstruction_counter.incr(11.6);storage.del('witness_produced');\n    }\n    exec(data) {\n        _IOSTInstruction_counter.incr(7);const bn = block.number;\n        _IOSTInstruction_counter.incr(14.5);const execBlockNumber = this._get('execBlockNumber');\n        if (_IOSTInstruction_counter.incr(6.3),_IOSTBinaryOp(bn, execBlockNumber, '===')) {\n            return true;\n        }\n        _IOSTInstruction_counter.incr(11.5);this._put('execBlockNumber', bn);\n        _IOSTInstruction_counter.incr(8);this._saveBlockInfo();\n        _IOSTInstruction_counter.incr(8);this._saveWitnessInfo();\n        if (_IOSTInstruction_counter.incr(29.7),_IOSTBinaryOp(_IOSTBinaryOp(bn, voteStatInterval, '%'), 0, '===') \u0026\u0026 _IOSTBinaryOp(data.parent[2], false, '===')) {\n            _IOSTInstruction_counter.incr(8);this._vote();\n            _IOSTInstruction_counter.incr(8);this._clearWitnessInfo();\n        }\n        if (_IOSTInstruction_counter.incr(12.399999999999999),_IOSTBinaryOp(_IOSTBinaryOp(bn, issueInterval, '%'), 0, '===')) {\n            _IOSTInstruction_counter.incr(8);this._issue();\n        }\n        _IOSTInstruction_counter.incr(8);this._bonus(data);\n    }\n}\n_IOSTInstruction_counter.incr(7);module.exports = Base;",
   "language":"javascript",
   "version":"1.0.0",
   "abis":[ 
      { 
         "name":"exec",
         "args":[ 
            "json"
         ],
         "amount_limit":[ 

         ]
      },
      { 
         "name":"can_update",
         "args":[ 
            "string"
         ],
         "amount_limit":[ 

         ]
      },
      { 
         "name":"initAdmin",
         "args":[ 
            "string"
         ],
         "amount_limit":[ 

         ]
      },
      { 
         "name":"init",
         "args":[ 

         ],
         "amount_limit":[ 

         ]
      }
   ]
}
```

